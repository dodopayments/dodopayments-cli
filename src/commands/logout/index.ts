import { clearConfig } from '../../utils/auth';
import type { CommandContext } from '../../ui/ink/CommandContext';

type LogoutChoice = 'all' | 'test_mode' | 'live_mode';

const modeLabels: Record<Exclude<LogoutChoice, 'all'>, string> = {
  test_mode: 'Test Mode',
  live_mode: 'Live Mode',
};

export async function handleLogout(ctx: CommandContext): Promise<void> {
  const target = await ctx.promptSelect('Which account would you like to logout from?', [
    { label: 'All accounts', value: 'all' },
    { label: 'Test Mode', value: 'test_mode' },
    { label: 'Live Mode', value: 'live_mode' },
  ]) as LogoutChoice;

  const confirmed = await ctx.promptConfirm(`Are you sure you want to logout from ${target}?`);
  if (!confirmed) {
    ctx.addBlock({ type: 'error', message: 'Logout cancelled.' });
    return;
  }

  const result = await clearConfig(target);

  if (result.hadInvalidConfig) {
    ctx.addBlock({ type: 'error', message: 'Stored credentials were invalid and have been cleared.' });
    process.exit(0);
    return;
  }

  if (target === 'all') {
    if (result.removedModes.length === 0) {
      ctx.addBlock({ type: 'error', message: 'No stored accounts were found.' });
      process.exit(0);
      return;
    }

    ctx.addBlock({ type: 'success', message: 'Logged out from all stored accounts.' });
    process.exit(0);
    return;
  }

  if (result.removedModes.length === 0) {
    ctx.addBlock({ type: 'error', message: `No ${modeLabels[target]} account is currently logged in.` });
    process.exit(0);
    return;
  }

  ctx.addBlock({ type: 'success', message: `Logged out from ${modeLabels[target]}.` });
  process.exit(0);
}
