import { readStore, storePath } from '../../utils/oauth';
import { apiKeyModes, getSessionMode } from '../../utils/auth';
import type { CommandContext } from '../../tui/CommandContext';

const modeLabel = (mode: string) => (mode === 'live_mode' ? 'Live Mode' : 'Test Mode');

/** Shows the stored sign-in: device name, modes, scopes, and the store file. */
export async function handleWhoami(ctx: CommandContext): Promise<void> {
  // The same order as `resolveCredentials`: the environment variable, a pasted key, then OAuth.
  if (process.env.DODO_PAYMENTS_API_KEY) {
    ctx.addBlock({ type: 'info', message: 'Signed in with the DODO_PAYMENTS_API_KEY environment variable. It overrides the stored sign-in.' });
    return;
  }
  const store = readStore();
  const active = getSessionMode();
  const pasted = await apiKeyModes();
  if (pasted.length > 0) {
    ctx.addBlock({
      type: 'info',
      message: `API key for: ${pasted.map(modeLabel).join(', ')}. A pasted key has priority over the browser sign-in for its mode.`,
    });
  }
  if (store) {
    ctx.addBlock({
      type: 'detail',
      data: {
        'Sign-in': 'Browser or device code (OAuth)',
        Device: store.device_name,
        Modes: store.modes.map(modeLabel).join(', '),
        'Active mode': active ? modeLabel(active) : modeLabel(store.modes[0] ?? 'test_mode'),
        Scopes: store.scope || 'Unknown',
        'Signed in at': store.signed_in_at,
        Storage: `${storePath()} (encrypted)`,
      },
    });
    return;
  }
  if (pasted.length > 0) {
    if (active) ctx.addBlock({ type: 'info', message: `Active mode: ${modeLabel(active)}.` });
    return;
  }
  ctx.addBlock({ type: 'error', message: 'Not signed in. Run `dodo login`.' });
  process.exitCode = 1;
}
