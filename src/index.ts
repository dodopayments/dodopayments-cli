#!/usr/bin/env node
import { spawnSync } from 'node:child_process';
import { homedir } from 'node:os';
import { join } from 'node:path';
import { fileURLToPath } from 'node:url';
import { renderHelp } from './ui';
import { version } from '../package.json';
import { configExists, resolveCredentials } from './utils/auth';
import { defaultContext } from './tui/DefaultContext';

async function relaunchTuiWithBun(): Promise<never> {
  const executable = process.platform === 'win32' ? 'bun.exe' : 'bun';
  const candidates = [
    executable,
    join(homedir(), '.bun', 'bin', executable),
    '/usr/local/bin/bun',
    '/opt/homebrew/bin/bun',
  ];

  for (const bunPath of candidates) {
    const child = spawnSync(
      bunPath,
      [fileURLToPath(import.meta.url), ...process.argv.slice(2)],
      {
        stdio: 'inherit',
        env: process.env,
      },
    );
    if (!child.error) {
      if (child.signal) process.kill(process.pid, child.signal);
      process.exit(child.status ?? 1);
    }
    if ((child.error as NodeJS.ErrnoException).code !== 'ENOENT') {
      process.stderr.write(`Failed to launch Bun: ${child.error.message}\n`);
      process.exit(1);
    }
  }

  process.stderr.write(
    'The interactive Dodo Payments TUI requires the Bun runtime.\n' +
      '\n' +
      'Install Bun:    https://bun.com/docs/installation\n' +
      'CLI subcommands such as `dodo login` work with Node.js alone.\n' +
      'Or download a standalone binary:\n' +
      '                https://github.com/dodopayments/dodopayments-cli/releases\n',
  );
  process.exit(1);
}

const positional = process.argv.slice(2).filter((a) => !a.startsWith('--'));

if (process.env.DODO_INTERNAL_MCP_CMD === 'remote') {
  await import('mcp-remote/dist/proxy.js');
  // mcp-remote uses floating promises at the module level, so import resolves immediately.
  // We must keep the process alive to allow it to run its stdio server.
  await new Promise(() => {});
} else if (process.env.DODO_INTERNAL_MCP_CMD === 'exec') {
  const { parseCLIOptions } = await import('dodopayments-mcp/options');
  const { launchStdioServer } = await import('dodopayments-mcp/stdio');
  const { configureLogger, getLogger } = await import('dodopayments-mcp/logger');
  const { selectTools } = await import('dodopayments-mcp/server');
  
  const options = parseCLIOptions();
  configureLogger({
      level: options.debug ? 'debug' : 'info',
      pretty: options.logFormat === 'pretty',
  });
  const selectedTools = selectTools(options);
  if (selectedTools.length === 0) {
      getLogger().error('No tools match the provided filters');
      process.exit(1);
  }
  getLogger().info({ tools: selectedTools.map((e: any) => e.tool.name) }, `MCP Server starting with ${selectedTools.length} tools`);
  await launchStdioServer(options);
} else {
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
  if (typeof Bun === 'undefined') await relaunchTuiWithBun();
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
      const { handleInitCommand } = await import('./commands/init');
      await handleInitCommand(subCommand, extraArgs[0]);
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
}