import DodoPayments from 'dodopayments';
import { currencyToSymbolMap } from '../../utils/currency-to-symbol-map';
import { isDodoPaymentsAPIError } from '../../utils/error';
import { paginationTip } from '../../utils/tips';
import type { CommandContext } from '../../ui/ink/CommandContext';

export async function handleRefund(
  client: DodoPayments,
  subCommand: string | undefined,
  ctx: CommandContext,
  args: string[] = [],
) {
  if (subCommand === 'list') {
    const pageNum = parseInt(args[0] ?? '1') || 1;

    const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Loading refunds…' });
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
      ctx.addBlock({ type: 'streaming', text: paginationTip('/refunds list') });
    } catch (e: any) {
      ctx.removeBlock(spinnerId);
      if (isDodoPaymentsAPIError(e)) {
        ctx.addBlock({ type: 'error', message: `Couldn't load refunds. ${e.error.message}` });
      } else {
        ctx.addBlock({ type: 'error', message: e.message });
      }
    }
  } else if (subCommand === 'info') {
    const refund_id = args[0];
    if (!refund_id) {
      ctx.addBlock({ type: 'error', message: 'Refund ID required. Usage: /refunds info <id>' });
      return;
    }

    const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Loading refund…' });
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
        ctx.addBlock({ type: 'error', message: 'Refund not found. Check the ID and try again.' });
      } else if (isDodoPaymentsAPIError(e)) {
        ctx.addBlock({ type: 'error', message: `Couldn't load this refund. ${e.error.message}` });
      } else {
        ctx.addBlock({ type: 'error', message: e.message });
      }
    }
  } else {
    ctx.addBlock({ type: 'help' });
  }
}
