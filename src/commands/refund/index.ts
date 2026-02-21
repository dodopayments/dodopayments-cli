import DodoPayments from 'dodopayments';
import { input } from '@inquirer/prompts';
import { CurrencyToSymbolMap } from '../../utils/currency-to-symbol-map';
import { usage } from '../../utils/usage-help';
import { isDodoPaymentsAPIError } from '../../utils/error';

export async function handleRefund( client: DodoPayments,subCommand?: string,) {
    if (subCommand === 'list') {
        const page = await input({ message: 'Enter page:', default: '1', validate: e => e.trim() !== '' });
        const refunds = await client.refunds.list({ page_number: parseInt(page) - 1, page_size: 100 });
        console.table(refunds.items.map(e => ({
            id: e.refund_id,
            'payment id': e.payment_id,
            price: `${CurrencyToSymbolMap[e.currency || ''] || e.currency + ' '}${(e.amount || 0) * 0.01}`,
            'created on': new Date(e.created_at).toLocaleString()
        })));

    } else if (subCommand === 'info') {
        try {
            const refund_id = await input({ message: 'Enter refund ID:', validate: e => e.startsWith('ref_') || 'Please enter a valid refund ID!' });
            const info = await client.refunds.retrieve(refund_id);
            console.table({
                'refund id': info.refund_id,
                'payment id': info.payment_id,
                ...Object.keys(info.metadata).length > 0 && { metadata: info.metadata },
                'customer id': info.customer.email,
                'refund type': info.is_partial ? 'Partial' : 'Full',
                price: `${CurrencyToSymbolMap[info.currency || ''] || info.currency + ' '}${(info.amount || 0) * 0.01}`,
                ...info.reason?.trim() !== '' && { reason: info.reason },
                created_at: new Date(info.created_at).toLocaleString(),
            });
        } catch (e) {
            if (isDodoPaymentsAPIError(e) && e.error.code === 'NOT_FOUND') {
                console.log('Incorrect refund ID!');
            } else {
                console.error(e);
            }
        }
    } else {
        usage.refunds!.forEach(e => console.log(`dodo refunds ${e.command} - ${e.description}`));
    }
}