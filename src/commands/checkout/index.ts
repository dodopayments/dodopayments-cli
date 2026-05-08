import DodoPayments from 'dodopayments';
import { usage } from '../../utils/usage-help';
import { isDodoPaymentsAPIError } from '../../utils/error';
import type { CommandContext } from '../../ui/ink/CommandContext';

export async function handleCheckout(
  client: DodoPayments,
  subCommand: string | undefined,
  ctx: CommandContext,
  args: string[] = [],
) {
  if (subCommand === 'new') {
    let config: DodoPayments.CheckoutSessions.CheckoutSessionCreateParams = {
      product_cart: [],
    };

    const product = await ctx.promptInput('Enter product:');
    config.product_cart = [{ product_id: product, quantity: 1 }];

    const useAdvancedStr = await ctx.promptSelect('Use advanced options?', [
      { label: 'Yes', value: 'yes' },
      { label: 'No', value: 'no' },
    ]);
    const useAdvanced = useAdvancedStr === 'yes';

    if (useAdvanced) {
      const minAddress = await ctx.promptSelect('Enable minimal address:', [
        { label: 'Yes', value: 'yes' },
        { label: 'No', value: 'no' },
      ]);
      config.minimal_address = minAddress === 'yes';

      const return_url = await ctx.promptInput('Enter return URL (Optional):');
      if (return_url.trim() !== '') {
        config.return_url = return_url;
      }

      const force3ds = await ctx.promptSelect('Force 3DS?', [
        { label: 'Yes', value: 'yes' },
        { label: 'No', value: 'no' },
      ]);
      config.force_3ds = force3ds === 'yes';

      const disc_code = await ctx.promptInput('Enter discount code (Optional):');
      if (disc_code.trim() !== '') {
        config.discount_code = disc_code;
      }

      const metadata = await ctx.promptInput('Enter metadata (Optional, JSON stringified):');
      if (metadata.trim() !== '') {
        try {
          config.metadata = JSON.parse(metadata);
        } catch {
          ctx.addBlock({ type: 'error', message: 'Invalid JSON for metadata.' });
          return;
        }
      }
    }
    
    const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Creating checkout session...' });
    try {
      const session = await client.checkoutSessions.create(config);
      ctx.removeBlock(spinnerId);
      ctx.addBlock({ type: 'link', text: 'Checkout Session URL:', url: session.checkout_url ?? '' });
    } catch (e: any) {
      ctx.removeBlock(spinnerId);
      if (isDodoPaymentsAPIError(e)) {
        ctx.addBlock({ type: 'error', message: `Error: ${e.error.message}` });
      } else {
        ctx.addBlock({ type: 'error', message: e.message });
      }
    }
  } else {
    ctx.addBlock({ type: 'help' });
  }
}
