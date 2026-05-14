import DodoPayments from 'dodopayments';
import { isDodoPaymentsAPIError } from '../../utils/error';
import { paginationTip } from '../../utils/tips';
import type { CommandContext } from '../../tui/CommandContext';
import { unknownSubcommand } from '../../utils/usage-help';

export async function handleDiscounts(
  client: DodoPayments,
  subCommand: string | undefined,
  ctx: CommandContext,
  args: string[] = [],
) {
  if (subCommand === 'list') {
    const pageNum = parseInt(args[0] ?? '1') || 1;
    const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Loading discounts…' });
    try {
      const discounts = await client.discounts.list({
        page_number: pageNum - 1,
        page_size: 100,
      });
      ctx.removeBlock(spinnerId);

      if (discounts.items.length === 0) {
        ctx.addBlock({ type: 'empty' });
        return;
      }

      const discountsTable = discounts.items.map((e) => ({
        name: e.name,
        code: e.code,
        'discount id': e.discount_id,
        'created at': new Date(e.created_at).toLocaleString(),
        ...(e.type === 'percentage'
          ? {
              amount: `${(e.amount * 0.01).toFixed(2)}%`,
            }
          : {
              amount: e.amount,
            }),
      }));

      ctx.addBlock({ type: 'table', data: discountsTable });
      ctx.addBlock({ type: 'link', text: 'To view a discount, go to', url: 'https://app.dodopayments.com/sales/discounts/edit?id={discount_id}' });
      ctx.addBlock({ type: 'streaming', text: paginationTip('/discounts list') });
    } catch (e: any) {
      ctx.removeBlock(spinnerId);
      if (isDodoPaymentsAPIError(e)) {
        ctx.addBlock({ type: 'error', message: `Couldn't load discounts. ${e.error.message}` });
      } else {
        ctx.addBlock({ type: 'error', message: e.message });
      }
    }
  } else if (subCommand === 'create') {
    const name = await ctx.promptInput('Discount name');
    const percentage = await ctx.promptInput('Discount percentage');
    const code = await ctx.promptInput('Discount code (optional)');
    const cycles = await ctx.promptInput('Discount cycles (optional)');

    const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Creating discount…' });
    try {
      const newDiscount = await client.discounts.create({
        name,
        code: code.trim() !== '' ? code : null,
        amount: parseFloat(percentage) * 100,
        type: 'percentage',
        ...(cycles.trim() !== '' && { subscription_cycles: parseInt(cycles) }),
      });

      ctx.removeBlock(spinnerId);
      ctx.addBlock({ type: 'success', message: 'Discount created.' });
      ctx.addBlock({
        type: 'detail',
        data: {
          name: newDiscount.name,
          code: newDiscount.code,
          'discount id': newDiscount.discount_id,
          ...(cycles.trim() !== '' && {
            'subscription cycles': newDiscount.subscription_cycles,
          }),
        }
      });
    } catch (e: any) {
      ctx.removeBlock(spinnerId);
      if (isDodoPaymentsAPIError(e)) {
        ctx.addBlock({ type: 'error', message: `Couldn't create discount. ${e.error.message}` });
      } else {
        ctx.addBlock({ type: 'error', message: e.message });
      }
    }
  } else if (subCommand === 'delete') {
    const discount_id = args[0];
    if (!discount_id) {
      ctx.addBlock({ type: 'error', message: 'Discount ID required. Usage: /discounts delete <id>' });
      return;
    }

    const confirmed = await ctx.promptConfirm(`Delete discount ${discount_id}?`);
    if (!confirmed) {
      ctx.addBlock({ type: 'error', message: 'Deletion cancelled.' });
      return;
    }

    const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Deleting discount…' });
    try {
      await client.discounts.delete(discount_id);
      ctx.removeBlock(spinnerId);
      ctx.addBlock({ type: 'success', message: 'Discount deleted.' });
    } catch (e: any) {
      ctx.removeBlock(spinnerId);
      if (isDodoPaymentsAPIError(e) && e.error.code === 'NOT_FOUND') {
        ctx.addBlock({ type: 'error', message: 'Discount not found. Check the ID and try again.' });
      } else if (isDodoPaymentsAPIError(e)) {
        ctx.addBlock({ type: 'error', message: `Couldn't delete discount. ${e.error.message}` });
      } else {
        ctx.addBlock({ type: 'error', message: e.message });
      }
    }
  } else {
    unknownSubcommand(ctx, 'discounts', subCommand);
  }
}
