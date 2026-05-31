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
    if (category && subCommand) {
      const { isInteractiveOnly } = await import('./utils/usage-help');
      if (isInteractiveOnly(category, subCommand)) {
        console.error(`\`dodo ${category} ${subCommand}\` is interactive — open the TUI with \`dodo\`.`);
        process.exit(0);
      }
    }

    if (category === 'login') {
      const { handleLogin } = await import('./commands/login');
      await handleLogin(defaultContext, subCommand, extraArgs[0]);
    } else if (category === 'logout') {
      const { handleLogout } = await import('./commands/logout');
      await handleLogout(defaultContext, undefined, subCommand);
    } else if (category === 'wh' && subCommand === 'trigger') {
      const { handleWebhook } = await import('./commands/webhook');
      await handleWebhook(subCommand, defaultContext, undefined, extraArgs);

    
    } else if (category === 'init') {
      const pluginFlag = flags.find((f) => f.startsWith('--plugins='));
    
      const pluginsArg = extraArgs[0] ?? (pluginFlag ? pluginFlag.split('=')[1] : undefined);

      const { handleInitCommand } = await import('./commands/init');
      await handleInitCommand(subCommand, pluginsArg);

    } else if (!(await configExists())) {
      if (category && !subCommand) {
        renderHelp(category);
        console.log();
        const { categoryNotes } = await import('./utils/usage-help');
        const note = categoryNotes[category];
        if (note) {
          console.log(note);
          process.exit(0);
        }
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
          await handleWebhook(subCommand, defaultContext, { apiKey, client: dodoClient }, extraArgs);
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