#!/usr/bin/env node
import { renderHelp } from './ui';
import { version } from '../package.json';
import { configExists, resolveCredentials } from './utils/auth';
import { defaultContext } from './tui/DefaultContext';

const positional = process.argv.slice(2).filter((a) => !a.startsWith('--'));
const flags = process.argv.slice(2).filter((a) => a.startsWith('--'));

const category = positional[0];
const subCommand = positional[1];
const extraArgs = positional.slice(2);

if (
  category === '--version' ||
  category === '-v' ||
  flags.includes('--version') ||
  flags.includes('-v')
) {
  console.log(`v${version}`);
  process.exit(0);
}

if (process.stdout.isTTY && !category) {
  const { mountTui } = await import('./tui/bootstrap');
  await mountTui();
} else {
  process.on('SIGINT', () => process.exit(0));
  process.on('SIGTERM', () => process.exit(0));
  try {
    if (category === 'login') {
      const { handleLogin } = await import('./commands/login');
      await handleLogin(defaultContext);
    } else if (category === 'logout') {
      const { handleLogout } = await import('./commands/logout');
      await handleLogout(defaultContext);
    } else if (category === 'wh' && subCommand === 'trigger') {
      const { handleWebhook } = await import('./commands/webhook');
      await handleWebhook(subCommand, defaultContext);
    } else if (!(await configExists())) {
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
        'Please login with `dodo login` to use authenticated commands!',
      );
      process.exit(0);
    } else {
      const { apiKey, mode } = await resolveCredentials(defaultContext);
      const { default: DodoPayments } = await import('dodopayments');

      const dodoClient = new DodoPayments({
        bearerToken: apiKey,
        environment: mode,
      });

      switch (category) {
        case 'products': {
          const { handleProducts } = await import('./commands/products');
          await handleProducts(dodoClient, subCommand, defaultContext, extraArgs);
          break;
        }
        case 'payments': {
          const { handlePayments } = await import('./commands/payments');
          await handlePayments(dodoClient, subCommand, defaultContext, extraArgs);
          break;
        }
        case 'customers': {
          const { handleCustomers } = await import('./commands/customers');
          await handleCustomers(dodoClient, subCommand, defaultContext, extraArgs);
          break;
        }
        case 'discounts': {
          const { handleDiscounts } = await import('./commands/discounts');
          await handleDiscounts(dodoClient, subCommand, defaultContext, extraArgs);
          break;
        }
        case 'licences': {
          const { handleLicences } = await import('./commands/licences');
          await handleLicences(dodoClient, subCommand, defaultContext, extraArgs);
          break;
        }
        case 'addons': {
          const { handleAddons } = await import('./commands/addons');
          await handleAddons(dodoClient, subCommand, defaultContext, extraArgs);
          break;
        }
        case 'refunds': {
          const { handleRefund } = await import('./commands/refund');
          await handleRefund(dodoClient, subCommand, defaultContext, extraArgs);
          break;
        }
        case 'checkout': {
          const { handleCheckout } = await import('./commands/checkout');
          await handleCheckout(dodoClient, subCommand, defaultContext, extraArgs);
          break;
        }
        case 'wh': {
          const { handleWebhook } = await import('./commands/webhook');
          await handleWebhook(subCommand, defaultContext, { apiKey, client: dodoClient });
          break;
        }
        default:
          renderHelp();
      }
    }
  } catch (e: any) {
    if (e?.name === 'ExitPromptError') process.exit(0);
    console.error('Unexpected error.', e);
    process.exit(1);
  }
}
