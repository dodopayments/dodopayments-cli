import { afterAll, beforeAll, beforeEach, describe, expect, spyOn, test } from 'bun:test';
import { createCipheriv, pbkdf2Sync, randomBytes } from 'node:crypto';
import { existsSync, mkdirSync, mkdtempSync, rmSync, writeFileSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { dirname, join } from 'node:path';
import { version } from '../../package.json';
import { configPath, unreadableStoreMessages } from './auth';
import * as oauth from './oauth';
import type { OAuthEndpoints } from './oauth';

// scripts/test-preload.ts sets DODO_CONFIG_DIR to a temporary directory.

type Form = URLSearchParams;
type Reply = { status: number; body: unknown };
let handler: (path: string, form: Form) => Reply | Promise<Reply> = () => ({ status: 500, body: {} });
const calls: { path: string; form: Form; headers: Headers }[] = [];

const server = Bun.serve({
  port: 0,
  async fetch(request) {
    const path = new URL(request.url).pathname;
    const form = new URLSearchParams(await request.text());
    calls.push({ path, form, headers: request.headers });
    const { status, body } = await handler(path, form);
    return Response.json(body, { status });
  },
});

const ep: OAuthEndpoints = {
  issuer: `http://127.0.0.1:${server.port}`,
  resources: { test_mode: 'https://test.example', live_mode: 'https://live.example', mcp: 'https://mcp.example/mcp' },
};

let counter = 0;
const token = (resource: string) => ({
  access_token: `at_${resource}_${++counter}`,
  refresh_token: `rt_${counter}`,
  expires_in: 3600,
  scope: 'payments:read',
});

beforeAll(() => {
  process.env.DODO_DEVICE_NAME = 'tester@box';
});

beforeEach(() => {
  calls.length = 0;
  oauth.deleteStore();
});

afterAll(() => {
  server.stop(true);
});

/** A grant with one mode: the resource of the other mode gets `invalid_target`. */
const singleModeGrant =
  (mode: 'test_mode' | 'live_mode') =>
  (path: string, form: Form): { status: number; body: unknown } => {
    if (path === '/oauth2/device_authorization') {
      return {
        status: 200,
        body: { device_code: 'dc', user_code: 'BCDF-GHJK', verification_uri: `${ep.issuer}/device`, expires_in: 60, interval: 0 },
      };
    }
    if (path === '/oauth2/token') {
      const grant = form.get('grant_type');
      if (grant === 'urn:ietf:params:oauth:grant-type:device_code') {
        const pending = calls.filter((c) => c.form.get('grant_type') === grant).length < 2;
        return pending ? { status: 400, body: { error: 'authorization_pending' } } : { status: 200, body: token('mcp') };
      }
      const other = mode === 'test_mode' ? ep.resources.live_mode : ep.resources.test_mode;
      if (form.get('resource') === other) return { status: 400, body: { error: 'invalid_target' } };
      return { status: 200, body: token(String(form.get('resource'))) };
    }
    if (path === '/oauth2/revoke') return { status: 200, body: {} };
    return { status: 404, body: {} };
  };
const testOnlyGrant = singleModeGrant('test_mode');
const liveOnlyGrant = singleModeGrant('live_mode');

async function signIn(handle: typeof testOnlyGrant) {
  handler = handle;
  const result = await oauth.pollDeviceFlow(await oauth.startDeviceFlow(ep), ep);
  calls.length = 0;
  return result;
}

async function writeRaw(value: unknown) {
  const { writeEncrypted } = await import('./secure-store');
  writeEncrypted(oauth.storePath(), JSON.stringify(value));
}

describe('device flow', () => {
  test('polls until the approval and stores the modes of the grant', async () => {
    handler = testOnlyGrant;
    const start = await oauth.startDeviceFlow(ep);
    expect(start.user_code).toBe('BCDF-GHJK');
    const authorize = calls[0]!.form;
    expect(authorize.get('client_id')).toBe('dodo_cli');
    expect(authorize.getAll('resource')).toEqual([ep.resources.test_mode, ep.resources.live_mode, ep.resources.mcp]);

    const { grant, replaced } = await oauth.pollDeviceFlow(start, ep);
    expect(grant.modes).toEqual(['test_mode']);
    expect(grant.device_name).toBe('tester@box');
    expect(replaced).toEqual([]);
    expect(oauth.oauthModes()).toEqual(['test_mode']);
    expect(await oauth.oauthCredentials('test_mode', ep)).toBe(grant.tokens[ep.resources.test_mode]!.access_token);
  });

  test('sends the CLI User-Agent on each call and the device name on device_authorization', async () => {
    handler = testOnlyGrant;
    await oauth.pollDeviceFlow(await oauth.startDeviceFlow(ep), ep);
    expect(calls[0]!.path).toBe('/oauth2/device_authorization');
    expect(calls[0]!.form.get('device_name')).toBe('tester@box');
    const userAgent = `dodopayments-cli/${version} (${process.platform} ${process.arch})`;
    expect(calls.length).toBeGreaterThan(1);
    for (const call of calls) expect(call.headers.get('user-agent')).toBe(userAgent);
  });

  test('stops on access_denied', async () => {
    handler = (path, form) =>
      path === '/oauth2/token' ? { status: 400, body: { error: 'access_denied' } } : testOnlyGrant(path, form);
    const start = await oauth.startDeviceFlow(ep);
    await expect(oauth.pollDeviceFlow(start, ep)).rejects.toMatchObject({ code: 'access_denied' });
    expect(oauth.readGrants()).toEqual([]);
  });
});

describe('one grant for each mode', () => {
  test('a live sign-in keeps the test sign-in, and each mode uses its own grant', async () => {
    const test = (await signIn(testOnlyGrant)).grant;
    const live = await signIn(liveOnlyGrant);
    expect(live.replaced).toEqual([]);
    expect(calls.some((c) => c.path === '/oauth2/revoke')).toBe(false);
    expect(oauth.oauthModes()).toEqual(['test_mode', 'live_mode']);
    expect(oauth.grantFor('test_mode')!.id).toBe(test.id);
    expect(oauth.grantFor('live_mode')!.id).toBe(live.grant.id);
    expect(await oauth.oauthCredentials('live_mode', ep)).toBe(live.grant.tokens[ep.resources.live_mode]!.access_token);
  });

  test('a second sign-in of the same mode replaces the first and revokes it', async () => {
    const first = (await signIn(testOnlyGrant)).grant;
    handler = testOnlyGrant;
    const second = await oauth.pollDeviceFlow(await oauth.startDeviceFlow(ep), ep);
    expect(second.replaced.map((grant) => grant.id)).toEqual([first.id]);
    expect(second.revokeErrors).toEqual([]);
    const revoke = calls.find((c) => c.path === '/oauth2/revoke');
    expect(revoke!.form.get('token')).toBe(first.refresh_token);
    expect(oauth.readGrants().map((grant) => grant.id)).toEqual([second.grant.id]);
  });

  test('reads the v3.7 store with one grant', async () => {
    const { grant } = await signIn(testOnlyGrant);
    const { id: _id, ...legacy } = { ...grant, modes: ['test_mode', 'live_mode'] };
    await writeRaw(legacy);
    expect(oauth.oauthModes()).toEqual(['test_mode', 'live_mode']);
    expect(oauth.grantFor('live_mode')!.id).toBe('v3.7');
  });
});

describe('refresh', () => {
  test('refreshes a token that is near the expiry and keeps the new refresh token', async () => {
    const { grant } = await signIn(testOnlyGrant);
    const cached = grant.tokens[ep.resources.mcp]!.access_token;
    expect(await oauth.mcpAccessToken('test_mode', ep)).toBe(cached);

    grant.tokens[ep.resources.mcp]!.expires_at = Date.now();
    await writeRaw({ version: 2, grants: [grant] });

    const next = await oauth.mcpAccessToken('test_mode', ep);
    expect(next).not.toBe(cached);
    expect(calls[0]!.form.get('refresh_token')).toBe(grant.refresh_token);
    expect(oauth.grantFor('test_mode')!.refresh_token).not.toBe(grant.refresh_token);
  });

  test('deletes only the grant that gets invalid_grant', async () => {
    const test = (await signIn(testOnlyGrant)).grant;
    const live = (await signIn(liveOnlyGrant)).grant;
    test.tokens = {};
    await writeRaw({ version: 2, grants: [test, live] });
    handler = () => ({ status: 400, body: { error: 'invalid_grant' } });
    await expect(oauth.oauthCredentials('test_mode', ep)).rejects.toMatchObject({ code: 'invalid_grant' });
    expect(oauth.oauthModes()).toEqual(['live_mode']);
  });

  test('says which mode has no sign-in', async () => {
    await signIn(testOnlyGrant);
    await expect(oauth.oauthCredentials('live_mode', ep)).rejects.toThrow('No Live Mode sign-in');
  });
});

describe('refresh lock', () => {
  const noRefreshReply = (path: string, form: Form): Reply | Promise<Reply> =>
    form.get('grant_type') === 'refresh_token' ? new Promise<Reply>(() => {}) : testOnlyGrant(path, form);

  test('a refresh with no reply stops at the timeout and releases the lock', async () => {
    const { grant } = await signIn(testOnlyGrant);
    grant.tokens = {};
    await writeRaw({ version: 2, grants: [grant] });
    handler = noRefreshReply;
    const timeout = AbortSignal.timeout.bind(AbortSignal);
    const spy = spyOn(AbortSignal, 'timeout').mockImplementation(() => timeout(50));
    try {
      await expect(oauth.oauthCredentials('test_mode', ep)).rejects.toMatchObject({ code: 'timeout' });
      expect(spy).toHaveBeenCalledWith(10_000);
    } finally {
      spy.mockRestore();
    }
    expect(existsSync(join(dirname(oauth.storePath()), 'refresh.lock'))).toBe(false);
  });

  test('Ctrl-C during a refresh exits with 130 and removes the lock', async () => {
    const childHome = mkdtempSync(join(tmpdir(), 'dodo-sigint-test-'));
    const configDir = join(childHome, '.dodopayments');
    mkdirSync(configDir, { mode: 0o700 });
    const lockPath = join(configDir, 'refresh.lock');
    const { writeEncrypted } = await import('./secure-store');
    const grant = {
      id: 'sigint',
      issuer: ep.issuer,
      refresh_token: 'rt_sigint',
      modes: ['test_mode'],
      scope: 'payments:read',
      device_name: 'tester@box',
      signed_in_at: new Date().toISOString(),
      tokens: {},
    };
    writeEncrypted(join(configDir, 'oauth.json'), JSON.stringify({ version: 2, grants: [grant] }));
    handler = noRefreshReply;

    const { DODO_PAYMENTS_API_KEY: _apiKey, ...env } = process.env;
    const child = Bun.spawn([process.execPath, 'src/index.ts', 'products', 'list'], {
      cwd: join(import.meta.dir, '..', '..'),
      env: {
        ...env,
        HOME: childHome,
        USERPROFILE: childHome,
        DODO_CONFIG_DIR: configDir,
        DODO_AUTH_ISSUER: ep.issuer,
        DODO_AUTH_RESOURCE_TEST: ep.resources.test_mode,
      },
      stdin: 'ignore',
      stdout: 'pipe',
      stderr: 'pipe',
    });
    try {
      while (!(calls.some((c) => c.form.get('refresh_token') === 'rt_sigint') && existsSync(lockPath))) {
        if (child.exitCode !== null) {
          const output = (await new Response(child.stdout).text()) + (await new Response(child.stderr).text());
          throw new Error(`dodo stopped before the refresh: ${output}`);
        }
        await Bun.sleep(50);
      }
      child.kill('SIGINT');
      expect(await child.exited).toBe(130);
      expect(existsSync(lockPath)).toBe(false);
    } finally {
      if (child.exitCode === null) child.kill();
      rmSync(childHome, { recursive: true, force: true });
    }
  }, 30_000);
});

describe('logout', () => {
  test('logout of one mode revokes only the grant of that mode', async () => {
    const test = (await signIn(testOnlyGrant)).grant;
    await signIn(liveOnlyGrant);
    const { removed, revokeErrors } = await oauth.oauthLogout('test_mode', ep);
    expect(removed.map((grant) => grant.id)).toEqual([test.id]);
    expect(revokeErrors).toEqual([]);
    expect(calls.map((c) => c.form.get('token'))).toEqual([test.refresh_token]);
    expect(oauth.oauthModes()).toEqual(['live_mode']);
  });

  test('logout all revokes each grant and deletes the store', async () => {
    await signIn(testOnlyGrant);
    await signIn(liveOnlyGrant);
    const { removed } = await oauth.oauthLogout('all', ep);
    expect(removed).toHaveLength(2);
    expect(calls.filter((c) => c.path === '/oauth2/revoke')).toHaveLength(2);
    expect(oauth.readGrants()).toEqual([]);
    expect((await oauth.oauthLogout('all', ep)).removed).toEqual([]);
  });

  test('reports a failed server revoke and still deletes the grant', async () => {
    await signIn(testOnlyGrant);
    handler = () => ({ status: 503, body: {} });
    const { removed, revokeErrors } = await oauth.oauthLogout('all', ep);
    expect(removed).toHaveLength(1);
    expect(revokeErrors).toEqual(['Test Mode: HTTP 503']);
    expect(oauth.readGrants()).toEqual([]);
  });

  test('logout all deletes a store that this computer cannot read', async () => {
    writeFileSync(oauth.storePath(), 'not encrypted');
    expect(oauth.readStore()).toEqual({ grants: [], unreadable: true });
    expect(await oauth.oauthLogout('all', ep)).toEqual({ removed: [], revokeErrors: [], unreadable: true });
    expect(existsSync(oauth.storePath())).toBe(false);
  });

  test('logout of one mode also deletes a store that this computer cannot read', async () => {
    writeForeignStore(oauth.storePath(), { version: 2, grants: [] });
    expect((await oauth.oauthLogout('live_mode', ep)).unreadable).toBe(true);
    expect(existsSync(oauth.storePath())).toBe(false);
  });
});

/** Encrypts `value` with the key of another computer, as after a change of the machine ID or the host name. */
function writeForeignStore(file: string, value: unknown) {
  const key = pbkdf2Sync('other-host-other-user', 'dodopayments-cli-salt', 100000, 32, 'sha256');
  const iv = randomBytes(16);
  const cipher = createCipheriv('aes-256-gcm', key, iv);
  const encrypted = cipher.update(JSON.stringify(value), 'utf8', 'hex') + cipher.final('hex');
  mkdirSync(dirname(file), { recursive: true, mode: 0o700 });
  writeFileSync(file, `${iv.toString('hex')}:${cipher.getAuthTag().toString('hex')}:${encrypted}`, { mode: 0o600 });
}

/** Runs the CLI in a child process. The child inherits DODO_CONFIG_DIR. */
async function runDodo(...args: string[]) {
  const { DODO_PAYMENTS_API_KEY: _apiKey, ...env } = process.env;
  const child = Bun.spawn([process.execPath, 'src/index.ts', ...args], {
    cwd: join(import.meta.dir, '..', '..'),
    env: { ...env, DODO_AUTH_ISSUER: ep.issuer, DODO_AUTH_RESOURCE_TEST: ep.resources.test_mode },
    stdin: 'ignore',
    stdout: 'pipe',
    stderr: 'pipe',
  });
  const [code, stdout, stderr] = await Promise.all([
    child.exited,
    new Response(child.stdout).text(),
    new Response(child.stderr).text(),
  ]);
  return { code, stdout, stderr };
}

describe('a store that this computer cannot read', () => {
  const foreignGrant = {
    id: 'foreign',
    issuer: 'https://auth.example',
    refresh_token: 'rt_foreign',
    modes: ['test_mode'],
    scope: 'payments:read',
    device_name: 'old@host',
    signed_in_at: new Date().toISOString(),
    tokens: {},
  };

  test('whoami explains it, keeps the file, and exits with 1', async () => {
    writeForeignStore(oauth.storePath(), { version: 2, grants: [foreignGrant] });
    expect(oauth.readStore()).toEqual({ grants: [], unreadable: true });

    const { code, stderr } = await runDodo('whoami');
    expect(code).toBe(1);
    expect(stderr).toContain(`The stored sign-in in ${oauth.storePath()} cannot be read on this machine.`);
    expect(stderr).toContain('The machine ID or the host name changed.');
    expect(stderr).toContain('revoke it in the dashboard (Profile, Connected apps)');
    expect(stderr).not.toContain('Not signed in');
    expect(existsSync(oauth.storePath())).toBe(true);
  }, 30_000);

  test('a command that needs credentials explains it and exits with 1', async () => {
    writeForeignStore(oauth.storePath(), { version: 2, grants: [foreignGrant] });
    const { code, stdout, stderr } = await runDodo('products', 'list');
    expect(code).toBe(1);
    expect(stderr).toContain('cannot be read on this machine');
    expect(stdout).not.toContain('Please login');
    expect(existsSync(oauth.storePath())).toBe(true);
  }, 30_000);

  test('a new sign-in replaces it', async () => {
    writeForeignStore(oauth.storePath(), { version: 2, grants: [foreignGrant] });
    const { grant, replaced } = await signIn(testOnlyGrant);
    expect(replaced).toEqual([]);
    expect(oauth.readStore()).toEqual({ grants: [grant], unreadable: false });
    expect(await unreadableStoreMessages()).toEqual([]);
  });

  test('logout deletes it and says that the server sign-in is not revoked', async () => {
    writeForeignStore(oauth.storePath(), { version: 2, grants: [foreignGrant] });
    const { code, stdout, stderr } = await runDodo('logout', 'all');
    expect(code).toBe(0);
    expect(stderr).toContain('The CLI did not revoke the sign-in on the server.');
    expect(stderr).toContain('revoke it in the dashboard (Profile, Connected apps)');
    expect(stdout).toContain('Logged out from all accounts.');
    expect(existsSync(oauth.storePath())).toBe(false);
    expect(calls.some((c) => c.path === '/oauth2/revoke')).toBe(false);
  }, 30_000);

  test('an API key file that this computer cannot read gets its own message', async () => {
    writeForeignStore(configPath(), { test_mode: 'sk_test' });
    try {
      const messages = await unreadableStoreMessages();
      expect(messages).toEqual([expect.stringContaining(`The stored API keys in ${configPath()} cannot be read on this machine.`)]);
      expect(existsSync(configPath())).toBe(true);
    } finally {
      rmSync(configPath(), { force: true });
    }
  });
});

describe('loopback flow', () => {
  test('puts the device name in the authorize URL', async () => {
    let opened = '';
    const noBrowser = async (url: string) => {
      opened = url;
      throw new Error('no browser');
    };
    await expect(oauth.loopbackSignIn(noBrowser, ep)).rejects.toThrow('no browser');
    const url = new URL(opened);
    expect(url.pathname).toBe('/oauth2/authorize');
    expect(url.searchParams.get('device_name')).toBe('tester@box');
  });
});

describe('flow choice', () => {
  test('uses the device flow over SSH and on Linux with no display', () => {
    expect(oauth.preferDeviceFlow({ SSH_CONNECTION: '1 2 3 4' }, 'darwin')).toBe(true);
    expect(oauth.preferDeviceFlow({}, 'linux')).toBe(true);
    expect(oauth.preferDeviceFlow({ DISPLAY: ':0' }, 'linux')).toBe(false);
    expect(oauth.preferDeviceFlow({}, 'darwin')).toBe(false);
  });
});
