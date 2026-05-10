/**
 * Slash-command router. Parses a `/cmd subcmd args...` string and dispatches
 * to the matching handler in `src/commands/`. Adds an auth gate for
 * commands that need a Dodo Payments client. Unrecognized commands surface
 * a friendly error block.
 */

import type { CommandContext } from './CommandContext';
import { resolveCredentials } from '../utils/auth';

export async function handleCommand(input: string, ctx: CommandContext, exit: () => void) {
  const trimmedInput = input.trim();
  if (!trimmedInput) return;

  if (!trimmedInput.startsWith('/')) {
    const { handleAI } = await import('../commands/ai');
    await handleAI(trimmedInput, ctx);
    return;
  }

  const parts = trimmedInput.split(' ').filter(Boolean);
  const cmd = parts[0];
  const subCmd = parts[1];
  const extraArgs = parts.slice(2);

  if (cmd === '/help') {
    ctx.addBlock({ type: 'help' });
    return;
  }

  if (cmd === '/update') {
    const { handleUpdate } = await import('../commands/update');
    await handleUpdate(ctx);
    return;
  }

  if (cmd === '/login') {
    const { handleLogin } = await import('../commands/login');
    await handleLogin(ctx);
    return;
  }

  if (cmd === '/logout') {
    const { handleLogout } = await import('../commands/logout');
    await handleLogout(ctx);
    return;
  }

  if (cmd === '/ai') {
    const { handleAI } = await import('../commands/ai');
    await handleAI(parts.slice(1).join(' '), ctx);
    return;
  }

  const { apiKey, mode } = await resolveCredentials(ctx);
  const { default: DodoPayments } = await import('dodopayments');

  const client = new DodoPayments({
    bearerToken: apiKey,
    environment: mode,
  });

  switch (cmd) {
    case '/payments': {
      const { handlePayments } = await import('../commands/payments');
      await handlePayments(client, subCmd, ctx, extraArgs);
      break;
    }
    case '/products': {
      const { handleProducts } = await import('../commands/products');
      await handleProducts(client, subCmd, ctx, extraArgs);
      break;
    }
    case '/customers': {
      const { handleCustomers } = await import('../commands/customers');
      await handleCustomers(client, subCmd, ctx, extraArgs);
      break;
    }
    case '/discounts': {
      const { handleDiscounts } = await import('../commands/discounts');
      await handleDiscounts(client, subCmd, ctx, extraArgs);
      break;
    }
    case '/licences': {
      const { handleLicences } = await import('../commands/licences');
      await handleLicences(client, subCmd, ctx, extraArgs);
      break;
    }
    case '/addons': {
      const { handleAddons } = await import('../commands/addons');
      await handleAddons(client, subCmd, ctx, extraArgs);
      break;
    }
    case '/refunds': {
      const { handleRefund } = await import('../commands/refund');
      await handleRefund(client, subCmd, ctx, extraArgs);
      break;
    }
    case '/checkout': {
      const { handleCheckout } = await import('../commands/checkout');
      await handleCheckout(client, subCmd, ctx, extraArgs);
      break;
    }
    case '/wh': {
      const { handleWebhook } = await import('../commands/webhook');
      await handleWebhook(subCmd, ctx, { apiKey, client });
      break;
    }
    case '/clear':
    case '/exit': {
      if (cmd === '/exit') exit();
      break;
    }
    default:
      ctx.addBlock({ type: 'error', message: `Unknown command: ${cmd}. Try /help.` });
      break;
  }
}
