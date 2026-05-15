import DodoPayments from 'dodopayments';
import open from 'open';
import { saveConfig, type Mode } from '../../utils/auth';
import type { CommandContext } from '../../tui/CommandContext';

const USAGE = 'Usage: dodo login <api-key> <test|live>';

function normalizeMode(value: string): Mode | null {
  if (value === 'test' || value === 'test_mode') return 'test_mode';
  if (value === 'live' || value === 'live_mode') return 'live_mode';
  return null;
}

export async function handleLogin(
  ctx: CommandContext,
  apiKeyArg?: string,
  modeArg?: string,
): Promise<boolean> {
  let apiKey: string;
  let mode: Mode;

  if (ctx.invocation === 'cli') {
    if (!apiKeyArg || !modeArg) {
      ctx.addBlock({ type: 'error', message: 'API key and mode are required.' });
      ctx.addBlock({ type: 'info', message: USAGE });
      return false;
    }
    const normalized = normalizeMode(modeArg);
    if (!normalized) {
      ctx.addBlock({ type: 'error', message: `Invalid mode '${modeArg}'. Use 'test' or 'live'.` });
      ctx.addBlock({ type: 'info', message: USAGE });
      return false;
    }
    apiKey = apiKeyArg;
    mode = normalized;
  } else {
    await open('https://app.dodopayments.com/developer/api-keys');
    apiKey = await ctx.promptInput('API key');
    mode = (await ctx.promptSelect('Environment', [
      { label: 'Test Mode', value: 'test_mode' },
      { label: 'Live Mode', value: 'live_mode' },
    ])) as Mode;
  }

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
