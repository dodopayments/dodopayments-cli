import { afterAll, beforeAll, beforeEach, describe, expect, test } from 'bun:test';
import { mkdtempSync, rmSync } from 'node:fs';
import { tmpdir } from 'node:os';
import { join } from 'node:path';

const home = mkdtempSync(join(tmpdir(), 'dodo-oauth-test-'));
process.env.HOME = home;
process.env.USERPROFILE = home;

import type { OAuthEndpoints } from './oauth';

const oauth = await import('./oauth');

type Form = URLSearchParams;
let handler: (path: string, form: Form) => { status: number; body: unknown } = () => ({ status: 500, body: {} });
const calls: { path: string; form: Form }[] = [];

const server = Bun.serve({
  port: 0,
  async fetch(request) {
    const path = new URL(request.url).pathname;
    const form = new URLSearchParams(await request.text());
    calls.push({ path, form });
    const { status, body } = handler(path, form);
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
  rmSync(home, { recursive: true, force: true });
});

/** A grant with test mode only: the live resource gets `invalid_target`. */
function testOnlyGrant(path: string, form: Form) {
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
    if (form.get('resource') === ep.resources.live_mode) return { status: 400, body: { error: 'invalid_target' } };
    return { status: 200, body: token(String(form.get('resource'))) };
  }
  if (path === '/oauth2/revoke') return { status: 200, body: {} };
  return { status: 404, body: {} };
}

describe('device flow', () => {
  test('polls until the approval and stores the modes of the grant', async () => {
    handler = testOnlyGrant;
    const start = await oauth.startDeviceFlow(ep);
    expect(start.user_code).toBe('BCDF-GHJK');
    const authorize = calls[0]!.form;
    expect(authorize.get('client_id')).toBe('dodo_cli');
    expect(authorize.getAll('resource')).toEqual([ep.resources.test_mode, ep.resources.live_mode, ep.resources.mcp]);

    const store = await oauth.pollDeviceFlow(start, ep);
    expect(store.modes).toEqual(['test_mode']);
    expect(store.device_name).toBe('tester@box');
    expect(oauth.oauthModes()).toEqual(['test_mode']);
    expect(await oauth.oauthCredentials('test_mode', ep)).toBe(store.tokens[ep.resources.test_mode]!.access_token);
  });

  test('stops on access_denied', async () => {
    handler = (path, form) =>
      path === '/oauth2/token' ? { status: 400, body: { error: 'access_denied' } } : testOnlyGrant(path, form);
    const start = await oauth.startDeviceFlow(ep);
    await expect(oauth.pollDeviceFlow(start, ep)).rejects.toMatchObject({ code: 'access_denied' });
    expect(oauth.readStore()).toBeNull();
  });
});

describe('refresh', () => {
  test('refreshes a token that is near the expiry and keeps the new refresh token', async () => {
    handler = testOnlyGrant;
    const store = await oauth.pollDeviceFlow(await oauth.startDeviceFlow(ep), ep);
    const cached = store.tokens[ep.resources.mcp]!.access_token;
    expect(await oauth.accessTokenFor(ep.resources.mcp, ep)).toBe(cached);

    // Expire the MCP token in the store.
    const { writeEncrypted } = await import('./secure-store');
    store.tokens[ep.resources.mcp]!.expires_at = Date.now();
    writeEncrypted(oauth.storePath(), JSON.stringify(store));

    calls.length = 0;
    const next = await oauth.accessTokenFor(ep.resources.mcp, ep);
    expect(next).not.toBe(cached);
    expect(calls[0]!.form.get('refresh_token')).toBe(store.refresh_token);
    expect(oauth.readStore()!.refresh_token).not.toBe(store.refresh_token);
  });

  test('deletes the sign-in on invalid_grant', async () => {
    handler = testOnlyGrant;
    const store = await oauth.pollDeviceFlow(await oauth.startDeviceFlow(ep), ep);
    handler = () => ({ status: 400, body: { error: 'invalid_grant' } });
    const { writeEncrypted } = await import('./secure-store');
    store.tokens = {};
    writeEncrypted(oauth.storePath(), JSON.stringify(store));
    await expect(oauth.accessTokenFor(ep.resources.test_mode, ep)).rejects.toMatchObject({ code: 'invalid_grant' });
    expect(oauth.readStore()).toBeNull();
  });
});

describe('logout', () => {
  test('revokes the refresh token and deletes the store', async () => {
    handler = testOnlyGrant;
    const store = await oauth.pollDeviceFlow(await oauth.startDeviceFlow(ep), ep);
    calls.length = 0;
    expect(await oauth.oauthLogout(ep)).toBe(true);
    expect(calls[0]!.path).toBe('/oauth2/revoke');
    expect(calls[0]!.form.get('token')).toBe(store.refresh_token);
    expect(oauth.readStore()).toBeNull();
    expect(await oauth.oauthLogout(ep)).toBe(false);
  });

  test('reports a failed server revoke and still deletes the store', async () => {
    handler = testOnlyGrant;
    await oauth.pollDeviceFlow(await oauth.startDeviceFlow(ep), ep);
    handler = () => ({ status: 503, body: {} });
    await expect(oauth.oauthLogout(ep)).rejects.toMatchObject({ code: 'server_error' });
    expect(oauth.readStore()).toBeNull();
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
