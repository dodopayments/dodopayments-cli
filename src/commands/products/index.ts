import DodoPayments from 'dodopayments';
import open from 'open';
import type { Price } from 'dodopayments/resources';
import { currencyToSymbolMap } from '../../utils/currency-to-symbol-map';
import { isDodoPaymentsAPIError } from '../../utils/error';
import { paginationTip } from '../../utils/tips';
import type { CommandContext } from '../../tui/CommandContext';

export async function handleProducts(
  client: DodoPayments,
  subCommand: string | undefined,
  ctx: CommandContext,
  args: string[] = [],
) {
  if (subCommand === 'list') {
    const pageNum = parseInt(args[0] ?? '1') || 1;
    const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Loading products…' });
    try {
      const fetchedData = await client.products.list({
        page_number: pageNum - 1,
        page_size: 100,
      });
      ctx.removeBlock(spinnerId);

      if (fetchedData.items.length === 0) {
        ctx.addBlock({ type: 'empty' });
        return;
      }

      const table = fetchedData.items.map((e) => ({
        name: e.name,
        'product id': e.product_id,
        created_at: new Date(e.created_at).toLocaleString(),
        ...(e.is_recurring
          ? {
              price: `${currencyToSymbolMap[e.price_detail!.currency] || e.price_detail!.currency + ' '}${(e.price! * 0.01).toFixed(2)} Every ${(e.price_detail as Price.RecurringPrice).payment_frequency_count} ${(e.price_detail as Price.RecurringPrice)?.payment_frequency_interval}`,
            }
          : {
              price: `${currencyToSymbolMap[e.price_detail!.currency] || e.price_detail!.currency + ' '}${(e.price! * 0.01).toFixed(2)} (One Time)`,
            }),
      }));

      ctx.addBlock({ type: 'table', data: table });
      ctx.addBlock({ type: 'link', text: 'To edit a product, go to', url: 'https://app.dodopayments.com/products/edit?id={product_id}' });
      ctx.addBlock({ type: 'streaming', text: paginationTip('/products list') });
    } catch (e: any) {
      ctx.removeBlock(spinnerId);
      if (isDodoPaymentsAPIError(e)) {
        ctx.addBlock({ type: 'error', message: `Couldn't load products. ${e.error.message}` });
      } else {
        ctx.addBlock({ type: 'error', message: e.message });
      }
    }
  } else if (subCommand === 'create') {
    try {
      await open('https://app.dodopayments.com/products/create');
      ctx.addBlock({ type: 'success', message: 'Browser opened to create a product.' });
    } catch (e: any) {
      ctx.addBlock({ type: 'error', message: `Couldn't open browser. ${e.message}` });
    }
  } else if (subCommand === 'info') {
    const product_id = args[0];
    if (!product_id) {
      ctx.addBlock({ type: 'error', message: 'Product ID required. Usage: /products info <id>' });
      return;
    }

    const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Loading product…' });
    try {
      const info = await client.products.retrieve(product_id);
      ctx.removeBlock(spinnerId);

      ctx.addBlock({
        type: 'detail',
        data: {
          'product id': info.product_id,
          name: info.name,
          ...(info.description?.trim() !== '' && {
            description: info.description,
          }),
          created_at: new Date(info.created_at).toLocaleString(),
          updated_at: new Date(info.updated_at).toLocaleString(),
          ...(info.is_recurring
            ? {
                // .fixed_price for usage based billing
                price: `${currencyToSymbolMap[info.price.currency] || info.price.currency + ' '}${(((info.price as Price.RecurringPrice).price || (info.price as Price.UsageBasedPrice).fixed_price) * 0.01).toFixed(2)} Every ${(info.price as Price.RecurringPrice).payment_frequency_count} ${(info.price as Price.RecurringPrice).payment_frequency_interval}`,
              }
            : {
                price: `${currencyToSymbolMap[info.price.currency] || info.price.currency + ' '}${((info.price as Price.OneTimePrice).price * 0.01).toFixed(2)} (One Time)`,
              }),
          tax_category: info.tax_category,
        }
      });

      ctx.addBlock({ type: 'link', text: 'To edit the product, go to', url: `https://app.dodopayments.com/products/edit?id=${info.product_id}` });
    } catch (e: any) {
      ctx.removeBlock(spinnerId);
      if (isDodoPaymentsAPIError(e) && e.error.code === 'NOT_FOUND') {
        ctx.addBlock({ type: 'error', message: 'Product not found. Check the ID and try again.' });
      } else if (isDodoPaymentsAPIError(e)) {
        ctx.addBlock({ type: 'error', message: `Couldn't load this product. ${e.error.message}` });
      } else {
        ctx.addBlock({ type: 'error', message: e.message });
      }
    }
  } else {
    ctx.addBlock({ type: 'help' });
  }
}
