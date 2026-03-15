#!/usr/bin/env node
import { renderHelp } from './ui';
import { configExists, resolveCredentials } from './utils/auth';

process.on('SIGINT', () => process.exit(0));
process.on('SIGTERM', () => process.exit(0));

const category = process.argv[2];
const subCommand = process.argv[3];

try {
  if (category === 'login') {
    const { handleLogin } = await import('./commands/login');
    await handleLogin();
  } else if (category === 'logout') {
    const { handleLogout } = await import('./commands/logout');
    await handleLogout();
  } else if (category === 'wh' && subCommand === 'trigger') {
    const { handleWebhook } = await import('./commands/webhook');
    await handleWebhook(subCommand);
  } else if (!configExists()) {
    if (category === 'wh' && !subCommand) {
      renderHelp(category);
      console.log();
      console.log(
        'Run `dodo wh trigger` without logging in, or `dodo login` to use `dodo wh listen`.',
      );
      process.exit(0);
    }

    if (category && !subCommand) {
      renderHelp(category);
      console.log();
    } else if (!category && !subCommand) {
      renderHelp();
      console.log();
    }

    console.log(
      'Login with `dodo login` to use authenticated commands. `dodo logout` and `dodo wh trigger` are available without login.',
    );
    process.exit(0);
  } else {
    const { apiKey, mode } = await resolveCredentials();
    const { default: DodoPayments } = await import('dodopayments');

    const dodoClient = new DodoPayments({
      bearerToken: apiKey,
      environment: mode,
    });

    switch (category) {
      case 'products': {
        const { handleProducts } = await import('./commands/products');
        await handleProducts(dodoClient, subCommand);
        break;
      }
      case 'payments': {
        const { handlePayments } = await import('./commands/payments');
        await handlePayments(dodoClient, subCommand);
        break;
      }
      case 'customers': {
        const { handleCustomers } = await import('./commands/customers');
        await handleCustomers(dodoClient, subCommand);
        break;
      }
      case 'discounts': {
        const { handleDiscounts } = await import('./commands/discounts');
        await handleDiscounts(dodoClient, subCommand);
        break;
      }
      case 'licences': {
        const { handleLicences } = await import('./commands/licences');
        await handleLicences(dodoClient, subCommand);
        break;
      }
      case 'addons': {
        const { handleAddons } = await import('./commands/addons');
        await handleAddons(dodoClient, subCommand);
        break;
      }
      case 'refunds': {
        const { handleRefund } = await import('./commands/refund');
        await handleRefund(dodoClient, subCommand);
        break;
      }
      case 'checkout': {
        const { handleCheckout } = await import('./commands/checkout');
        await handleCheckout(dodoClient, subCommand);
        break;
      }
      case 'wh': {
        const { handleWebhook } = await import('./commands/webhook');
        await handleWebhook(subCommand, { apiKey, client: dodoClient });
        break;
      }
      default:
        renderHelp();
    }
  }
} catch (e: any) {
  if (e?.name === 'ExitPromptError') process.exit(0);
  console.error('An unexpected error occurred:', e);
  process.exit(1);
}
