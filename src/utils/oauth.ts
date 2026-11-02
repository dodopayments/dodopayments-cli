import fs from 'node:fs';
import os from 'node:os';
import path from 'node:path';
import crypto from 'node:crypto';
import http from 'node:http';
import { version } from '../../package.json';
import { CONFIG_DIR, decrypt, ensureConfigDir, writeEncrypted } from './secure-store';

export type Mode = 'test_mode' | 'live_mode';

const CLIENT_ID = 'dodo_cli';
const DEVICE_GRANT = 'urn:ietf:params:oauth:grant-type:device_code';
const STORE_PATH = path.join(CONFIG_DIR, 'oauth.json');
const LOCK_PATH = path.join(CONFIG_DIR, 'refresh.lock');
const REFRESH_MARGIN_MS = 60_000;
const REFRESH_TIMEOUT_MS = 10_000;
/** Two times the refresh timeout. A live process releases the lock before this age. */
const LOCK_STALE_MS = 20_000;
/** More than the stale age. A waiter can then remove the lock of a killed process and continue. */
const LOCK_WAIT_MS = 25_000;
const LOOPBACK_TIMEOUT_MS = 5 * 60_000;
const ALL_MODES: Mode[] = ['test_mode', 'live_mode'];

/** The CLI asks for the Developer preset. The consent page can make it smaller. */
const REQUESTED_SCOPES = [
  'payments', 'refunds', 'disputes', 'subscriptions', 'products', 'customers', 'checkout_sessions',
  'discounts', 'license_keys', 'webhooks', 'events', 'meters', 'credit_entitlements', 'entitlements', 'brands',
]
  .flatMap((name) => [`${name}:read`, `${name}:write`])
  .concat(['payouts:read', 'balances:read'])
  .join(' ');

export interface OAuthEndpoints {
  issuer: string;
  resources: Record<Mode, string> & { mcp: string };
}

export function endpoints(env: NodeJS.ProcessEnv = process.env): OAuthEndpoints {
  return {
    issuer: (env.DODO_AUTH_ISSUER ?? 'https://auth.dodopayments.com').replace(/\/$/, ''),
    resources: {
      test_mode: env.DODO_AUTH_RESOURCE_TEST ?? 'https://test.dodopayments.com',
      live_mode: env.DODO_AUTH_RESOURCE_LIVE ?? 'https://live.dodopayments.com',
      mcp: env.DODO_AUTH_RESOURCE_MCP ?? 'https://mcp.dodopayments.com/mcp',
    },
  };
}

interface CachedToken {
  access_token: string;
  expires_at: number;
}

/** One approval on the consent page. A grant has one mode. A grant from v3.7 can have both modes. */
export interface OAuthStore {
  id: string;
  issuer: string;
  refresh_token: string;
  modes: Mode[];
  scope: string;
  device_name: string;
  signed_in_at: string;
  /** The access tokens, by resource. */
  tokens: Record<string, CachedToken>;
}

interface TokenResponse {
  access_token?: string;
  refresh_token?: string;
  expires_in?: number;
  scope?: string;
  error?: string;
  error_description?: string;
  interval?: number;
}

export class OAuthError extends Error {
  constructor(
    readonly code: string,
    message: string,
  ) {
    super(message);
    this.name = 'OAuthError';
  }
}

export interface SignInResult {
  grant: OAuthStore;
  /** The earlier grants of the same modes. The new grant replaces them, and the CLI revokes them. */
  replaced: OAuthStore[];
  /** The revokes that the server did not confirm. */
  revokeErrors: string[];
}

const isGrant = (value: Partial<OAuthStore> | undefined): value is OAuthStore =>
  Boolean(value?.refresh_token) && Array.isArray(value?.modes) && value.modes.length > 0;

/**
 * The stored grants. v3.7 stored one grant with no `grants` list and no `id`. A store that this
 * computer cannot decrypt counts as no sign-in. The next sign-in or sign-out writes it again.
 */
export function readGrants(): OAuthStore[] {
  if (!fs.existsSync(STORE_PATH)) return [];
  let data: { grants?: Partial<OAuthStore>[] } & Partial<OAuthStore>;
  try {
    data = JSON.parse(decrypt(fs.readFileSync(STORE_PATH, 'utf-8')));
  } catch {
    return [];
  }
  const list = Array.isArray(data.grants) ? data.grants : [data];
  return list.filter(isGrant).map((grant) => ({ ...grant, id: grant.id ?? 'v3.7', tokens: grant.tokens ?? {} }));
}

function writeGrants(grants: OAuthStore[]) {
  if (grants.length === 0) {
    deleteStore();
    return;
  }
  writeEncrypted(STORE_PATH, JSON.stringify({ version: 2, grants }));
}

/** The grant that holds `mode`, or null. */
export function grantFor(mode: Mode): OAuthStore | null {
  return readGrants().find((grant) => grant.modes.includes(mode)) ?? null;
}

export function deleteStore() {
  fs.rmSync(STORE_PATH, { force: true });
}

export function storePath(): string {
  return STORE_PATH;
}

/** The modes of the stored grants. An empty list means no OAuth sign-in. */
export function oauthModes(): Mode[] {
  const grants = readGrants();
  return ALL_MODES.filter((mode) => grants.some((grant) => grant.modes.includes(mode)));
}

export function deviceName(env: NodeJS.ProcessEnv = process.env): string {
  return env.DODO_DEVICE_NAME ?? `${os.userInfo().username}@${os.hostname()}`;
}

/** The sign-in server shows this value on the device confirm page. */
const USER_AGENT = `dodopayments-cli/${version} (${process.platform} ${process.arch})`;

async function postForm(
  url: string,
  fields: Record<string, string | string[]>,
  signal?: AbortSignal,
): Promise<{ status: number; body: TokenResponse }> {
  const params = new URLSearchParams();
  for (const [key, value] of Object.entries(fields)) {
    for (const item of Array.isArray(value) ? value : [value]) params.append(key, item);
  }
  const response = await fetch(url, {
    method: 'POST',
    headers: {
      'content-type': 'application/x-www-form-urlencoded',
      accept: 'application/json',
      'user-agent': USER_AGENT,
    },
    body: params,
    signal,
  });
  const text = await response.text();
  let body: TokenResponse = {};
  try {
    body = text ? (JSON.parse(text) as TokenResponse) : {};
  } catch {
    body = { error: 'invalid_response', error_description: `HTTP ${response.status}` };
  }
  return { status: response.status, body };
}

function tokenOf(body: TokenResponse, now: number): CachedToken {
  if (!body.access_token) throw new OAuthError('invalid_response', 'The sign-in server returned no access token.');
  return { access_token: body.access_token, expires_at: now + (body.expires_in ?? 3600) * 1000 };
}

async function refreshFor(
  ep: OAuthEndpoints,
  refreshToken: string,
  resource: string,
): Promise<{ status: number; body: TokenResponse }> {
  const signal = AbortSignal.timeout(REFRESH_TIMEOUT_MS);
  try {
    return await postForm(
      `${ep.issuer}/oauth2/token`,
      { grant_type: 'refresh_token', refresh_token: refreshToken, client_id: CLIENT_ID, resource },
      signal,
    );
  } catch (error) {
    if (signal.aborted) {
      throw new OAuthError('timeout', `The sign-in server did not reply in ${REFRESH_TIMEOUT_MS / 1000} seconds. Try again.`);
    }
    throw error;
  }
}

/**
 * Makes the store from the first token response. The grant modes are not in the response, so the
 * CLI asks for the API token of each mode. The grant refuses a mode that the user did not approve.
 */
async function finishSignIn(ep: OAuthEndpoints, first: TokenResponse): Promise<SignInResult> {
  const now = Date.now();
  if (!first.refresh_token) throw new OAuthError('invalid_response', 'The sign-in server returned no refresh token.');
  const store: OAuthStore = {
    id: crypto.randomUUID(),
    issuer: ep.issuer,
    refresh_token: first.refresh_token,
    modes: [],
    scope: first.scope ?? '',
    device_name: deviceName(),
    signed_in_at: new Date(now).toISOString(),
    tokens: { [ep.resources.mcp]: tokenOf(first, now) },
  };
  for (const mode of ALL_MODES) {
    const { status, body } = await refreshFor(ep, store.refresh_token, ep.resources[mode]);
    if (status === 200) {
      store.modes.push(mode);
      store.tokens[ep.resources[mode]] = tokenOf(body, Date.now());
      if (body.refresh_token) store.refresh_token = body.refresh_token;
    } else if (body.error !== 'invalid_target') {
      throw new OAuthError(body.error ?? 'server_error', body.error_description ?? `HTTP ${status}`);
    }
  }
  if (store.modes.length === 0) throw new OAuthError('invalid_grant', 'The approval has no test mode or live mode.');
  const replaced = await withRefreshLock(async () => {
    const grants = readGrants();
    const old = grants.filter((grant) => grant.modes.some((mode) => store.modes.includes(mode)));
    writeGrants([...grants.filter((grant) => !old.includes(grant)), store]);
    return old;
  });
  return { grant: store, replaced, revokeErrors: await revokeAll(replaced, ep) };
}

function sleep(ms: number) {
  return new Promise((resolve) => setTimeout(resolve, ms));
}

export interface DeviceStart {
  device_code: string;
  user_code: string;
  verification_uri: string;
  verification_uri_complete?: string;
  expires_in: number;
  interval?: number;
}

function authorizationResources(ep: OAuthEndpoints): string[] {
  return [ep.resources.test_mode, ep.resources.live_mode, ep.resources.mcp];
}

export async function startDeviceFlow(ep: OAuthEndpoints = endpoints()): Promise<DeviceStart> {
  const { status, body } = await postForm(`${ep.issuer}/oauth2/device_authorization`, {
    client_id: CLIENT_ID,
    scope: REQUESTED_SCOPES,
    resource: authorizationResources(ep),
    device_name: deviceName(),
  });
  const start = body as unknown as Partial<DeviceStart> & TokenResponse;
  if (status !== 200 || !start.device_code || !start.user_code || !start.verification_uri) {
    throw new OAuthError(start.error ?? 'server_error', start.error_description ?? `HTTP ${status}`);
  }
  return start as DeviceStart;
}

/** Polls the token endpoint until the user approves, denies, or the code expires. */
export async function pollDeviceFlow(start: DeviceStart, ep: OAuthEndpoints = endpoints()): Promise<SignInResult> {
  let interval = (start.interval ?? 5) * 1000;
  const deadline = Date.now() + start.expires_in * 1000;
  while (Date.now() < deadline) {
    await sleep(interval);
    const { status, body } = await postForm(`${ep.issuer}/oauth2/token`, {
      grant_type: DEVICE_GRANT,
      device_code: start.device_code,
      client_id: CLIENT_ID,
      resource: ep.resources.mcp,
    });
    if (status === 200) return finishSignIn(ep, body);
    if (body.error === 'authorization_pending') continue;
    if (body.error === 'slow_down') {
      interval += 5000;
      continue;
    }
    if (body.error === 'access_denied') throw new OAuthError('access_denied', 'The sign-in was cancelled.');
    if (body.error === 'expired_token') break;
    throw new OAuthError(body.error ?? 'server_error', body.error_description ?? `HTTP ${status}`);
  }
  throw new OAuthError('expired_token', 'The code expired. Run dodo login again.');
}

function base64url(buffer: Buffer): string {
  return buffer.toString('base64url');
}

const CALLBACK_PAGE = (title: string, text: string) =>
  `<!DOCTYPE html><html><head><meta charset="utf-8"><title>${title}</title></head>` +
  `<body style="font-family:sans-serif;max-width:420px;margin:80px auto"><h1>${title}</h1><p>${text}</p></body></html>`;

/**
 * Sign-in with a browser on this computer (RFC 8252, loopback redirect with PKCE). `openUrl` opens
 * the browser. It must throw when no browser can open, so that the caller can use the device flow.
 */
export async function loopbackSignIn(
  openUrl: (url: string) => Promise<void>,
  ep: OAuthEndpoints = endpoints(),
): Promise<SignInResult> {
  const verifier = base64url(crypto.randomBytes(32));
  const challenge = base64url(crypto.createHash('sha256').update(verifier).digest());
  const state = base64url(crypto.randomBytes(16));

  let server: http.Server | undefined;
  try {
    const result = await new Promise<{ code: string; redirectUri: string }>((resolve, reject) => {
      const timer = setTimeout(
        () => reject(new OAuthError('timeout', 'The sign-in took too long. Run dodo login again.')),
        LOOPBACK_TIMEOUT_MS,
      );
      let redirectUri = '';
      server = http.createServer((req, res) => {
        const url = new URL(req.url ?? '/', 'http://127.0.0.1');
        if (url.pathname !== '/callback') {
          res.writeHead(404).end();
          return;
        }
        res.setHeader('content-type', 'text/html; charset=utf-8');
        const error = url.searchParams.get('error');
        const code = url.searchParams.get('code');
        const iss = url.searchParams.get('iss');
        if (url.searchParams.get('state') !== state || (iss !== null && iss !== ep.issuer)) {
          res.writeHead(400).end(CALLBACK_PAGE('Sign-in failed', 'The response does not match this sign-in.'));
          return;
        }
        clearTimeout(timer);
        if (error || !code) {
          res.writeHead(200).end(CALLBACK_PAGE('Sign-in cancelled', 'You can close this page.'));
          reject(new OAuthError(error ?? 'access_denied', 'The sign-in was cancelled.'));
          return;
        }
        res.writeHead(200).end(CALLBACK_PAGE('The CLI is connected', 'Go back to your terminal. You can close this page.'));
        resolve({ code, redirectUri });
      });
      server.on('error', reject);
      server.listen(0, '127.0.0.1', () => {
        const { port } = server!.address() as { port: number };
        redirectUri = `http://127.0.0.1:${port}/callback`;
        const params = new URLSearchParams({
          response_type: 'code',
          client_id: CLIENT_ID,
          redirect_uri: redirectUri,
          scope: REQUESTED_SCOPES,
          state,
          code_challenge: challenge,
          code_challenge_method: 'S256',
          device_name: deviceName(),
        });
        for (const resource of authorizationResources(ep)) params.append('resource', resource);
        openUrl(`${ep.issuer}/oauth2/authorize?${params.toString()}`).catch((e: unknown) => {
          clearTimeout(timer);
          reject(e);
        });
      });
    });
    const { status, body } = await postForm(`${ep.issuer}/oauth2/token`, {
      grant_type: 'authorization_code',
      code: result.code,
      code_verifier: verifier,
      client_id: CLIENT_ID,
      redirect_uri: result.redirectUri,
      resource: ep.resources.mcp,
    });
    if (status !== 200) throw new OAuthError(body.error ?? 'server_error', body.error_description ?? `HTTP ${status}`);
    return await finishSignIn(ep, body);
  } finally {
    server?.close();
  }
}

/** The device flow is the default when no local browser can open (SSH, Linux with no display). */
export function preferDeviceFlow(env: NodeJS.ProcessEnv = process.env, platform = process.platform): boolean {
  if (env.SSH_CONNECTION || env.SSH_TTY) return true;
  return platform === 'linux' && !env.DISPLAY && !env.WAYLAND_DISPLAY;
}

/** The lock file that this process holds. */
let heldLock: string | undefined;

// process.exit() does not run `finally` blocks. Ctrl-C calls process.exit(), so the lock must also go here.
process.on('exit', () => {
  if (heldLock === undefined) return;
  try {
    fs.unlinkSync(heldLock);
  } catch (error) {
    if ((error as NodeJS.ErrnoException).code !== 'ENOENT') throw error;
  }
});

async function withRefreshLock<T>(work: () => Promise<T>): Promise<T> {
  ensureConfigDir();
  const deadline = Date.now() + LOCK_WAIT_MS;
  let fd: number | undefined;
  while (fd === undefined) {
    try {
      fd = fs.openSync(LOCK_PATH, 'wx', 0o600);
    } catch (error) {
      if ((error as NodeJS.ErrnoException).code !== 'EEXIST') throw error;
      try {
        if (Date.now() - fs.statSync(LOCK_PATH).mtimeMs > LOCK_STALE_MS) fs.rmSync(LOCK_PATH, { force: true });
      } catch {
        // The other process removed the lock.
      }
      if (Date.now() > deadline) throw new OAuthError('lock_timeout', `Another dodo process holds ${LOCK_PATH}.`);
      await sleep(100);
    }
  }
  heldLock = LOCK_PATH;
  try {
    return await work();
  } finally {
    fs.closeSync(fd);
    fs.rmSync(LOCK_PATH, { force: true });
    heldLock = undefined;
  }
}

function fresh(token: CachedToken | undefined, now: number): token is CachedToken {
  return token !== undefined && token.expires_at - REFRESH_MARGIN_MS > now;
}

const modeLabel = (mode: Mode) => (mode === 'test_mode' ? 'Test Mode' : 'Live Mode');

/** An access token for one resource of the grant of `mode`. It refreshes 60 seconds before the expiry. */
async function accessTokenFor(mode: Mode, resource: string, ep: OAuthEndpoints): Promise<string> {
  const notSignedIn = () => new OAuthError('not_signed_in', `No ${modeLabel(mode)} sign-in. Run dodo login.`);
  const cached = grantFor(mode);
  if (!cached) throw notSignedIn();
  if (fresh(cached.tokens[resource], Date.now())) return cached.tokens[resource].access_token;

  return withRefreshLock(async () => {
    // Another process can refresh while this one waits for the lock.
    const grants = readGrants();
    const grant = grants.find(({ id }) => id === cached.id) ?? grants.find(({ modes }) => modes.includes(mode));
    if (!grant) throw notSignedIn();
    if (fresh(grant.tokens[resource], Date.now())) return grant.tokens[resource].access_token;
    const { status, body } = await refreshFor(ep, grant.refresh_token, resource);
    if (status !== 200) {
      if (body.error === 'invalid_grant') {
        writeGrants(grants.filter((other) => other !== grant));
        throw new OAuthError('invalid_grant', `Your ${modeLabel(mode)} sign-in ended or was revoked. Run dodo login again.`);
      }
      throw new OAuthError(body.error ?? 'server_error', body.error_description ?? `HTTP ${status}`);
    }
    const token = tokenOf(body, Date.now());
    grant.tokens[resource] = token;
    if (body.refresh_token) grant.refresh_token = body.refresh_token;
    writeGrants(grants);
    return token.access_token;
  });
}

/** The API token of one mode. */
export async function oauthCredentials(mode: Mode, ep: OAuthEndpoints = endpoints()): Promise<string> {
  return accessTokenFor(mode, ep.resources[mode], ep);
}

/** The token for the remote Dodo MCP server, from the grant of `mode`. */
export async function mcpAccessToken(mode: Mode, ep: OAuthEndpoints = endpoints()): Promise<string> {
  return accessTokenFor(mode, ep.resources.mcp, ep);
}

async function revoke(grant: OAuthStore, ep: OAuthEndpoints): Promise<void> {
  const { status, body } = await postForm(`${grant.issuer || ep.issuer}/oauth2/revoke`, {
    token: grant.refresh_token,
    token_type_hint: 'refresh_token',
    client_id: CLIENT_ID,
  });
  if (status !== 200) throw new OAuthError(body.error ?? 'server_error', body.error_description ?? `HTTP ${status}`);
}

/** Revokes each grant on the server. Gives the errors, one for each revoke that failed. */
async function revokeAll(grants: OAuthStore[], ep: OAuthEndpoints): Promise<string[]> {
  const errors: string[] = [];
  for (const grant of grants) {
    try {
      await revoke(grant, ep);
    } catch (error) {
      errors.push(`${grant.modes.map(modeLabel).join(', ')}: ${error instanceof Error ? error.message : String(error)}`);
    }
  }
  return errors;
}

/**
 * Deletes the grants of `target` from this computer, then revokes them on the server. A v3.7 grant
 * with both modes ends for both modes.
 */
export async function oauthLogout(
  target: Mode | 'all',
  ep: OAuthEndpoints = endpoints(),
): Promise<{ removed: OAuthStore[]; revokeErrors: string[] }> {
  const removed = await withRefreshLock(async () => {
    const grants = readGrants();
    const matches = grants.filter((grant) => target === 'all' || grant.modes.includes(target));
    if (target === 'all') deleteStore();
    else if (matches.length > 0) writeGrants(grants.filter((grant) => !matches.includes(grant)));
    return matches;
  });
  return { removed, revokeErrors: await revokeAll(removed, ep) };
}
