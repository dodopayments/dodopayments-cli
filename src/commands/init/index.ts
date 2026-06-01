import fs from 'fs';
import path from 'path';
import { execSync } from 'child_process';

const c = {
  reset:   '\x1b[0m',
  bold:    '\x1b[1m',
  dim:     '\x1b[2m',
  green:   '\x1b[32m',
  yellow:  '\x1b[33m',
  blue:    '\x1b[34m',
  cyan:    '\x1b[36m',
  red:     '\x1b[31m',
  gray:    '\x1b[90m',
};

function ok(msg: string)    { console.log(`${c.green}✔${c.reset}  ${msg}`); }
function info(msg: string)  { console.log(`${c.cyan}ℹ${c.reset}  ${msg}`); }
function warn(msg: string)  { console.log(`${c.yellow}⚠${c.reset}  ${msg}`); }
function error(msg: string) { console.log(`${c.red}✖${c.reset}  ${msg}`); }
function dim(msg: string)   { console.log(`${c.dim}${msg}${c.reset}`); }
function divider()          { console.log(`${c.dim}${'─'.repeat(52)}${c.reset}`); }

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
  return fs.existsSync(
    path.join(process.cwd(), 'tsconfig.json'),
  );
}

function writeFile(relPath: string, content: string) {
  const abs = path.join(process.cwd(), relPath);

  fs.mkdirSync(path.dirname(abs), { recursive: true });

  const existed = fs.existsSync(abs);

  if (existed) {
    warn(`Skipped existing file: ${relPath}`);
    return;
  }

  fs.writeFileSync(abs, content, 'utf8');

  console.log(
    `  ${c.green}+ ${c.reset}${c.gray}${relPath}${c.reset}`,
  );
}

function installDeps(packages: string): boolean {
  const pm = getPackageManager();
  const cmd = pm === 'npm' ? `npm install ${packages}` : `${pm} add ${packages}`;
  info(`Running: ${c.dim}${cmd}${c.reset}`);
  try {
    execSync(cmd, { stdio: 'inherit' });
    return true;
  } catch {
    warn('Dependency install failed — run the command above manually.');
    return false;
  }
}
function hasEnvKey(content: string, key: string): boolean {
  const escapedKey = key.replace(/[.*+?^${}()|[\]\\]/g, '\\$&');
  return new RegExp(`^\\s*${escapedKey}=`, 'm').test(content);
}

function appendEnv(vars: string) {
  const envPath = path.join(process.cwd(), '.env');

  if (!fs.existsSync(envPath)) {
    fs.writeFileSync(envPath, vars.trimStart() + '\n', 'utf8');
    ok('Created .env');
    return;
  }

  const existing = fs.readFileSync(envPath, 'utf8');

  if (!hasEnvKey(existing, 'DODO_PAYMENTS_API_KEY')) {
    const sep =
      existing.endsWith('\n') || existing.length === 0 ? '' : '\n';

    fs.appendFileSync(
      envPath,
      sep + vars.trimStart() + '\n',
      'utf8',
    );

    ok('Appended Dodo vars to existing .env');
  } else {
    warn('.env already contains DODO_PAYMENTS_API_KEY, skipping');
  }
}

const boilerplates = {
  nextjs: {
    checkout: `import { Checkout } from "@dodopayments/nextjs";

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
    portal: `import { CustomerPortal } from "@dodopayments/nextjs";

export const GET = CustomerPortal({
  bearerToken: process.env.DODO_PAYMENTS_API_KEY,
  environment: process.env.DODO_PAYMENTS_ENVIRONMENT as "live_mode" | "test_mode" | undefined,
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
    checkout: `import { checkoutHandler } from '@dodopayments/express';
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
    portal: `import { CustomerPortal } from "@dodopayments/express";
import express from 'express';

const router = express.Router();

router.get('/customer-portal', CustomerPortal({
  bearerToken: process.env.DODO_PAYMENTS_API_KEY,
  environment: process.env.DODO_PAYMENTS_ENVIRONMENT as "live_mode" | "test_mode" | undefined,
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

function scaffoldNextjs() {
  console.log('');
  console.log(`${c.bold}Scaffolding Next.js App Router billing routes…${c.reset}`);
  divider();

  const installed = installDeps('@dodopayments/nextjs');
  if (!installed) {
    error('Aborting scaffold — please install dependencies first.');
    process.exit(1);
  }
  console.log('');

  const appRoot = hasSrcDir() ? 'src/app' : 'app';

  writeFile(`${appRoot}/checkout/route.ts`,                      boilerplates.nextjs.checkout);
  writeFile(`${appRoot}/customer-portal/route.ts`,               boilerplates.nextjs.portal);
  writeFile(`${appRoot}/api/webhook/dodo-payments/route.ts`,     boilerplates.nextjs.webhook);

  appendEnv(`
DODO_PAYMENTS_API_KEY="your_api_key_here"
DODO_PAYMENTS_RETURN_URL="http://localhost:3000/success"
DODO_PAYMENTS_ENVIRONMENT="test_mode"
DODO_PAYMENTS_WEBHOOK_KEY="your_webhook_key_here"
`);

  divider();
  console.log('');
  ok(`${c.bold}Next.js billing routes ready!${c.reset}`);
  console.log('');
  info('Routes created:');
  dim('  GET/POST  /checkout');
  dim('  GET       /customer-portal');
  dim('  POST      /api/webhook/dodo-payments');
  console.log('');
}

function scaffoldExpress() {
  console.log('');
  console.log(`${c.bold}Scaffolding Express billing routes…${c.reset}`);
  divider();

  const installed = installDeps('@dodopayments/express');
  if (!installed) {
    error('Aborting scaffold — please install dependencies first.');
    process.exit(1);
  }
  console.log('');

  const routesDir = hasSrcDir() ? 'src/routes' : 'routes';

  writeFile(`${routesDir}/checkout.ts`,       boilerplates.express.checkout);
  writeFile(`${routesDir}/customerPortal.ts`, boilerplates.express.portal);
  writeFile(`${routesDir}/webhook.ts`,        boilerplates.express.webhook);

  appendEnv(`
DODO_PAYMENTS_API_KEY="your_api_key_here"
DODO_PAYMENTS_RETURN_URL="http://localhost:3000/success"
DODO_PAYMENTS_ENVIRONMENT="test_mode"
DODO_PAYMENTS_WEBHOOK_KEY="your_webhook_key_here"
`);

  divider();
  console.log('');
  ok(`${c.bold}Express billing routes ready!${c.reset}`);
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

function scaffoldBetterAuth(plugins: BetterAuthPlugin[] = ALL_PLUGINS) {
  console.log('');
  console.log(`${c.bold}Scaffolding Better-Auth plugin (plugins: ${plugins.join(', ')})…${c.reset}`);
  divider();

  const installed = installDeps('@dodopayments/better-auth better-auth dodopayments zod');
  if (!installed) {
    error('Aborting scaffold — please install dependencies first.');
    process.exit(1);
  }
  console.log('');

  const libDir = hasSrcDir() ? 'src/lib' : 'lib';

  writeFile(`${libDir}/auth.ts`,        generateBetterAuthServer(plugins));
  writeFile(`${libDir}/auth-client.ts`, betterAuthClientBoilerplate);

  let envVars = `
DODO_PAYMENTS_API_KEY="your_api_key_here"
DODO_PAYMENTS_ENVIRONMENT="test_mode"
BETTER_AUTH_URL="http://localhost:3000"
BETTER_AUTH_SECRET="your_better_auth_secret_32_chars"
`;
  if (plugins.includes('webhooks')) {
    envVars += `DODO_PAYMENTS_WEBHOOK_SECRET="your_webhook_secret_here"\n`;
  }
  appendEnv(envVars);

  divider();
  console.log('');
  ok(`${c.bold}Better-Auth plugin config ready!${c.reset}`);
  console.log('');
  info('Files created:');
  dim(`  ${libDir}/auth.ts        — server auth instance`);
  dim(`  ${libDir}/auth-client.ts — client-side hooks`);
  console.log('');
}

function printUsage() {
  console.log('');
  console.log(`${c.bold}${c.cyan}dodo init${c.reset} — scaffold Dodo Payments into your project`);
  console.log('');
  console.log('Usage:');
  dim('  dodo init <framework>');
  console.log('');
  console.log('Available scaffolds:');
  console.log(`  ${c.green}nextjs${c.reset}        Next.js App Router billing routes`);
  console.log(`  ${c.green}express${c.reset}       Express server billing routes`);
  console.log(`  ${c.green}better-auth${c.reset}   Better-Auth plugin configuration`);
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

export async function handleInitCommand(subCommand?: string, pluginsArg?: string) {
  switch (subCommand) {
    case 'nextjs':
      scaffoldNextjs();
      break;

    case 'express':
      scaffoldExpress();
      break;

    case 'better-auth': {
      let plugins: BetterAuthPlugin[] = ALL_PLUGINS;

      if (pluginsArg && pluginsArg !== 'all') {
        const requested = pluginsArg.split(',').map((p) => p.trim());

        const invalid = requested.filter((p) => !ALL_PLUGINS.includes(p as BetterAuthPlugin));
        if (invalid.length > 0) {
          error(`Unknown plugin(s): ${invalid.join(', ')}`);
          info(`Available plugins: ${ALL_PLUGINS.join(', ')}`);
          process.exit(1);
        }

        plugins = requested as BetterAuthPlugin[];
      }

      scaffoldBetterAuth(plugins);
      break;
    }

    default:
      printUsage();
  }
}