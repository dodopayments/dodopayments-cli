import type { CommandContext } from '../tui/CommandContext';

type UsageCommand = {
  command: string;
  description: string;
};

// For help commands
export const standaloneUsage: UsageCommand[] = [
  {
    command: 'login',
    description: 'Login with a Dodo Payments API key',
  },
  {
    command: 'logout',
    description: 'Logout from stored Dodo Payments accounts',
  },
];

export const usage: Record<string, UsageCommand[]> = {
  products: [
    { command: 'list', description: 'List your products' },
    { command: 'create', description: 'Create a new product' },
    { command: 'info <id>', description: 'Get info about a product' },
  ],
  payments: [
    { command: 'list', description: 'List your payments' },
    { command: 'info <id>', description: 'Information about a payment' },
  ],
  customers: [
    { command: 'list', description: 'List your customers' },
    { command: 'create', description: 'Create a customer' },
    { command: 'update <id>', description: 'Update a customer' },
  ],
  discounts: [
    { command: 'list', description: 'List your discounts' },
    { command: 'create', description: 'Create a discount' },
    { command: 'delete <id>', description: 'Remove a discount' },
  ],
  licences: [{ command: 'list', description: 'List licences' }],
  addons: [
    { command: 'create', description: 'Create an addon' },
    { command: 'list', description: 'List addons' },
    { command: 'info <id>', description: 'Get addon info' },
  ],
  refunds: [
    { command: 'list', description: 'List refunds' },
    { command: 'info <id>', description: 'Get refund info' },
  ],
  wh: [
    {
      command: 'listen <url>',
      description: 'Listen to webhook events directly from Dodo Payments',
    },
    { command: 'trigger <event> <url>', description: 'Trigger a webhook event offline' },
  ],
  checkout: [{ command: 'new', description: 'Create a checkout session' }],
};

export const categoryNotes: Record<string, string> = {
  wh: 'Run `dodo wh trigger` without logging in, or `dodo login` to use `dodo wh listen`.',
};

export function unknownSubcommand(
  ctx: CommandContext,
  category: string,
  subCommand: string | undefined,
): void {
  const invocation = ctx.invocation === 'cli' ? `dodo ${category}` : `/${category}`;

  if (subCommand) {
    ctx.addBlock({
      type: 'error',
      message: `Unknown subcommand '${subCommand}' for ${invocation}.`,
    });
  } else {
    ctx.addBlock({ type: 'error', message: `Subcommand required for ${invocation}.` });
  }

  const commands = usage[category];
  if (commands) {
    const lines = [
      'Usage:',
      ...commands.map((c) => `  ${invocation} ${c.command} - ${c.description}`),
    ].join('\n');
    ctx.addBlock({ type: 'info', message: lines });
  }

  const note = categoryNotes[category];
  if (note && ctx.invocation === 'cli') {
    ctx.addBlock({ type: 'info', message: note });
  }
}
