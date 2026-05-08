import DodoPayments from 'dodopayments';
import chalk from 'chalk';
import { currencyToSymbolMap } from '../../utils/currency-to-symbol-map';
import { usage } from '../../utils/usage-help';
import { isDodoPaymentsAPIError } from '../../utils/error';
import type { CommandContext } from '../../ui/ink/CommandContext';

export async function handlePayments(
  client: DodoPayments,
  subCommand: string | undefined,
  ctx: CommandContext,
  args: string[] = [],
) {
  if (subCommand === 'list') {
    const pageNum = parseInt(args[0] ?? '1') || 1;
    const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Fetching payments...' });
    try {
      const payments = (
        await client.payments.list({
          page_number: pageNum - 1,
          page_size: 100,
        })
      ).items;

      ctx.removeBlock(spinnerId);
      
      if (payments.length === 0) {
        ctx.addBlock({ type: 'empty' });
        return;
      }

      ctx.addBlock({
        type: 'table',
        data: payments.map((payment) => ({
          'payment id': payment.payment_id,
          'created at': new Date(payment.created_at).toLocaleString(),
          'subscription id': payment.subscription_id,
          'total amount': `${currencyToSymbolMap[payment.currency] || payment.currency + ' '}${(payment.total_amount * 0.01).toFixed(2)}`,
          status: payment.status,
        })),
        statusColumn: 'status',
      });

      ctx.addBlock({ type: 'link', text: 'To view a payment, go to', url: 'https://app.dodopayments.com/transactions/payments/{payment_id}' });
      ctx.addBlock({ type: 'streaming', text: chalk.dim('\nTip: Use ') + chalk.cyan('/payments list <page>') + chalk.dim(' to see more pages.') });
    } catch (e: any) {
      ctx.removeBlock(spinnerId);
      if (isDodoPaymentsAPIError(e)) {
        ctx.addBlock({ type: 'error', message: `Failed to fetch payments: ${e.error.message}` });
      } else {
        ctx.addBlock({ type: 'error', message: e.message });
      }
    }
  } else if (subCommand === 'info') {
    let payment_id = args[0];
    if (!payment_id) {
      ctx.addBlock({ type: 'error', message: 'Please provide a payment ID! (Usage: /payments info <id>)' });
      return;
    }
    
    const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Fetching payment...' });
    try {
      const info = await client.payments.retrieve(payment_id);
      ctx.removeBlock(spinnerId);

      ctx.addBlock({
        type: 'detail',
        data: {
          'payment id': info.payment_id,
          status: info.status,
          'total amount': `${currencyToSymbolMap[info.currency] || info.currency + ' '}${(info.total_amount * 0.01).toFixed(2)}`,
          'payment method': info.payment_method,
          createdAt: new Date(info.created_at).toLocaleString(),
          customer: info.customer.customer_id,
          'customer email': info.customer.email,
          ...(info.subscription_id && {
            'subscription id': info.subscription_id,
          }),
          'billing address street': info.billing.street,
          'billing address state': info.billing.state,
          'billing address city': info.billing.city,
          'billing address country': info.billing.country,
          'billing address zipcode': info.billing.zipcode,
        }
      });

      ctx.addBlock({ type: 'link', text: 'To view the payment, go to', url: `https://app.dodopayments.com/transactions/payments/${info.payment_id}` });
    } catch (e: any) {
      ctx.removeBlock(spinnerId);
      if (isDodoPaymentsAPIError(e) && e.error.code === 'NOT_FOUND') {
        ctx.addBlock({ type: 'error', message: 'Incorrect payment ID!' });
      } else {
        ctx.addBlock({ type: 'error', message: e.message });
      }
    }
  } else {
    ctx.addBlock({ type: 'help' });
  }
}
