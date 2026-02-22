import DodoPayments from 'dodopayments';
import { input } from '@inquirer/prompts';
import open from 'open';
import { currencyToSymbolMap } from '../../utils/currency-to-symbol-map';
import { usage } from '../../utils/usage-help';
import { isDodoPaymentsAPIError } from '../../utils/error';

export async function handleAddons(client: DodoPayments, subCommand?: string) {
  if (subCommand === 'create') {
    open('https://app.dodopayments.com/products/create/add-on');
  } else if (subCommand === 'list') {
    const page = await input({
      message: 'Enter page:',
      default: '1',
      validate: (e) => e.trim() !== '',
    });
    try {
      const addons = await client.addons.list({
        page_number: parseInt(page) - 1,
        page_size: 100,
      });

      const addonsList = addons.items.map((e) => ({
        id: e.id,
        name: e.name,
        price: `${currencyToSymbolMap[e.currency] || e.currency + ' '}${(e.price * 0.01).toFixed(2)}`,
        'created on': new Date(e.created_at).toLocaleString(),
      }));

      console.table(addonsList);
    } catch (e) {
      if (isDodoPaymentsAPIError(e)) {
        console.log('Failed to fetch addons:', e.error.message);
      } else {
        console.error(e);
      }
    }
  } else if (subCommand === 'info') {
    const addon_id = await input({
      message: 'Enter addon ID:',
      validate: (e) => e.startsWith('adn_') || 'Please enter a valid addon ID!',
    });
    try {
      const info = await client.addons.retrieve(addon_id);

      console.table({
        'addon id': info.id,
        name: info.name,
        price: `${currencyToSymbolMap[info.currency] || info.currency + ' '}${info.price * 0.01}`,
        ...(info.description?.trim() !== '' && {
          description: info.description,
        }),
        created_at: new Date(info.created_at).toLocaleString(),
        updated_at: new Date(info.updated_at).toLocaleString(),
        tax_category: info.tax_category,
      });

      console.log(
        `To edit the addon, go to https://app.dodopayments.com/products/edit/add-on?id=${info.id}`,
      );
    } catch (e) {
      if (isDodoPaymentsAPIError(e) && e.error.code === 'NOT_FOUND') {
        console.log('Incorrect addon ID!');
      } else if (isDodoPaymentsAPIError(e)) {
        console.log('Failed to fetch addon:', e.error.message);
      } else {
        console.error(e);
      }
    }
  } else {
    usage.addons!.forEach((e) =>
      console.log(`dodo addons ${e.command} - ${e.description}`),
    );
  }
}
