import DodoPayments from 'dodopayments';
import { input } from '@inquirer/prompts';
import open from 'open';
import type { Price } from 'dodopayments/resources';
import { currencyToSymbolMap } from '../../utils/currency-to-symbol-map';
import { usage } from '../../utils/usage-help';
import { isDodoPaymentsAPIError } from '../../utils/error';

export async function handleProducts( client: DodoPayments,subCommand?: string) {
    if (subCommand === 'list') {
        const page = await input({
            message: 'Enter page:',
            default: "1",
            validate: (e => e.trim() !== '')
        });

        const fetchedData = await client.products.list({ page_number: parseInt(page) - 1, page_size: 100 });
        const table = fetchedData.items.map(e => ({
            name: e.name,
            'product id': e.product_id,
            created_at: new Date(e.created_at).toLocaleString(),
            ...e.is_recurring ? {
                price: `${currencyToSymbolMap[e.price_detail!.currency] || (e.price_detail!.currency + ' ')}${(e.price! * 0.01).toFixed(2)} Every ${(e.price_detail as Price.RecurringPrice).payment_frequency_count} ${(e.price_detail as Price.RecurringPrice)?.payment_frequency_interval}`,
            } : {
                price: `${currencyToSymbolMap[e.price_detail!.currency] || (e.price_detail!.currency + ' ')}${(e.price! * 0.01).toFixed(2)} (One Time)`,
            },
        }));

        console.table(table);
        console.log("To edit a product, go to https://app.dodopayments.com/products/edit?id={product_id}")
    } else if (subCommand === 'create') {
        open('https://app.dodopayments.com/products/create');

    } else if (subCommand === 'info') {
         try {
                    const product_id = await input({
                        message: "Enter product ID:",
                        validate: (e => e.startsWith('pdt_') || 'Please enter a valid product ID!')
                    });
        
                    const info = await DodoClient.products.retrieve(product_id);
                    console.table({
                        'product id': info.product_id,
                        name: info.name,
                        ...info.description?.trim() !== '' && { description: info.description },
                        created_at: new Date(info.created_at).toLocaleString(),
                        updated_at: new Date(info.updated_at).toLocaleString(),
                        ...info.is_recurring ? {
                            // .fixed_price for usage based billing
                            price: `${currencyToSymbolMap[info.price.currency] || (info.price.currency + ' ')}${(((info.price as Price.RecurringPrice).price || (info.price as Price.UsageBasedPrice).fixed_price) * 0.01).toFixed(2)} Every ${(info.price as Price.RecurringPrice).payment_frequency_count} ${(info.price as Price.RecurringPrice).payment_frequency_interval}`,
                        } : {
                            price: `${currencyToSymbolMap[info.price.currency] || (info.price.currency + ' ')}${((info.price as Price.OneTimePrice).price * 0.01).toFixed(2)} (One Time)`,
                        },
                        tax_category: info.tax_category,
                    });
                    console.log(`To edit the product, go to https://app.dodopayments.com/products/edit?id=${info.product_id}`)
                } catch (e) {
                    if (isDodoPaymentsAPIError(e) && e.error.code === "NOT_FOUND") {
                        console.log("Incorrect product ID!");
                    } else {
                        console.error(e);
                    }
                }
    } else {
        usage.products!.forEach(e => console.log(`dodo products ${e.command} - ${e.description}`));
    }
}