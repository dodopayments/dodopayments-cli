import DodoPayments from 'dodopayments';
import { isDodoPaymentsAPIError } from '../../utils/error';
import type { CommandContext } from '../../tui/CommandContext';

export async function handleCustomerPortal(
  client: DodoPayments,
  customerId: string | undefined,
  ctx: CommandContext,
) {
  if (!customerId) {
    ctx.addBlock({
      type: 'error',
      message: 'Customer ID required. Usage: /customer-portal <customer id>',
    });
    return;
  }

  const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Creating customer portal session…' });
  try {
    const session = await client.customers.customerPortal.create(customerId);
    ctx.removeBlock(spinnerId);

    ctx.addBlock({ type: 'success', message: 'Customer portal session created.' });
    ctx.addBlock({
      type: 'link',
      text: 'URL: ',
      url: session.link,
    });
  } catch (e: any) {
    ctx.removeBlock(spinnerId);
    if (isDodoPaymentsAPIError(e)) {
      ctx.addBlock({
        type: 'error',
        message: `Couldn't create customer portal session. ${e.error.message}`,
      });
    } else {
      ctx.addBlock({ type: 'error', message: e.message });
    }
  }
}
