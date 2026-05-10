import DodoPayments from 'dodopayments';
import open from 'open';
import { saveConfig } from '../../utils/auth';
import type { CommandContext } from '../../tui/CommandContext';

export async function handleLogin(ctx: CommandContext): Promise<void> {
  await open('https://app.dodopayments.com/developer/api-keys');

  const apiKey = await ctx.promptInput('API key');

  const mode = (await ctx.promptSelect('Environment', [
    { label: 'Test Mode', value: 'test_mode' },
    { label: 'Live Mode', value: 'live_mode' },
  ])) as 'test_mode' | 'live_mode';

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
    return;
  }

  await saveConfig(mode, apiKey);
  ctx.addBlock({ type: 'success', message: 'Setup complete.' });
}
