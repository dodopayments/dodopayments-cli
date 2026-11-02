import { readGrants, storePath } from '../../utils/oauth';
import { apiKeyModes, configPath, getSessionMode, unreadableStoreMessages, type Mode } from '../../utils/auth';
import type { CommandContext } from '../../tui/CommandContext';

const ALL_MODES: Mode[] = ['test_mode', 'live_mode'];
const modeLabel = (mode: Mode) => (mode === 'live_mode' ? 'Live Mode' : 'Test Mode');

/** Shows each signed-in mode: the sign-in method, the device name, the scopes, and the store file. */
export async function handleWhoami(ctx: CommandContext): Promise<void> {
  // The same order as `resolveCredentials`: the environment variable, a pasted key, then OAuth.
  if (process.env.DODO_PAYMENTS_API_KEY) {
    ctx.addBlock({ type: 'info', message: 'Signed in with the DODO_PAYMENTS_API_KEY environment variable. It overrides the stored sign-in.' });
    return;
  }
  const grants = readGrants();
  const pasted = await apiKeyModes();
  const modes = ALL_MODES.filter((mode) => pasted.includes(mode) || grants.some((grant) => grant.modes.includes(mode)));
  const unreadable = await unreadableStoreMessages();
  for (const message of unreadable) ctx.addBlock({ type: 'error', message });
  if (unreadable.length > 0) process.exitCode = 1;
  if (modes.length === 0) {
    if (unreadable.length === 0) ctx.addBlock({ type: 'error', message: 'Not signed in. Run `dodo login`.' });
    process.exitCode = 1;
    return;
  }

  // `resolveCredentials` uses the saved mode when it is signed in, and else the first mode.
  const saved = getSessionMode();
  const active = saved && modes.includes(saved) ? saved : modes[0];

  for (const mode of modes) {
    const grant = grants.find((candidate) => candidate.modes.includes(mode));
    const data: Record<string, string> = { Mode: mode === active ? `${modeLabel(mode)} (active)` : modeLabel(mode) };
    if (pasted.includes(mode)) {
      data['Sign-in'] = 'API key';
      if (grant) data.Note = 'The API key has priority over the browser sign-in of this mode.';
      data.Storage = `${configPath()} (encrypted)`;
    } else if (grant) {
      data['Sign-in'] = 'Browser or device code (OAuth)';
      data.Device = grant.device_name;
      data.Scopes = grant.scope || 'Unknown';
      data['Signed in at'] = grant.signed_in_at;
      data.Storage = `${storePath()} (encrypted)`;
    }
    ctx.addBlock({ type: 'detail', data });
  }
  if (modes.length > 1) {
    ctx.addBlock({ type: 'info', message: 'To change the active mode, run /switch in the TUI.' });
  }
}
