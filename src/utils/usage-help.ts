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
    { command: 'trigger', description: 'Trigger a webhook event offline' },
  ],
  checkout: [{ command: 'new', description: 'Create a checkout session' }],
};
