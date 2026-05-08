import DodoPayments from 'dodopayments';
import chalk from 'chalk';
import { currencyToSymbolMap } from '../../utils/currency-to-symbol-map';
import { usage } from '../../utils/usage-help';
import { isDodoPaymentsAPIError } from '../../utils/error';
import type { CommandContext } from '../../ui/ink/CommandContext';

export async function handleRefund(
  client: DodoPayments,
  subCommand: string | undefined,
  ctx: CommandContext,
  args: string[] = [],
) {
  if (subCommand === 'list') {
    const pageNum = parseInt(args[0] ?? '1') || 1;

    const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Fetching refunds...' });
    try {
      const refunds = await client.refunds.list({
        page_number: pageNum - 1,
        page_size: 100,
      });
      ctx.removeBlock(spinnerId);

      if (refunds.items.length === 0) {
        ctx.addBlock({ type: 'empty' });
        return;
      }

      const refundsList = refunds.items.map((e) => ({
        id: e.refund_id,
        'payment id': e.payment_id,
        price: `${currencyToSymbolMap[e.currency || ''] || e.currency + ' '}${((e.amount || 0) * 0.01).toFixed(2)}`,
        'created on': new Date(e.created_at).toLocaleString(),
      }));

      ctx.addBlock({ type: 'table', data: refundsList });
      ctx.addBlock({ type: 'streaming', text: chalk.dim('\nTip: Use ') + chalk.cyan('/refunds list <page>') + chalk.dim(' to see more pages.') });
    } catch (e: any) {
      ctx.removeBlock(spinnerId);
      if (isDodoPaymentsAPIError(e)) {
        ctx.addBlock({ type: 'error', message: `Failed to fetch refunds: ${e.error.message}` });
      } else {
        ctx.addBlock({ type: 'error', message: e.message });
      }
    }
  } else if (subCommand === 'info') {
    const refund_id = args[0];
    if (!refund_id) {
      ctx.addBlock({ type: 'error', message: 'Please provide a refund ID! (Usage: /refunds info <id>)' });
      return;
    }
    
    const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Fetching refund...' });
    try {
      const info = await client.refunds.retrieve(refund_id);
      ctx.removeBlock(spinnerId);

      ctx.addBlock({
        type: 'detail',
        data: {
          'refund id': info.refund_id,
          'payment id': info.payment_id,
          ...(Object.keys(info.metadata).length > 0 && {
            metadata: info.metadata,
          }),
          'customer id': info.customer.email,
          'refund type': info.is_partial ? 'Partial' : 'Full',
          price: `${currencyToSymbolMap[info.currency || ''] || info.currency + ' '}${((info.amount || 0) * 0.01).toFixed(2)}`,
          ...(info.reason?.trim() !== '' && { reason: info.reason }),
          created_at: new Date(info.created_at).toLocaleString(),
        }
      });
    } catch (e: any) {
      ctx.removeBlock(spinnerId);
      if (isDodoPaymentsAPIError(e) && e.error.code === 'NOT_FOUND') {
        ctx.addBlock({ type: 'error', message: 'Incorrect refund ID!' });
      } else if (isDodoPaymentsAPIError(e)) {
        ctx.addBlock({ type: 'error', message: `Failed to fetch refund: ${e.error.message}` });
      } else {
        ctx.addBlock({ type: 'error', message: e.message });
      }
    }
  } else {
    ctx.addBlock({ type: 'help' });
  }
}
