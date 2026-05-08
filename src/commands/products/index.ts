import DodoPayments from 'dodopayments';
import chalk from 'chalk';
import open from 'open';
import type { Price } from 'dodopayments/resources';
import { currencyToSymbolMap } from '../../utils/currency-to-symbol-map';
import { usage } from '../../utils/usage-help';
import { isDodoPaymentsAPIError } from '../../utils/error';
import type { CommandContext } from '../../ui/ink/CommandContext';

export async function handleProducts(
  client: DodoPayments,
  subCommand: string | undefined,
  ctx: CommandContext,
  args: string[] = [],
) {
  if (subCommand === 'list') {
    const pageNum = parseInt(args[0] ?? '1') || 1;
    const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Fetching products...' });
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
      ctx.addBlock({ type: 'streaming', text: chalk.dim('\nTip: Use ') + chalk.cyan('/products list <page>') + chalk.dim(' to see more pages.') });
    } catch (e: any) {
      ctx.removeBlock(spinnerId);
      if (isDodoPaymentsAPIError(e)) {
        ctx.addBlock({ type: 'error', message: `Failed to fetch products: ${e.error.message}` });
      } else {
        ctx.addBlock({ type: 'error', message: e.message });
      }
    }
  } else if (subCommand === 'create') {
    try {
      await open('https://app.dodopayments.com/products/create');
      ctx.addBlock({ type: 'success', message: 'Opened browser to create product.' });
    } catch (e: any) {
      ctx.addBlock({ type: 'error', message: `Failed to open browser: ${e.message}` });
    }
  } else if (subCommand === 'info') {
    const product_id = args[0];
    if (!product_id) {
      ctx.addBlock({ type: 'error', message: 'Please provide a product ID! (Usage: /products info <id>)' });
      return;
    }

    const spinnerId = ctx.addBlock({ type: 'spinner', label: 'Fetching product...' });
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
        ctx.addBlock({ type: 'error', message: 'Incorrect product ID!' });
      } else {
        ctx.addBlock({ type: 'error', message: e.message });
      }
    }
  } else {
    ctx.addBlock({ type: 'help' });
  }
}
