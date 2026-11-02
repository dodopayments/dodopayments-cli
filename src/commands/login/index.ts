import DodoPayments from 'dodopayments';
import open from 'open';
import { saveConfig, setSessionMode, type Mode } from '../../utils/auth';
import {
  loopbackSignIn,
  OAuthError,
  pollDeviceFlow,
  preferDeviceFlow,
  startDeviceFlow,
  type SignInResult,
} from '../../utils/oauth';
import type { CommandContext } from '../../tui/CommandContext';

const USAGE =
  'Usage: dodo login [--device]\n' +
  '       dodo login <api-key> <test|live>';

export interface LoginOptions {
  /** Use the device flow, also when a browser can open. */
  device?: boolean;
}

function normalizeMode(value: string): Mode | null {
  if (value === 'test' || value === 'test_mode') return 'test_mode';
  if (value === 'live' || value === 'live_mode') return 'live_mode';
  return null;
}

const modeLabel = (mode: Mode) => (mode === 'test_mode' ? 'Test Mode' : 'Live Mode');

async function apiKeyLogin(ctx: CommandContext, apiKey: string, mode: Mode): Promise<boolean> {
  const client = new DodoPayments({ bearerToken: apiKey, environment: mode });

  const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Verifying API key…' });
  try {
    await client.products.list({ page_size: 1 });
    ctx.removeBlock(spinnerId);
    ctx.addBlock({ type: 'success', message: 'API key verified.' });
  } catch {
    ctx.removeBlock(spinnerId);
    ctx.addBlock({
      type: 'error',
      message: 'Authentication failed. Check your API key and selected environment.',
    });
    process.exitCode = 1;
    return false;
  }

  await saveConfig(mode, apiKey);
  ctx.addBlock({ type: 'success', message: 'Setup complete.' });
  return true;
}

async function deviceLogin(ctx: CommandContext): Promise<SignInResult> {
  const start = await startDeviceFlow();
  const url = start.verification_uri_complete ?? start.verification_uri;
  ctx.addBlock({
    type: 'info',
    message:
      `Open ${start.verification_uri} on any device and type the code ${start.user_code}.\n` +
      `Or open: ${url}\n` +
      'Type the code only if you started this sign-in yourself, just now.',
  });
  const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Waiting for the approval…' });
  try {
    return await pollDeviceFlow(start);
  } finally {
    ctx.removeBlock(spinnerId);
  }
}

async function browserLogin(ctx: CommandContext): Promise<SignInResult> {
  const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Waiting for the approval in your browser…' });
  try {
    return await loopbackSignIn(async (url) => {
      ctx.addBlock({ type: 'info', message: `If the browser does not open, go to:\n${url}` });
      await open(url);
    });
  } finally {
    ctx.removeBlock(spinnerId);
  }
}

async function oauthLogin(ctx: CommandContext, options: LoginOptions): Promise<boolean> {
  let result: SignInResult;
  try {
    if (options.device || preferDeviceFlow()) {
      result = await deviceLogin(ctx);
    } else {
      try {
        result = await browserLogin(ctx);
      } catch (error) {
        if (error instanceof OAuthError) throw error;
        // No browser could open. The device flow works from any device.
        result = await deviceLogin(ctx);
      }
    }
  } catch (error) {
    const message = error instanceof Error ? error.message : String(error);
    ctx.addBlock({ type: 'error', message: `Sign-in failed. ${message}` });
    process.exitCode = 1;
    return false;
  }

  const { grant, replaced, revokeErrors } = result;
  const modes = grant.modes.map(modeLabel).join(', ');
  if (replaced.length > 0) {
    ctx.addBlock({ type: 'info', message: `This sign-in replaces the earlier ${modes} sign-in. The CLI revoked the earlier one.` });
  }
  for (const error of revokeErrors) {
    ctx.addBlock({
      type: 'error',
      message: `The server did not confirm the revoke of the earlier sign-in (${error}). Revoke it in Settings > Connected Apps.`,
    });
  }
  setSessionMode(grant.modes[0]!);
  ctx.addBlock({ type: 'success', message: `Signed in to ${modes} as ${grant.device_name}.` });
  return true;
}

export async function handleLogin(
  ctx: CommandContext,
  apiKeyArg?: string,
  modeArg?: string,
  options: LoginOptions = {},
): Promise<boolean> {
  if (ctx.invocation === 'cli') {
    if (apiKeyArg || modeArg) {
      const normalized = modeArg ? normalizeMode(modeArg) : null;
      if (!apiKeyArg || !normalized) {
        ctx.addBlock({ type: 'error', message: modeArg ? `Invalid mode '${modeArg}'. Use 'test' or 'live'.` : 'API key and mode are required.' });
        ctx.addBlock({ type: 'info', message: USAGE });
        process.exitCode = 1;
        return false;
      }
      return apiKeyLogin(ctx, apiKeyArg, normalized);
    }
    if (!options.device && !process.stdout.isTTY) {
      ctx.addBlock({
        type: 'error',
        message: 'No terminal for the sign-in. Set DODO_PAYMENTS_API_KEY, or run `dodo login --device`.',
      });
      ctx.addBlock({ type: 'info', message: USAGE });
      return false;
    }
    return oauthLogin(ctx, options);
  }

  const method = await ctx.promptSelect('Sign in with', [
    { label: 'Browser (recommended)', value: 'browser' },
    { label: 'Code on another device', value: 'device' },
    { label: 'API key', value: 'api_key' },
  ]);
  if (method !== 'api_key') return oauthLogin(ctx, { device: method === 'device' });

  await open('https://app.dodopayments.com/developer/api-keys');
  const apiKey = await ctx.promptInput('API key');
  const mode = (await ctx.promptSelect('Environment', [
    { label: 'Test Mode', value: 'test_mode' },
    { label: 'Live Mode', value: 'live_mode' },
  ])) as Mode;
  return apiKeyLogin(ctx, apiKey, mode);
}
