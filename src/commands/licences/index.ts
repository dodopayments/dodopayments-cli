import DodoPayments from 'dodopayments';
import { isDodoPaymentsAPIError } from '../../utils/error';
import { paginationTip } from '../../utils/tips';
import type { CommandContext } from '../../tui/CommandContext';
import { unknownSubcommand } from '../../utils/usage-help';

export async function handleLicences(
  client: DodoPayments,
  subCommand: string | undefined,
  ctx: CommandContext,
  args: string[] = [],
) {
  if (subCommand === 'list') {
    const pageNum = parseInt(args[0] ?? '1') || 1;

    const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Loading licences…' });
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
      ctx.addBlock({ type: 'streaming', text: paginationTip('/licences list') });
    } catch (e: any) {
      ctx.removeBlock(spinnerId);
      if (isDodoPaymentsAPIError(e)) {
        ctx.addBlock({ type: 'error', message: `Couldn't load licences. ${e.error.message}` });
      } else {
        ctx.addBlock({ type: 'error', message: e.message });
      }
    }
  } else {
    unknownSubcommand(ctx, 'licences', subCommand);
  }
}
