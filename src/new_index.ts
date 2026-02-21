#!/usr/bin/env node
import DodoPayments from 'dodopayments';
import { renderHelp } from './ui';
import { usage } from './utils/usage-help';
import { configExists, resolveCredentials } from './utils/auth';

import { handleLogin } from './commands/login';
import { handleProducts } from './commands/products';
import { handlePayments } from './commands/payments';
import { handleCustomers } from './commands/customers';
import { handleDiscounts } from './commands/discounts';
import { handleLicences } from './commands/licences';
import { handleAddons } from './commands/addons';
import { handleRefund } from './commands/refund';
import { handleCheckout } from './commands/checkout';
import WebhookListener from './commands/webhook/listen';

const args = process.argv;
const category = args[2];
const subCommand = args[3];

if (category === 'login') {
    await handleLogin();
    process.exit(0);
}

if (!configExists()) {
    if (category && category in usage) {
        console.log(`Category: ${category}`);
        
        usage[category as keyof typeof usage]!.forEach(e => {
            console.log(`dodo ${category} ${e.command} - ${e.description}`);
        });
    } else {
        renderHelp();
    }
    console.log('\nPlease login using `dodo login` command first!');
    process.exit(0);
}

// Resolve credentials — handles single/multi-mode and config errors
const { apiKey, mode } = await resolveCredentials();

const dodoClient = new DodoPayments({
    bearerToken: apiKey,
    environment: mode,
});

// Route to command modules
switch (category) {
    case 'products':
        await handleProducts(dodoClient, subCommand);
        break;

    case 'payments':
        await handlePayments(dodoClient, subCommand);
        break;

    case 'customers':
        await handleCustomers(dodoClient, subCommand);
        break;

    case 'discounts':
        await handleDiscounts(dodoClient, subCommand);
        break;

    case 'licences':
        await handleLicences(dodoClient, subCommand);
        break;

    case 'addons':
        await handleAddons(dodoClient, subCommand);
        break;

    case 'refund':
        await handleRefund(dodoClient, subCommand);
        break;

    case 'checkout':
        await handleCheckout(dodoClient, subCommand);
        break;

    case 'wh':
        switch (subCommand) {
            case 'listen':
                WebhookListener({ API_KEY: apiKey, dodoClient });
                break;
            case 'trigger':
                import('./commands/webhook/trigger');
                break;
            default:
                usage.wh!.forEach(e => console.log(`dodo wh ${e.command} - ${e.description}`));
        }
        break;

    default:
        renderHelp();
}