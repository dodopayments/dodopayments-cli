import DodoPayments from 'dodopayments';
import { input } from '@inquirer/prompts';
import { usage } from '../../utils/usage-help';

export async function handleDiscounts( client: DodoPayments,subCommand?: string) {
    if (subCommand === 'list') {
        const page = await input({
            message: 'Enter page:',
            default: "1",
            validate: (e => e.trim() !== '')
        });
        const discounts = await client.discounts.list({ page_number: parseInt(page) - 1, page_size: 100 });
        const discountsTable = discounts.items.map(e => (
            {
                name: e.name,
                code: e.code,
                'discount id': e.discount_id,
                'created at': new Date(e.created_at).toLocaleString(),
                ...e.type === 'percentage' ? {
                    amount: `${(e.amount * 0.01).toFixed(2)}%`
                } : {
                    // I just added this in case of a breaking change in the future
                    amount: e.amount
                },
            }
        ));

        console.table(discountsTable);
        console.log(`To view a discount, go to https://app.dodopayments.com/sales/discounts/edit?id={discount_id}`)
    } else if (subCommand === 'create') {
        const name = await input({
            message: "Enter discount name:",
            validate: (e => e.trim() !== '')
        });

        const percentage = await input({
            message: "Enter discount percentage:",
            // Make sure user enters valid value
            validate: (e => {
                const parsed = parseFloat(e);
                if (!Number.isNaN(parsed) && parsed > 0 && parsed <= 100) {
                    return true;
                } else {
                    return false;
                }
            })
        });

        const code = await input({
            message: "Enter discount code (Optional):"
        });

        const cycles = await input({
            message: "Enter discount cycles (Optional):"
        });

        const newDiscount = await client.discounts.create({
            name,
            code: code.trim() !== '' ? code : null,
            amount: parseFloat(percentage) * 100,
            type: 'percentage',
            // If the subscription cycles is provided
            ...cycles.trim() !== '' && {
                subscription_cycles: parseInt(cycles)
            }
        });

        console.log('Discount created successfully!');
        console.table({
            name: newDiscount.name,
            code: newDiscount.code,
            'discount id': newDiscount.discount_id,
            ...cycles.trim() !== '' && {
                'subscription cycles': newDiscount.subscription_cycles
            }
        });
    } else if (subCommand === 'delete') {
        await client.discounts.delete(await input({
            message: "Enter discount ID to be deleted:",
            validate: (e => e.startsWith('dsc_'))
        }));

        console.log("Successfully deleted discount!");
    } else {
        usage.discounts!.forEach(e => {
            console.log(`dodo discounts ${e.command} - ${e.description}`)
        });
    }
}