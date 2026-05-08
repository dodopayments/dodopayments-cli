import type { CommandContext } from '../../ui/ink/CommandContext';
import { version as currentVersion } from '../../../package.json';
import {
  checkForUpdates,
  detectInstallMethod,
  runForegroundUpdate,
} from '../../utils/update';

export async function handleUpdate(ctx: CommandContext): Promise<void> {
  const checkSpinnerId = ctx.addBlock({
    type: 'spinner',
    label: 'Checking for updates…',
  });

  const info = await checkForUpdates(currentVersion);
  ctx.removeBlock(checkSpinnerId);

  if (!info) {
    ctx.addBlock({
      type: 'success',
      message: `Already on the latest version (v${currentVersion}).`,
    });
    return;
  }

  const method = detectInstallMethod();

  if (method === 'binary' || method === 'unknown') {
    ctx.addBlock({
      type: 'info',
      message:
        `v${info.latestVersion} is available (you're on v${info.currentVersion}). ` +
        `This install can't be upgraded automatically — download the latest release from ` +
        `https://github.com/dodopayments/dodopayments-cli/releases.`,
    });
    return;
  }

  const installLabel = method === 'bun' ? 'bun add -g' : 'npm install -g';
  const updateSpinnerId = ctx.addBlock({
    type: 'spinner',
    label: `Updating to v${info.latestVersion} via ${installLabel}…`,
  });

  const result = runForegroundUpdate(method);
  ctx.removeBlock(updateSpinnerId);

  if (result.ok) {
    ctx.addBlock({
      type: 'success',
      message: `Updated to v${info.latestVersion}. Restart the CLI to use the new version.`,
    });
    return;
  }

  const detail = (result.stderr || result.stdout || '').trim().split('\n').slice(-3).join('\n');
  ctx.addBlock({
    type: 'error',
    message:
      `Update failed (${result.command}).` +
      (detail ? `\n${detail}` : '') +
      `\nYou can run \`${result.command}\` manually.`,
  });
}
