import { clearConfig } from '../../utils/auth';
import type { CommandContext } from '../../tui/CommandContext';

type LogoutChoice = 'all' | 'test_mode' | 'live_mode';

const modeLabels: Record<Exclude<LogoutChoice, 'all'>, string> = {
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
    target = (await ctx.promptSelect('Sign out from', [
      { label: 'All accounts', value: 'all' },
      { label: 'Test Mode', value: 'test_mode' },
      { label: 'Live Mode', value: 'live_mode' },
    ])) as LogoutChoice;

    const targetLabel = target === 'all' ? 'all accounts' : modeLabels[target];
    const confirmed = await ctx.promptConfirm(`Sign out from ${targetLabel}?`);
    if (!confirmed) {
      ctx.addBlock({ type: 'error', message: 'Logout cancelled.' });
      return;
    }
  }

  const result = await clearConfig(target);

  if (result.hadInvalidConfig) {
    ctx.addBlock({ type: 'error', message: 'Stored credentials were invalid and have been cleared.' });
    exitCli();
    return;
  }

  if (target === 'all') {
    if (result.removedModes.length === 0) {
      ctx.addBlock({ type: 'error', message: 'No stored accounts found.' });
      exitCli();
      return;
    }

    ctx.addBlock({ type: 'success', message: 'Logged out from all accounts.' });
    exitCli();
    return;
  }

  if (result.removedModes.length === 0) {
    ctx.addBlock({ type: 'error', message: `No ${modeLabels[target]} account is signed in.` });
    exitCli();
    return;
  }

  ctx.addBlock({ type: 'success', message: `Logged out from ${modeLabels[target]}.` });
  exitCli();
}
