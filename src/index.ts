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
import { handleWebhook } from './commands/webhook';

process.on('SIGINT', () => process.exit(0));
process.on('SIGTERM', () => process.exit(0));

const category = process.argv[2];
const subCommand = process.argv[3];

try {
  if (category === 'login') {
    await handleLogin();
    process.exit(0);
  }

  if (!configExists()) {
    renderHelp(category);
    console.log('\nPlease login using `dodo login` command first!');
    process.exit(0);
  }

  const { apiKey, mode } = await resolveCredentials();

  const dodoClient = new DodoPayments({
    bearerToken: apiKey,
    environment: mode,
  });

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
      await handleWebhook(dodoClient, apiKey, subCommand);
      break;
    default:
      renderHelp();
  }
} catch (e: any) {
  if (e?.name === 'ExitPromptError') process.exit(0);
  console.error('An unexpected error occurred:', e);
  process.exit(1);
}
