import fs from 'fs';
import path from 'path';
import { execSync } from 'child_process';
import chalk from 'chalk';
import { colors } from '../../tui/theme';
import type { CommandContext } from '../../tui/CommandContext';

function ok(msg: string, ctx?: CommandContext) {
  if (ctx) {
    ctx.addBlock({ type: 'success', message: msg });
  } else {
    console.log(`${chalk.hex(colors.success)('✔')}  ${msg}`);
  }
}

function info(msg: string, ctx?: CommandContext) {
  if (ctx) {
    ctx.addBlock({ type: 'info', message: msg });
  } else {
    console.log(`${chalk.hex(colors.info)('ℹ')}  ${msg}`);
  }
}

function warn(msg: string, ctx?: CommandContext) {
  if (ctx) {
    ctx.addBlock({ type: 'info', message: chalk.yellow(`⚠  ${msg}`) });
  } else {
    console.log(`${chalk.hex(colors.warning)('⚠')}  ${msg}`);
  }
}

function error(msg: string, ctx?: CommandContext) {
  if (ctx) {
    ctx.addBlock({ type: 'error', message: msg });
  } else {
    console.log(`${chalk.hex(colors.error)('✖')}  ${msg}`);
  }
}

function dim(msg: string, ctx?: CommandContext) {
  if (ctx) {
    ctx.addBlock({ type: 'info', message: chalk.dim(msg) });
  } else {
    console.log(chalk.dim(msg));
  }
}

function divider(ctx?: CommandContext) {
  if (ctx) {
    ctx.addBlock({ type: 'empty' });
  } else {
    console.log(chalk.dim('─'.repeat(52)));
  }
}

function getPackageManager(): 'bun' | 'pnpm' | 'yarn' | 'npm' {
  const cwd = process.cwd();
  if (fs.existsSync(path.join(cwd, 'bun.lockb')))       return 'bun';
  if (fs.existsSync(path.join(cwd, 'bun.lock')))        return 'bun';
  if (fs.existsSync(path.join(cwd, 'pnpm-lock.yaml')))  return 'pnpm';
  if (fs.existsSync(path.join(cwd, 'yarn.lock')))       return 'yarn';
  return 'npm';
}

function hasSrcDir(): boolean {
  return fs.existsSync(path.join(process.cwd(), 'src'));
}

function isTypeScriptProject(): boolean {
  return fs.existsSync(path.join(process.cwd(), 'tsconfig.json'));
}

function writeFile(relPath: string, content: string, ctx?: CommandContext) {
  const abs = path.join(process.cwd(), relPath);

  fs.mkdirSync(path.dirname(abs), { recursive: true });

  const existed = fs.existsSync(abs);

  if (existed) {
    warn(`Skipped existing file: ${relPath}`, ctx);
    return;
  }

  fs.writeFileSync(abs, content, 'utf8');

  if (ctx) {
    ctx.addBlock({
      type: 'success',
      message: `Created file: ${relPath}`,
    });
  } else {
    console.log(`  ${chalk.hex(colors.success)('+')} ${chalk.dim(relPath)}`);
  }
}

function installDeps(packages: string, ctx?: CommandContext): Promise<boolean> | boolean {
  const pm = getPackageManager();
  const cmd = pm === 'npm' ? `npm install ${packages}` : `${pm} add ${packages}`;
  if (ctx) {
    const spinnerId = ctx.addBlock({
      type: 'spinner',
      label: `Installing dependencies: ${cmd}…`,
    });
    return new Promise<boolean>((resolve) => {
      const { exec } = require('child_process');
      exec(cmd, (err: any) => {
        ctx.removeBlock(spinnerId);
        if (err) {
          ctx.addBlock({
            type: 'error',
            message: `Dependency install failed — run \`${cmd}\` manually.`,
          });
          resolve(false);
        } else {
          ctx.addBlock({
            type: 'success',
            message: 'Dependencies installed successfully.',
          });
          resolve(true);
        }
      });
    });
  } else {
    info(`Running: ${chalk.dim(cmd)}`);
    try {
      execSync(cmd, { stdio: 'inherit' });
      return true;
    } catch {
      warn('Dependency install failed — run the command above manually.');
      return false;
    }
  }
}

function hasEnvKey(content: string, key: string): boolean {
  const escapedKey = key.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
  return new RegExp(`^\\s*${escapedKey}=`, 'm').test(content);
}

function appendEnv(vars: Record<string, string>, ctx?: CommandContext) {
  const envPath = path.join(process.cwd(), '.env');

  const existing = fs.existsSync(envPath)
    ? fs.readFileSync(envPath, 'utf8')
    : '';

  const missing = Object.entries(vars).filter(
    ([key]) => !hasEnvKey(existing, key),
  );

  if (missing.length === 0) {
    warn('.env already contains all Dodo vars, skipping', ctx);
    return;
  }

  const lines = missing.map(([key, val]) => `${key}=${val}`).join('\n');
  const sep = existing.length > 0 && !existing.endsWith('\n') ? '\n' : '';

  if (!fs.existsSync(envPath)) {
    fs.writeFileSync(envPath, lines + '\n', 'utf8');
    ok('Created .env', ctx);
  } else {
    fs.appendFileSync(envPath, sep + lines + '\n', 'utf8');
    ok(`Appended ${missing.length} Dodo var(s) to existing .env`, ctx);
  }
}


const boilerplates = {
  nextjs: {
    checkoutTs: `import { Checkout } from "@dodopayments/nextjs";

export const GET = Checkout({
  bearerToken: process.env.DODO_PAYMENTS_API_KEY,
  returnUrl: process.env.DODO_PAYMENTS_RETURN_URL,
  environment: process.env.DODO_PAYMENTS_ENVIRONMENT as "live_mode" | "test_mode" | undefined,
  type: "static",
});

export const POST = Checkout({
  bearerToken: process.env.DODO_PAYMENTS_API_KEY,
  returnUrl: process.env.DODO_PAYMENTS_RETURN_URL,
  environment: process.env.DODO_PAYMENTS_ENVIRONMENT as "live_mode" | "test_mode" | undefined,
  type: "session",
});
`,
    checkoutJs: `import { Checkout } from "@dodopayments/nextjs";

export const GET = Checkout({
  bearerToken: process.env.DODO_PAYMENTS_API_KEY,
  returnUrl: process.env.DODO_PAYMENTS_RETURN_URL,
  environment: process.env.DODO_PAYMENTS_ENVIRONMENT,
  type: "static",
});

export const POST = Checkout({
  bearerToken: process.env.DODO_PAYMENTS_API_KEY,
  returnUrl: process.env.DODO_PAYMENTS_RETURN_URL,
  environment: process.env.DODO_PAYMENTS_ENVIRONMENT,
  type: "session",
});
`,
    portalTs: `import { CustomerPortal } from "@dodopayments/nextjs";

export const GET = CustomerPortal({
  bearerToken: process.env.DODO_PAYMENTS_API_KEY,
  environment: process.env.DODO_PAYMENTS_ENVIRONMENT as "live_mode" | "test_mode" | undefined,
});
`,
    portalJs: `import { CustomerPortal } from "@dodopayments/nextjs";

export const GET = CustomerPortal({
  bearerToken: process.env.DODO_PAYMENTS_API_KEY,
  environment: process.env.DODO_PAYMENTS_ENVIRONMENT,
});
`,
    webhook: `import { Webhooks } from "@dodopayments/nextjs";

export const POST = Webhooks({
  webhookKey: process.env.DODO_PAYMENTS_WEBHOOK_KEY,
  onPayload: async (payload) => {
    console.log("Webhook received:", payload);
  },
});
`,
  },
  express: {
    checkoutTs: `import { checkoutHandler } from '@dodopayments/express';
import express from 'express';

const router = express.Router();

router.get('/checkout', checkoutHandler({
  bearerToken: process.env.DODO_PAYMENTS_API_KEY,
  returnUrl: process.env.DODO_PAYMENTS_RETURN_URL,
  environment: process.env.DODO_PAYMENTS_ENVIRONMENT as "live_mode" | "test_mode" | undefined,
  type: "static",
}));

export default router;
`,
    checkoutJs: `import { checkoutHandler } from '@dodopayments/express';
import express from 'express';

const router = express.Router();

router.get('/checkout', checkoutHandler({
  bearerToken: process.env.DODO_PAYMENTS_API_KEY,
  returnUrl: process.env.DODO_PAYMENTS_RETURN_URL,
  environment: process.env.DODO_PAYMENTS_ENVIRONMENT,
  type: "static",
}));

export default router;
`,
    portalTs: `import { CustomerPortal } from "@dodopayments/express";
import express from 'express';

const router = express.Router();

router.get('/customer-portal', CustomerPortal({
  bearerToken: process.env.DODO_PAYMENTS_API_KEY,
  environment: process.env.DODO_PAYMENTS_ENVIRONMENT as "live_mode" | "test_mode" | undefined,
}));

export default router;
`,
    portalJs: `import { CustomerPortal } from "@dodopayments/express";
import express from 'express';

const router = express.Router();

router.get('/customer-portal', CustomerPortal({
  bearerToken: process.env.DODO_PAYMENTS_API_KEY,
  environment: process.env.DODO_PAYMENTS_ENVIRONMENT,
}));

export default router;
`,
    webhook: `import { Webhooks } from "@dodopayments/express";
import express from 'express';

const router = express.Router();

router.post('/webhook', Webhooks({
  webhookKey: process.env.DODO_PAYMENTS_WEBHOOK_KEY,
  onPayload: async (payload) => {
    console.log("Webhook received:", payload);
  },
}));

export default router;
`,
  },
};

export type BetterAuthPlugin = 'checkout' | 'portal' | 'usage' | 'webhooks';
export const ALL_PLUGINS: BetterAuthPlugin[] = ['checkout', 'portal', 'usage', 'webhooks'];

function generateBetterAuthServer(plugins: BetterAuthPlugin[]): string {
  const pluginBodies = plugins.map((p) => {
    if (p === 'checkout') {
      return `checkout({
          products: [
            {
              productId: "pdt_xxxxxxxxxxxxxxxxxxxxx",
              slug: "premium-plan",
            },
          ],
          successUrl: "/dashboard/success",
          authenticatedUsersOnly: true,
        })`;
    }
    if (p === 'webhooks') {
      return `webhooks({
          webhookKey: process.env.DODO_PAYMENTS_WEBHOOK_SECRET!,
          onPayload: async (payload) => {
            console.log("Received webhook:", payload.event_type);
          },
        })`;
    }
    return `${p}()`;
  });

  return `import { betterAuth } from "better-auth";
import { ${['dodopayments', ...plugins].join(', ')} } from "@dodopayments/better-auth";
import DodoPayments from "dodopayments";

export const dodoPayments = new DodoPayments({
  bearerToken: process.env.DODO_PAYMENTS_API_KEY!,
  environment: (process.env.DODO_PAYMENTS_ENVIRONMENT || "test_mode") as "live_mode" | "test_mode",
});

export const { auth, endpoints, client } = betterAuth({
  plugins: [
    dodopayments({
      client: dodoPayments,
      createCustomerOnSignUp: true,
      use: [
        ${pluginBodies.join(',\n        ')}
      ],
    }),
  ],
});
`;
}

const betterAuthClientBoilerplate = `import { createAuthClient } from "better-auth/react";
import { dodopaymentsClient } from "@dodopayments/better-auth/client";

export const authClient = createAuthClient({
  baseURL: process.env.BETTER_AUTH_URL || "http://localhost:3000",
  plugins: [dodopaymentsClient()],
});
`;


async function scaffoldNextjs(ctx?: CommandContext) {
  if (ctx) {
    ctx.addBlock({ type: 'info', message: 'Scaffolding Next.js App Router billing routes…' });
  } else {
    console.log('');
    console.log(chalk.bold('Scaffolding Next.js App Router billing routes…'));
    divider();
  }

  const installed = await installDeps('@dodopayments/nextjs', ctx);
  if (!installed) {
    throw new Error('Aborting scaffold — please install dependencies first.');
  }
  if (!ctx) console.log('');

  const isTs = isTypeScriptProject();
  if (!isTs) {
    warn('No tsconfig.json found — writing plain JS files (no type assertions).', ctx);
  }

  const appRoot = hasSrcDir() ? 'src/app' : 'app';
  const ext = isTs ? 'ts' : 'js';

  writeFile(
    `${appRoot}/checkout/route.${ext}`,
    isTs ? boilerplates.nextjs.checkoutTs : boilerplates.nextjs.checkoutJs,
    ctx,
  );
  writeFile(
    `${appRoot}/customer-portal/route.${ext}`,
    isTs ? boilerplates.nextjs.portalTs : boilerplates.nextjs.portalJs,
    ctx,
  );
  writeFile(
    `${appRoot}/api/webhook/dodo-payments/route.${ext}`,
    boilerplates.nextjs.webhook,
    ctx,
  );

  appendEnv({
    DODO_PAYMENTS_API_KEY: '"your_api_key_here"',
    DODO_PAYMENTS_RETURN_URL: '"http://localhost:3000/success"',
    DODO_PAYMENTS_ENVIRONMENT: '"test_mode"',
    DODO_PAYMENTS_WEBHOOK_KEY: '"your_webhook_key_here"',
  }, ctx);

  if (ctx) {
    const md = `### Next.js billing routes ready!

**Routes created:**
* \`GET/POST\`  \`/checkout\`
* \`GET\`       \`/customer-portal\`
* \`POST\`      \`/api/webhook/dodo-payments\``;
    ctx.addBlock({ type: 'markdown', text: md });
  } else {
    divider();
    console.log('');
    ok(chalk.bold('Next.js billing routes ready!'));
    console.log('');
    info('Routes created:');
    dim('  GET/POST  /checkout');
    dim('  GET       /customer-portal');
    dim('  POST      /api/webhook/dodo-payments');
    console.log('');
  }
}

async function scaffoldExpress(ctx?: CommandContext) {
  if (ctx) {
    ctx.addBlock({ type: 'info', message: 'Scaffolding Express billing routes…' });
  } else {
    console.log('');
    console.log(chalk.bold('Scaffolding Express billing routes…'));
    divider();
  }

  const installed = await installDeps('@dodopayments/express', ctx);
  if (!installed) {
    throw new Error('Aborting scaffold — please install dependencies first.');
  }
  if (!ctx) console.log('');

  const isTs = isTypeScriptProject();
  if (!isTs) {
    warn('No tsconfig.json found — writing plain JS files (no type assertions).', ctx);
  }

  const routesDir = hasSrcDir() ? 'src/routes' : 'routes';
  const ext = isTs ? 'ts' : 'js';

  writeFile(
    `${routesDir}/checkout.${ext}`,
    isTs ? boilerplates.express.checkoutTs : boilerplates.express.checkoutJs,
    ctx,
  );
  writeFile(
    `${routesDir}/customerPortal.${ext}`,
    isTs ? boilerplates.express.portalTs : boilerplates.express.portalJs,
    ctx,
  );
  writeFile(`${routesDir}/webhook.${ext}`, boilerplates.express.webhook, ctx);

  appendEnv({
    DODO_PAYMENTS_API_KEY: '"your_api_key_here"',
    DODO_PAYMENTS_RETURN_URL: '"http://localhost:3000/success"',
    DODO_PAYMENTS_ENVIRONMENT: '"test_mode"',
    DODO_PAYMENTS_WEBHOOK_KEY: '"your_webhook_key_here"',
  }, ctx);

  if (ctx) {
    const md = `### Express billing routes ready!

**Mount in your app:**
\`\`\`typescript
import checkoutRouter       from './${routesDir}/checkout';
import customerPortalRouter from './${routesDir}/customerPortal';
import webhookRouter        from './${routesDir}/webhook';

app.use(checkoutRouter);
app.use(customerPortalRouter);
app.use(webhookRouter);
\`\`\``;
    ctx.addBlock({ type: 'markdown', text: md });
  } else {
    divider();
    console.log('');
    ok(chalk.bold('Express billing routes ready!'));
    console.log('');
    info('Mount in your app:');
    dim(`  import checkoutRouter       from './${routesDir}/checkout';`);
    dim(`  import customerPortalRouter from './${routesDir}/customerPortal';`);
    dim(`  import webhookRouter        from './${routesDir}/webhook';`);
    dim('');
    dim('  app.use(checkoutRouter);');
    dim('  app.use(customerPortalRouter);');
    dim('  app.use(webhookRouter);');
    console.log('');
  }
}

async function scaffoldBetterAuth(plugins: BetterAuthPlugin[] = ALL_PLUGINS, ctx?: CommandContext) {
  if (ctx) {
    ctx.addBlock({ type: 'info', message: `Scaffolding Better-Auth plugin (plugins: ${plugins.join(', ')})…` });
  } else {
    console.log('');
    console.log(chalk.bold(`Scaffolding Better-Auth plugin (plugins: ${plugins.join(', ')})…`));
    divider();
  }

  const installed = await installDeps('@dodopayments/better-auth better-auth dodopayments zod', ctx);
  if (!installed) {
    throw new Error('Aborting scaffold — please install dependencies first.');
  }
  if (!ctx) console.log('');

  const libDir = hasSrcDir() ? 'src/lib' : 'lib';

  writeFile(`${libDir}/auth.ts`,        generateBetterAuthServer(plugins), ctx);
  writeFile(`${libDir}/auth-client.ts`, betterAuthClientBoilerplate, ctx);

  const envVars: Record<string, string> = {
    DODO_PAYMENTS_API_KEY: '"your_api_key_here"',
    DODO_PAYMENTS_ENVIRONMENT: '"test_mode"',
    BETTER_AUTH_URL: '"http://localhost:3000"',
    BETTER_AUTH_SECRET: '"your_better_auth_secret_32_chars"',
  };
  if (plugins.includes('webhooks')) {
    envVars['DODO_PAYMENTS_WEBHOOK_SECRET'] = '"your_webhook_secret_here"';
  }
  appendEnv(envVars, ctx);

  if (ctx) {
    const md = `### Better-Auth plugin config ready!

**Files created:**
* \`${libDir}/auth.ts\` — server auth instance
* \`${libDir}/auth-client.ts\` — client-side hooks`;
    ctx.addBlock({ type: 'markdown', text: md });
  } else {
    divider();
    console.log('');
    ok(chalk.bold('Better-Auth plugin config ready!'));
    console.log('');
    info('Files created:');
    dim(`  ${libDir}/auth.ts        — server auth instance`);
    dim(`  ${libDir}/auth-client.ts — client-side hooks`);
    console.log('');
  }
}

function printUsage(ctx?: CommandContext) {
  if (ctx) {
    const md = `### /init — scaffold Dodo Payments into your project

**Usage:**
  \`/init <framework>\`

**Available scaffolds:**
* **nextjs** — Next.js App Router billing routes
* **express** — Express server billing routes
* **better-auth** — Better-Auth plugin configuration

**Better-Auth plugin options (comma-separated, default: all):**
  \`checkout\`, \`portal\`, \`usage\`, \`webhooks\`

**Examples:**
* \`/init nextjs\`
* \`/init better-auth\`
* \`/init better-auth checkout,portal\``;

    ctx.addBlock({ type: 'markdown', text: md });
    return;
  }

  console.log('');
  console.log(`${chalk.bold(chalk.hex(colors.info)('dodo init'))} — scaffold Dodo Payments into your project`);
  console.log('');
  console.log('Usage:');
  dim('  dodo init <framework>');
  console.log('');
  console.log('Available scaffolds:');
  console.log(`  ${chalk.hex(colors.success)('nextjs')}        Next.js App Router billing routes`);
  console.log(`  ${chalk.hex(colors.success)('express')}       Express server billing routes`);
  console.log(`  ${chalk.hex(colors.success)('better-auth')}   Better-Auth plugin configuration`);
  console.log('');
  info('Better-Auth plugin options (comma-separated, default: all):');
  dim(`  ${ALL_PLUGINS.join(', ')}`);
  console.log('');
  info('Examples:');
  dim('  dodo init nextjs');
  dim('  dodo init better-auth');
  dim('  dodo init better-auth checkout,portal');
  dim('  dodo init better-auth all');
  console.log('');
}

export async function handleInitCommand(subCommand?: string, pluginsArg?: string, ctx?: CommandContext) {
  switch (subCommand) {
    case 'nextjs':
      await scaffoldNextjs(ctx);
      break;

    case 'express':
      await scaffoldExpress(ctx);
      break;

    case 'better-auth': {
      let plugins: BetterAuthPlugin[] = ALL_PLUGINS;

      if (pluginsArg && pluginsArg !== 'all') {
        const requested = pluginsArg.split(',').map((p) => p.trim());

        const invalid = requested.filter((p) => !ALL_PLUGINS.includes(p as BetterAuthPlugin));
        if (invalid.length > 0) {
          error(`Unknown plugin(s): ${invalid.join(', ')}`, ctx);
          info(`Available plugins: ${ALL_PLUGINS.join(', ')}`, ctx);
          throw new Error(`Unknown plugin(s): ${invalid.join(', ')}`);
        }

        plugins = requested as BetterAuthPlugin[];
      }

      await scaffoldBetterAuth(plugins, ctx);
      break;
    }

    default:
      printUsage(ctx);
  }
}