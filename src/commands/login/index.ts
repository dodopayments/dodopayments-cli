import DodoPayments from 'dodopayments';
import open from 'open';
import { saveConfig } from '../../utils/auth';
import type { CommandContext } from '../../ui/ink/CommandContext';

export async function handleLogin(ctx: CommandContext): Promise<void> {
  await open('https://app.dodopayments.com/developer/api-keys');

  const apiKey = await ctx.promptInput('Enter your Dodo Payments API Key:');

  const mode = (await ctx.promptSelect('Choose the environment:', [
    { label: 'Test Mode', value: 'test_mode' },
    { label: 'Live Mode', value: 'live_mode' },
  ])) as 'test_mode' | 'live_mode';

  const client = new DodoPayments({ bearerToken: apiKey, environment: mode });

  const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Verifying Dodo Payments API Key...' });
  try {
    await client.products.list({ page_size: 1 });
    ctx.removeBlock(spinnerId);
    ctx.addBlock({ type: 'success', message: 'Successfully verified your Dodo Payments API Key!' });
  } catch {
    ctx.removeBlock(spinnerId);
    ctx.addBlock({
      type: 'error',
      message: 'Something went wrong while authenticating. Please check your API key and selected environment.',
    });
    process.exitCode = 1;
    return;
  }

  await saveConfig(mode, apiKey);
  ctx.addBlock({ type: 'success', message: 'Setup complete successfully!' });
}
