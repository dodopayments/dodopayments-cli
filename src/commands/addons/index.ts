import DodoPayments from 'dodopayments';
import open from 'open';
import { currencyToSymbolMap } from '../../utils/currency-to-symbol-map';
import { isDodoPaymentsAPIError } from '../../utils/error';
import { paginationTip } from '../../utils/tips';
import type { CommandContext } from '../../ui/ink/CommandContext';

export async function handleAddons(
  client: DodoPayments,
  subCommand: string | undefined,
  ctx: CommandContext,
  args: string[] = [],
) {
  if (subCommand === 'create') {
    try {
      await open('https://app.dodopayments.com/products/create/add-on');
      ctx.addBlock({ type: 'success', message: 'Browser opened to create an addon.' });
    } catch (e: any) {
      ctx.addBlock({ type: 'error', message: `Couldn't open browser. ${e.message}` });
    }
  } else if (subCommand === 'list') {
    const pageNum = parseInt(args[0] ?? '1') || 1;

    const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Loading addons…' });
    try {
      const addons = await client.addons.list({
        page_number: pageNum - 1,
        page_size: 100,
      });
      ctx.removeBlock(spinnerId);

      if (addons.items.length === 0) {
        ctx.addBlock({ type: 'empty' });
        return;
      }

      const addonsList = addons.items.map((e) => ({
        id: e.id,
        name: e.name,
        price: `${currencyToSymbolMap[e.currency] || e.currency + ' '}${(e.price * 0.01).toFixed(2)}`,
        'created on': new Date(e.created_at).toLocaleString(),
      }));

      ctx.addBlock({ type: 'table', data: addonsList });
      ctx.addBlock({ type: 'streaming', text: paginationTip('/addons list') });
    } catch (e: any) {
      ctx.removeBlock(spinnerId);
      if (isDodoPaymentsAPIError(e)) {
        ctx.addBlock({ type: 'error', message: `Couldn't load addons. ${e.error.message}` });
      } else {
        ctx.addBlock({ type: 'error', message: e.message });
      }
    }
  } else if (subCommand === 'info') {
    const addon_id = args[0];
    if (!addon_id) {
      ctx.addBlock({ type: 'error', message: 'Addon ID required. Usage: /addons info <id>' });
      return;
    }

    const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Loading addon…' });
    try {
      const info = await client.addons.retrieve(addon_id);
      ctx.removeBlock(spinnerId);

      ctx.addBlock({
        type: 'detail',
        data: {
          'addon id': info.id,
          name: info.name,
          price: `${currencyToSymbolMap[info.currency] || info.currency + ' '}${info.price * 0.01}`,
          ...(info.description?.trim() !== '' && {
            description: info.description,
          }),
          created_at: new Date(info.created_at).toLocaleString(),
          updated_at: new Date(info.updated_at).toLocaleString(),
          tax_category: info.tax_category,
        }
      });

      ctx.addBlock({ type: 'link', text: 'To edit the addon, go to', url: `https://app.dodopayments.com/products/edit/add-on?id=${info.id}` });
    } catch (e: any) {
      ctx.removeBlock(spinnerId);
      if (isDodoPaymentsAPIError(e) && e.error.code === 'NOT_FOUND') {
        ctx.addBlock({ type: 'error', message: 'Addon not found. Check the ID and try again.' });
      } else if (isDodoPaymentsAPIError(e)) {
        ctx.addBlock({ type: 'error', message: `Couldn't load this addon. ${e.error.message}` });
      } else {
        ctx.addBlock({ type: 'error', message: e.message });
      }
    }
  } else {
    ctx.addBlock({ type: 'help' });
  }
}
