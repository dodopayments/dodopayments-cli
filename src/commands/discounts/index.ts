import DodoPayments from 'dodopayments';
import chalk from 'chalk';
import { usage } from '../../utils/usage-help';
import { isDodoPaymentsAPIError } from '../../utils/error';
import type { CommandContext } from '../../ui/ink/CommandContext';

export async function handleDiscounts(
  client: DodoPayments,
  subCommand: string | undefined,
  ctx: CommandContext,
  args: string[] = [],
) {
  if (subCommand === 'list') {
    const pageNum = parseInt(args[0] ?? '1') || 1;
    const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Fetching discounts...' });
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
      ctx.addBlock({ type: 'streaming', text: chalk.dim('\nTip: Use ') + chalk.cyan('/discounts list <page>') + chalk.dim(' to see more pages.') });
    } catch (e: any) {
      ctx.removeBlock(spinnerId);
      if (isDodoPaymentsAPIError(e)) {
        ctx.addBlock({ type: 'error', message: `Failed to fetch discounts: ${e.error.message}` });
      } else {
        ctx.addBlock({ type: 'error', message: e.message });
      }
    }
  } else if (subCommand === 'create') {
    const name = await ctx.promptInput('Enter discount name:');
    const percentage = await ctx.promptInput('Enter discount percentage:');
    const code = await ctx.promptInput('Enter discount code (Optional):');
    const cycles = await ctx.promptInput('Enter discount cycles (Optional):');
    
    const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Creating discount...' });
    try {
      const newDiscount = await client.discounts.create({
        name,
        code: code.trim() !== '' ? code : null,
        amount: parseFloat(percentage) * 100,
        type: 'percentage',
        ...(cycles.trim() !== '' && { subscription_cycles: parseInt(cycles) }),
      });

      ctx.removeBlock(spinnerId);
      ctx.addBlock({ type: 'success', message: 'Discount created successfully!' });
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
        ctx.addBlock({ type: 'error', message: `Failed to create discount: ${e.error.message}` });
      } else {
        ctx.addBlock({ type: 'error', message: e.message });
      }
    }
  } else if (subCommand === 'delete') {
    const discount_id = args[0];
    if (!discount_id) {
      ctx.addBlock({ type: 'error', message: 'Please provide a discount ID! (Usage: /discounts delete <id>)' });
      return;
    }
    
    const confirmed = await ctx.promptConfirm(`Are you sure you want to delete discount ${discount_id}?`);
    if (!confirmed) {
      ctx.addBlock({ type: 'error', message: 'Deletion cancelled.' });
      return;
    }

    const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Deleting discount...' });
    try {
      await client.discounts.delete(discount_id);
      ctx.removeBlock(spinnerId);
      ctx.addBlock({ type: 'success', message: 'Successfully deleted discount!' });
    } catch (e: any) {
      ctx.removeBlock(spinnerId);
      if (isDodoPaymentsAPIError(e) && e.error.code === 'NOT_FOUND') {
        ctx.addBlock({ type: 'error', message: 'Incorrect discount ID!' });
      } else if (isDodoPaymentsAPIError(e)) {
        ctx.addBlock({ type: 'error', message: `Failed to delete discount: ${e.error.message}` });
      } else {
        ctx.addBlock({ type: 'error', message: e.message });
      }
    }
  } else {
    ctx.addBlock({ type: 'help' });
  }
}
