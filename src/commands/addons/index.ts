import DodoPayments from 'dodopayments';
import { input } from '@inquirer/prompts';
import open from 'open';
import { CurrencyToSymbolMap } from '../../utils/currency-to-symbol-map';
import { usage } from '../../utils/usage-help';
import { isDodoPaymentsAPIError } from '../../utils/error';

export async function handleAddons( client: DodoPayments,subCommand?: string) {
    if (subCommand === 'create') {
        open('https://app.dodopayments.com/products/create/add-on');

    } else if (subCommand === 'list') {
        const page = await input({ message: 'Enter page:', default: '1', validate: e => e.trim() !== '' });
        const addons = await client.addons.list({ page_number: parseInt(page) - 1, page_size: 100 });
        console.table(addons.items.map(e => ({
            id: e.id,
            name: e.name,
            price: `${CurrencyToSymbolMap[e.currency] || e.currency + ' '}${e.price * 0.01}`,
            'created on': new Date(e.created_at).toLocaleString()
        })));

    } else if (subCommand === 'info') {
        try {
            const addon_id = await input({ message: 'Enter addon ID:', validate: e => e.startsWith('adn_') || 'Please enter a valid addon ID!' });
            const info = await client.addons.retrieve(addon_id);
            console.table({
                'addon id': info.id,
                name: info.name,
                price: `${CurrencyToSymbolMap[info.currency] || info.currency + ' '}${info.price * 0.01}`,
                ...info.description?.trim() !== '' && { description: info.description },
                created_at: new Date(info.created_at).toLocaleString(),
                updated_at: new Date(info.updated_at).toLocaleString(),
                tax_category: info.tax_category,
            });
            console.log(`To edit the addon, go to https://app.dodopayments.com/products/edit/add-on?id=${info.id}`);
        } catch (e) {
            if (isDodoPaymentsAPIError(e) && e.error.code === 'NOT_FOUND') {
                console.log('Incorrect addon ID!');
            } else {
                console.error(e);
            }
        }
    } else {
        usage.addons!.forEach(e => console.log(`dodo addons ${e.command} - ${e.description}`));
    }
}