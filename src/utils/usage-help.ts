// For help commands
export const usage: {
    [key: string]: {
        command: string,
        description: string
    }[]
} = {
    products: [
        { command: 'list', description: 'List your products' },
        { command: 'create', description: 'Create a new product' },
        { command: 'info', description: 'Get info about a product' }
    ],
    payments: [
        { command: 'list', description: 'List your payments' },
        { command: 'info', description: 'Information about a payment' },
    ],
    customers: [
        { command: 'list', description: 'List your customers' },
        { command: 'create', description: 'Create a customer' },
        { command: 'update', description: 'Update a customer' },
    ],
    discounts: [
        { command: 'list', description: 'List your discounts' },
        { command: 'create', description: 'Create a discount' },
        { command: 'delete', description: 'Remove a discount' },
    ],
    licences: [
        { command: 'list', description: 'List licences' }
    ],
    addons: [
        { command: 'create', description: 'Create an addon' },
        { command: 'list', description: 'List addons' },
        { command: 'info', description: 'Get addon info' }
    ],
    refunds: [
        { command: 'list', description: 'List refunds' },
        { command: 'info', description: 'Get refund info' }
    ],
    wh: [
        { command: 'listen', description: 'Listen to webhook events directly from Dodo Payments' },
        { command: 'trigger', description: 'Trigger a webhook event' },
    ],
    checkout: [
        { command: 'new', description: 'Create a checkout session' }
    ]
}