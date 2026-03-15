import { select } from '@inquirer/prompts';
import { clearConfig } from '../../utils/auth';

type LogoutChoice = 'all' | 'test_mode' | 'live_mode';

const modeLabels: Record<Exclude<LogoutChoice, 'all'>, string> = {
  test_mode: 'Test Mode',
  live_mode: 'Live Mode',
};

export async function handleLogout(): Promise<void> {
  const target = await select<LogoutChoice>({
    message: 'Which account would you like to logout from?',
    choices: [
      { name: 'All accounts', value: 'all' },
      { name: 'Test Mode', value: 'test_mode' },
      { name: 'Live Mode', value: 'live_mode' },
    ],
  });

  const result = clearConfig(target);

  if (result.hadInvalidConfig) {
    console.log('Stored credentials were invalid and have been cleared.');
    process.exit(0);
  }

  if (target === 'all') {
    if (result.removedModes.length === 0) {
      console.log('No stored accounts were found.');
      process.exit(0);
    }

    console.log('Logged out from all stored accounts.');
    process.exit(0);
  }

  if (result.removedModes.length === 0) {
    console.log(`No ${modeLabels[target]} account is currently logged in.`);
    process.exit(0);
  }

  console.log(`Logged out from ${modeLabels[target]}.`);
  process.exit(0);
}
