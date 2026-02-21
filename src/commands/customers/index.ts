import DodoPayments from 'dodopayments';
import { input } from '@inquirer/prompts';
import { usage } from '../../utils/usage-help';
import { isDodoPaymentsAPIError } from '../../utils/error';

export async function handleCustomers( client: DodoPayments,subCommand?: string) {
    if (subCommand === 'list') {
        const page = await input({ message: 'Enter page:', default: '1', validate: e => e.trim() !== '' });
        console.table((await client.customers.list({ page_number: parseInt(page) - 1, page_size: 100 })).items, ['customer_id', 'name', 'email', 'phone_number']);

    } else if (subCommand === 'create') {
        const name = await input({ message: 'Enter Name: ', validate: e => e.trim() !== '' });
        const email = await input({ message: 'Enter Email: ', validate: e => e.trim() !== '' });
        const phone = await input({ message: 'Enter Phone Number: ' });
        const creation = await client.customers.create({ name, email, phone_number: phone.trim() !== '' ? phone : null });
        console.log('Customer Successfully Created!');
        console.table([creation], ['customer_id', 'name', 'email', 'phone_number']);

    } else if (subCommand === 'update') {
        try {
            const customer_id = await input({ message: 'Enter customer ID:', validate: e => e.startsWith('cus_') || 'Please enter a valid customer ID!' });
            const existingInfo = await client.customers.retrieve(customer_id);
            const name = await input({ message: 'Enter customer name:', default: existingInfo.name });
            const phone = await input({ message: 'Enter customer phone:', default: existingInfo.phone_number?.toString() });
            const updated = await client.customers.update(customer_id, { name, phone_number: phone.trim() !== '' ? phone : null });
            console.table([updated], ['customer_id', 'name', 'email', 'phone_number']);
        } catch (e) {
            if (isDodoPaymentsAPIError(e) && e.error.code === 'NOT_FOUND') {
                console.log('Incorrect customer ID!');
            } else {
                console.error(e);
            }
        }
    } else {
        usage.customers!.forEach(e => console.log(`dodo customers ${e.command} - ${e.description}`));
    }
}