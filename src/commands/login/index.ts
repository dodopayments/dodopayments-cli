import DodoPayments from 'dodopayments';
import { input, select } from '@inquirer/prompts';
import open from 'open';
import { writeConfig } from '../../utils/auth';

export async function handleLogin(): Promise<void> {
    await open('https://app.dodopayments.com/developer/api-keys');

    const apiKey = await input({ message: 'Enter your Dodo Payments API Key:', required: true });

    const mode = await select({
        choices: [{ name: 'Test Mode', value: 'test_mode' }, { name: 'Live Mode', value: 'live_mode' }],
        message: 'Choose the environment:'
    }) as 'test_mode' | 'live_mode';

    const client = new DodoPayments({ bearerToken: apiKey, environment: mode });

    console.log('Verifying Dodo Payments API Key');
    try {
        await client.products.list({ page_size: 1 });
        console.log('Successfully verified your Dodo Payments API Key!');
    } catch {
        console.log('Something went wrong while authenticating, please check your API key!');
        process.exit(1);
    }

    console.log('Storing / Updating existing configuration...');
    writeConfig(mode, apiKey);
    console.log('Setup complete successfully!');
    process.exit(0);
}

