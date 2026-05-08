import DodoPayments from 'dodopayments';
import chalk from 'chalk';
import { usage } from '../../utils/usage-help';
import { isDodoPaymentsAPIError } from '../../utils/error';
import type { CommandContext } from '../../ui/ink/CommandContext';

export async function handleLicences(
  client: DodoPayments,
  subCommand: string | undefined,
  ctx: CommandContext,
  args: string[] = [],
) {
  if (subCommand === 'list') {
    const pageNum = parseInt(args[0] ?? '1') || 1;
    
    const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Fetching licences...' });
    try {
      const licences = await client.licenseKeys.list({
        page_number: pageNum - 1,
        page_size: 100,
      });

      ctx.removeBlock(spinnerId);

      if (licences.items.length === 0) {
        ctx.addBlock({ type: 'empty' });
        return;
      }

      ctx.addBlock({ type: 'table', data: licences.items });
      ctx.addBlock({ type: 'streaming', text: chalk.dim('\nTip: Use ') + chalk.cyan('/licences list <page>') + chalk.dim(' to see more pages.') });
    } catch (e: any) {
      ctx.removeBlock(spinnerId);
      if (isDodoPaymentsAPIError(e)) {
        ctx.addBlock({ type: 'error', message: `Failed to fetch licences: ${e.error.message}` });
      } else {
        ctx.addBlock({ type: 'error', message: e.message });
      }
    }
  } else {
    ctx.addBlock({ type: 'help' });
  }
}
