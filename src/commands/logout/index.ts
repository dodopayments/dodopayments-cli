import { clearConfig, clearSessionMode, getSessionMode, signedInModes, type Mode } from '../../utils/auth';
import { oauthLogout } from '../../utils/oauth';
import type { CommandContext } from '../../tui/CommandContext';

type LogoutChoice = 'all' | Mode;

const modeLabels: Record<Mode, string> = {
  test_mode: 'Test Mode',
  live_mode: 'Live Mode',
};

const USAGE = 'Usage: dodo logout <test|live|all>';

function normalizeTarget(value: string): LogoutChoice | null {
  if (value === 'all') return 'all';
  if (value === 'test' || value === 'test_mode') return 'test_mode';
  if (value === 'live' || value === 'live_mode') return 'live_mode';
  return null;
}

/**
 * Signs out of the API keys and the browser sign-ins of `target`, and revokes the browser sign-ins
 * on the server. In the TUI, the TUI closes when no account is left.
 */
export async function handleLogout(
  ctx: CommandContext,
  exit?: () => void,
  targetArg?: string,
): Promise<void> {
  const exitCli = exit ?? (() => process.exit(0));

  let target: LogoutChoice;

  if (ctx.invocation === 'cli') {
    if (!targetArg) {
      ctx.addBlock({ type: 'error', message: 'Target is required.' });
      ctx.addBlock({ type: 'info', message: USAGE });
      return;
    }
    const normalized = normalizeTarget(targetArg);
    if (!normalized) {
      ctx.addBlock({ type: 'error', message: `Invalid target '${targetArg}'. Use 'test', 'live', or 'all'.` });
      ctx.addBlock({ type: 'info', message: USAGE });
      return;
    }
    target = normalized;
  } else {
    const modes = await signedInModes();
    if (modes.length === 0) {
      ctx.addBlock({ type: 'error', message: 'Not signed in.' });
      return;
    }
    target = (await ctx.promptSelect('Sign out from', [
      { label: 'All accounts', value: 'all' },
      ...modes.map((mode) => ({ label: modeLabels[mode], value: mode })),
    ])) as LogoutChoice;

    const targetLabel = target === 'all' ? 'all accounts' : modeLabels[target];
    const confirmed = await ctx.promptConfirm(`Sign out from ${targetLabel}?`);
    if (!confirmed) {
      ctx.addBlock({ type: 'error', message: 'Logout cancelled.' });
      return;
    }
  }

  const apiKeys = await clearConfig(target);
  const oauth = await oauthLogout(target);

  for (const error of oauth.revokeErrors) {
    ctx.addBlock({
      type: 'error',
      message: `The server did not confirm the revoke (${error}). The sign-in is deleted from this computer. Revoke it in Settings > Connected Apps.`,
    });
  }
  if (apiKeys.hadInvalidConfig) {
    ctx.addBlock({ type: 'error', message: 'The stored API keys could not be read. They are deleted.' });
  }

  const oauthModes = oauth.removed.flatMap((grant) => grant.modes);
  const removed = [...new Set([...apiKeys.removedModes, ...oauthModes])];
  if (target !== 'all') {
    // A grant from v3.7 can hold both modes. A revoke ends it for both modes.
    const alsoRemoved = oauthModes.filter((mode) => mode !== target);
    if (alsoRemoved.length > 0) {
      ctx.addBlock({
        type: 'info',
        message: `This browser sign-in also held ${modeLabels[alsoRemoved[0]!]}. It is signed out too.`,
      });
    }
  }

  const active = getSessionMode();
  if (active && removed.includes(active)) clearSessionMode();

  if (removed.length === 0 && !apiKeys.hadInvalidConfig) {
    ctx.addBlock({
      type: 'error',
      message: target === 'all' ? 'No stored accounts found.' : `No ${modeLabels[target]} account is signed in.`,
    });
  } else {
    ctx.addBlock({
      type: 'success',
      message: target === 'all' ? 'Logged out from all accounts.' : `Logged out from ${modeLabels[target]}.`,
    });
  }

  if (ctx.invocation === 'cli' || (await signedInModes()).length === 0) exitCli();
}
