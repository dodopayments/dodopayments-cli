import DodoPayments from 'dodopayments';
import { input } from '@inquirer/prompts';
import { currencyToSymbolMap } from '../../utils/currency-to-symbol-map';
import { usage } from '../../utils/usage-help';
import { isDodoPaymentsAPIError } from '../../utils/error';

export async function handlePayments(client: DodoPayments, subCommand?: string) {
    switch (subCommand) {
        case 'list': {
            const page = await input({ message: 'Enter page:', default: '1', validate: e => e.trim() !== '' });
            try {
                const payments = (await client.payments.list({ page_number: parseInt(page) - 1, page_size: 100 })).items;
                console.table(payments.map(payment => ({
                    'payment id': payment.payment_id,
                    'created at': new Date(payment.created_at).toLocaleString(),
                    'subscription id': payment.subscription_id,
                    'total amount': `${currencyToSymbolMap[payment.currency] || (payment.currency + ' ')}${(payment.total_amount * 0.01).toFixed(2)}`,
                    status: payment.status
                })));
                console.log('To view a payment, go to https://app.dodopayments.com/transactions/payments/{payment_id}');
            } catch (e) {
                console.error('Failed to fetch payments:', e);
            }
            break;
        }

        case 'info': {
            try {
                const payment_id = await input({
                    message: 'Enter payment ID:',
                    validate: e => e.startsWith('pay_') || 'Please enter a valid payment ID!'
                });
                const info = await client.payments.retrieve(payment_id);
                console.table({
                    'payment id': info.payment_id,
                    status: info.status,
                    'total amount': `${currencyToSymbolMap[info.currency] || info.currency + ' '}${(info.total_amount * 0.01).toFixed(2)}`,
                    'payment method': info.payment_method,
                    createdAt: new Date(info.created_at).toLocaleString(),
                    customer: info.customer.customer_id,
                    'customer email': info.customer.email,
                    ...info.subscription_id && { 'subscription id': info.subscription_id },
                    'billing address street': info.billing.street,
                    'billing address state': info.billing.state,
                    'billing address city': info.billing.city,
                    'billing address country': info.billing.country,
                    'billing address zipcode': info.billing.zipcode,
                });
                console.log(`To view the payment, go to https://app.dodopayments.com/transactions/payments/${info.payment_id}`);
            } catch (e) {
                if (isDodoPaymentsAPIError(e) && e.error.code === 'NOT_FOUND') {
                    console.log('Incorrect payment ID!');
                } else {
                    console.error(e);
                }
            }
            break;
        }

        default:
            usage.payments!.forEach(e => console.log(`dodo payments ${e.command} - ${e.description}`));
    }
}