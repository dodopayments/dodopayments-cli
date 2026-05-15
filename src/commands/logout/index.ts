import { clearConfig } from '../../utils/auth';
import type { CommandContext } from '../../tui/CommandContext';

type LogoutChoice = 'all' | 'test_mode' | 'live_mode';

const modeLabels: Record<Exclude<LogoutChoice, 'all'>, string> = {
  test_mode: 'Test Mode',
  live_mode: 'Live Mode',
};

export async function handleLogout(ctx: CommandContext, exit?: () => void): Promise<void> {
  const exitCli = exit ?? (() => process.exit(0));

  const target = await ctx.promptSelect('Sign out from', [
    { label: 'All accounts', value: 'all' },
    { label: 'Test Mode', value: 'test_mode' },
    { label: 'Live Mode', value: 'live_mode' },
  ]) as LogoutChoice;

  const targetLabel = target === 'all' ? 'all accounts' : modeLabels[target];
  const confirmed = await ctx.promptConfirm(`Sign out from ${targetLabel}?`);
  if (!confirmed) {
    ctx.addBlock({ type: 'error', message: 'Logout cancelled.' });
    return;
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
