/**
 * Slash-command help catalogue. 1:1 port of `src/ui/ink/help-structure.ts`.
 * Edits made here MUST be mirrored to the Ink version through the cutover
 * (Phase 7) -- both TUIs render the same /help during the dual-stack period.
 */

export interface HelpItem {
  command: string;
  args?: string;
  description: string;
}

export interface HelpGroup {
  heading: string;
  items: HelpItem[];
}

export const HELP_GROUPS: HelpGroup[] = [
  {
    heading: 'PRODUCTS',
    items: [
      { command: '/products list', args: '<page>', description: 'List products' },
      { command: '/products create', description: 'Create a product (opens browser)' },
      { command: '/products info', args: '<id>', description: 'Get product detail' },
    ],
  },
  {
    heading: 'PAYMENTS',
    items: [
      { command: '/payments list', args: '<page>', description: 'List payments' },
      { command: '/payments info', args: '<id>', description: 'Get payment detail' },
    ],
  },
  {
    heading: 'CUSTOMERS',
    items: [
      { command: '/customers list', args: '<page>', description: 'List customers' },
      { command: '/customers create', description: 'Create a customer' },
      { command: '/customers update', args: '<id>', description: 'Update a customer' },
    ],
  },
  {
    heading: 'DISCOUNTS',
    items: [
      { command: '/discounts list', args: '<page>', description: 'List discounts' },
      { command: '/discounts create', description: 'Create a discount' },
      { command: '/discounts delete', args: '<id>', description: 'Delete a discount' },
    ],
  },
  {
    heading: 'LICENCES',
    items: [{ command: '/licences list', args: '<page>', description: 'List licences' }],
  },
  {
    heading: 'ADDONS',
    items: [
      { command: '/addons list', args: '<page>', description: 'List addons' },
      { command: '/addons create', description: 'Create an addon (opens browser)' },
      { command: '/addons info', args: '<id>', description: 'Get addon detail' },
    ],
  },
  {
    heading: 'REFUNDS',
    items: [
      { command: '/refunds list', args: '<page>', description: 'List refunds' },
      { command: '/refunds info', args: '<id>', description: 'Get refund detail' },
    ],
  },
  {
    heading: 'CHECKOUT',
    items: [{ command: '/checkout new', description: 'Create a checkout session' }],
  },
  {
    heading: 'WEBHOOKS',
    items: [
      { command: '/wh listen', description: 'Listen to live webhook events' },
      { command: '/wh trigger', description: 'Trigger a webhook event offline' },
    ],
  },
  {
    heading: 'AI',
    items: [{ command: '/ai', args: '<query>', description: 'Ask a question about your data' }],
  },
  {
    heading: 'AUTH',
    items: [
      { command: '/login', description: 'Sign in with a Dodo Payments API key' },
      { command: '/logout', description: 'Sign out from stored accounts' },
    ],
  },
  {
    heading: 'SESSION',
    items: [
      { command: '/clear', description: 'Clear message history' },
      { command: '/update', description: 'Update the CLI to the latest version' },
      { command: '/exit', description: 'Quit the CLI' },
    ],
  },
];

export const HELP_FOOTER = 'Tab to autocomplete  ·  ↑/↓ for history  ·  Esc to exit';
