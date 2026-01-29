#!/usr/bin/env node
import os from 'node:os';
import path from 'node:path';
import DodoPayments from 'dodopayments';
import { input, select } from '@inquirer/prompts';
import open from 'open';
import { CurrencyToSymbolMap } from './utils/currency-to-symbol-map';
import fs from 'node:fs';

const DodoCliLogo = `
 /$$$$$$$                  /$$                  /$$$$$$  /$$       /$$$$$$
| $$__  $$                | $$                 /$$__  $$| $$      |_  $$_/
| $$  \\ $$  /$$$$$$   /$$$$$$$  /$$$$$$       | $$  \\__/| $$        | $$  
| $$  | $$ /$$__  $$ /$$__  $$ /$$__  $$      | $$      | $$        | $$  
| $$  | $$| $$  \\ $$| $$  | $$| $$  \\ $$      | $$      | $$        | $$  
| $$  | $$| $$  | $$| $$  | $$| $$  | $$      | $$    $$| $$        | $$  
| $$$$$$$/|  $$$$$$/|  $$$$$$$|  $$$$$$/      |  $$$$$$/| $$$$$$$$ /$$$$$$
|_______/  \\______/  \\_______/ \\______/        \\______/ |________/|______/

The CLI to manage Dodo Payments!
`;

// The below is used to check if the error is a Dodo Payments error or not in the API Request
type DodoPaymentsAPIError = {
    error: {
        code: 'NOT_FOUND';
        message: string;
    }
}

function isDodoPaymentsAPIError(e: unknown): e is DodoPaymentsAPIError {
    return (
        typeof e === "object" &&
        e !== null &&
        "error" in e &&
        typeof (e as any).error?.code === "string"
    );
}

// For help commands
const usage: {
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
    wh: [
        { command: '', description: 'Send a webhook event' },
    ]
}


const args = process.argv;
const category = args[2];
const subCommand = args[3];
const homedir = os.homedir();

const RenderHelp = () => {
    console.log(DodoCliLogo);
    // List all available methods
    // todo: Add more comments to make it clear what's being done
    Object.keys(usage).forEach(e => {
        console.log(`Category: ${e}`);
        (usage as any)[e].forEach((y: { command: string, description: string }) => {
            console.log(`dodo ${e} ${y.command} - ${y.description}`)
        });
        // Blank space as a separator
        console.log("");
    });
}

// Added this to the top so that it can bypass all further auth that happens for the login route
if (category === 'login') {
    open('https://app.dodopayments.com/developer/api-keys');
    const API_KEY = await input({ message: 'Enter your Dodo Payments API Key:', required: true });
    const MODE = await select({
        choices: [{ name: "Test Mode", value: 'test_mode' }, { name: "Live Mode", value: 'live_mode' }],
        message: 'Choose the environment:'
    });

    // Initialize the Dodo Payment client to test the API key
    const newDodoClient = new DodoPayments({
        bearerToken: API_KEY,
        environment: (MODE as 'test_mode' | 'live_mode') // as 'test_mode' | 'live_mode' used to bypass ts error
    });

    console.log("Verifying Dodo Payments API Key");

    try {
        // Make this request just to confirm whether API key is correct or not
        await newDodoClient.products.list({ page_size: 1 });
        console.log('Successfully verified your Dodo Payments API Key!');
    } catch (err) {
        console.log("Something went wrong while authenticating:", err);
        process.exit(1);
    };


    console.log("Storing / Updating existing configuration...")
    let existingConfig;
    try {
        existingConfig = Object.create(JSON.parse(fs.readFileSync(path.join(homedir, '.dodopayments', 'api-key'), 'utf-8')));
    } catch {
        existingConfig = {};
    }

    existingConfig[MODE] = API_KEY;
    // Make the ~/.dodopayments directory if it's not present
    if (!fs.existsSync(path.join(homedir, '.dodopayments'))) {
        fs.mkdirSync(path.join(homedir, '.dodopayments'));
    }
    fs.writeFileSync(path.join(homedir, '.dodopayments', 'api-key'), JSON.stringify(existingConfig));

    // Mode will always be either test_mode or live_mode
    existingConfig[MODE] = API_KEY;
    console.log("Setup complete successfully!");
    process.exit(0);
}

// Webhook is managed completely by another file
if (category === 'wh') {
    await import('./dodo-webhooks/index.ts');
}

// Normal functions which require the API key to be present start from here
// Authentication part
// Read the API key config
if (!fs.existsSync(path.join(homedir, '.dodopayments', 'api-key'))) {
    if (category && subCommand) {
        console.log('Please login using `dodo login` command first!');
        process.exit(0);
    } else if (category) {
        if (category in usage) {
            console.log(`Category: ${category}`);
            (usage as any)[category]!.forEach((e: { command: string, description: string }) => {
                console.log(`dodo ${category} ${e.command} - ${e.description}`)
            });
            console.log('\nPlease login using `dodo login` command first!');
        } else {
            RenderHelp();
        }
        process.exit(0);
    } else {
        RenderHelp();
        console.log('Please login using `dodo login` command first!');
        process.exit(0);
    }
}

// Parse the API key config
let existingAPIKeyConfigParsed;
try {
    existingAPIKeyConfigParsed = JSON.parse(fs.readFileSync(path.join(homedir, '.dodopayments', 'api-key'), 'utf-8'));
} catch {
    // Delete API config if something fails with parsing
    fs.rmSync(path.join(homedir, '.dodopayments', 'api-key'), { force: true });
    console.log("Failed to decode API Key configuration. Your config has been reset. Please log in again using `dodo login`");
    process.exit(0);
}


// Final variables
let API_KEY;
let MODE;

// Retrive the keys of the parsed API key config to auto determine the environment if possible.
const existingAPIKeyConfigParsedKeys = Object.keys(existingAPIKeyConfigParsed);

// If there is only one mode auth mehtod then 
if (existingAPIKeyConfigParsedKeys.length === 1) {
    MODE = existingAPIKeyConfigParsedKeys[0]
    API_KEY = existingAPIKeyConfigParsed[MODE!];
}
else {
    // If there are multiple modes (i.e. both test mode & live mode) then prompt the user to select one environment to continue
    MODE = await select({
        choices: [{ name: "Test Mode", value: 'test_mode' }, { name: "Live Mode", value: 'live_mode' }],
        message: 'Choose the environment:'
    });
    API_KEY = existingAPIKeyConfigParsed[MODE];
}

// Initialize the Dodo Payments SDK to be used from now on
const DodoClient = new DodoPayments({
    bearerToken: API_KEY,
    environment: (MODE as 'test_mode' | 'live_mode') // as 'test_mode' | 'live_mode' used to bypass ts error
});

// Continuation of other functions that require api key
if (category === 'products') {
    if (subCommand === 'list') {
        const page = await input({
            message: 'Enter page:',
            default: "1",
            validate: (e => e.trim() !== '')
        });

        const fetchedData = await DodoClient.products.list({ page_number: parseInt(page) - 1, page_size: 100 });
        const table = fetchedData.items.map(e => ({
            name: e.name,
            product_id: e.product_id,
            created_at: new Date(e.created_at).toLocaleString(),
            ...e.is_recurring ? {
                price: `${CurrencyToSymbolMap[e.price_detail!.currency] || (e.price_detail!.currency + ' ')}${(e.price! * 0.01).toFixed(2)} Every ${e.price_detail?.payment_frequency_count} ${e.price_detail?.payment_frequency_interval}`,
            } : {
                price: `${CurrencyToSymbolMap[e.price_detail!.currency] || (e.price_detail!.currency + ' ')}${(e.price! * 0.01).toFixed(2)} (One Time)`,
            },
        }));

        console.table(table);
        console.log("To edit a product, go to https://app.dodopayments.com/products/edit?id={product_id}")
    } else if (subCommand === 'create') {
        open('https://app.dodopayments.com/products/create');
    } else if (subCommand === 'info') {
        try {
            const product_id = await input({
                message: "Enter product ID:",
                validate: (e => e.startsWith('pdt_') || 'Please enter a valid product ID!')
            });

            const info = await DodoClient.products.retrieve(product_id);
            console.table({
                product_id: info.product_id,
                name: info.name,
                ...info.description?.trim() !== '' && { description: info.description },
                created_at: new Date(info.created_at).toLocaleString(),
                updated_at: new Date(info.updated_at).toLocaleString(),
                ...info.is_recurring ? {
                    // .fixed_price for usage based billing
                    price: `${CurrencyToSymbolMap[info.price.currency] || (info.price.currency + ' ')}${((info.price.price || info.price.fixed_price) * 0.01).toFixed(2)} Every ${info.price.payment_frequency_count} ${info.price.payment_frequency_interval}`,
                } : {
                    price: `${CurrencyToSymbolMap[info.price.currency] || (info.price.currency + ' ')}${((info.price.price || info.price.fixed_price) * 0.01).toFixed(2)} (One Time)`,
                },
                tax_category: info.tax_category,
            });
            console.log(`To edit the product, go to https://app.dodopayments.com/products/edit?id=${info.product_id}`)
        } catch (e) {
            if (isDodoPaymentsAPIError(e) && e.error.code === "NOT_FOUND") {
                console.log("Incorrect product ID!");
            } else {
                console.error(e);
            }
        }
    } else {
        usage.products!.forEach(e => {
            console.log(`dodo products ${e.command} - ${e.description}`)
        });
    }
} else if (category === 'payments') {
    if (subCommand === 'list') {
        const page = await input({
            message: 'Enter page:',
            default: "1",
            validate: (e => e.trim() !== '')
        });
        const payments = (await DodoClient.payments.list({ page_number: parseInt(page) - 1, page_size: 100 })).items;
        const paymentsTable = payments.map(payment => {
            return {
                'payment id': payment.payment_id,
                'created at': new Date(payment.created_at).toLocaleString(),
                'subscription id': payment.subscription_id,
                'total amount': `${CurrencyToSymbolMap[payment.currency] || (payment.currency + ' ')}${(payment.total_amount * 0.01).toFixed(2)}`,
                status: payment.status
            };
        });
        console.table(paymentsTable);
        console.log("To view a payment, go to https://app.dodopayments.com/transactions/payments/{payment_id}")
    } else if (subCommand === 'info') {
        try {
            const payment_id = 'pay_0NWiGvZPWxeWeNWISbfat';
            const payment_info = await DodoClient.payments.retrieve(payment_id);
            console.log(payment_info);
            const payment_table = {
                'payment id': payment_info.payment_id,
                status: payment_info.status,
                'total amount': `${CurrencyToSymbolMap[payment_info.currency] || payment_info.currency + ' '}${(payment_info.total_amount * 0.01).toFixed(2)}`,
                'payment method': payment_info.payment_method,
                createdAt: new Date(payment_info.created_at).toLocaleString(),
                customer: payment_info.customer.customer_id,
                'customer email': payment_info.customer.email,
                ...payment_info.subscription_id && {
                    'subscription id': payment_info.subscription_id
                },
                'billing address street': `${payment_info.billing.street}`,
                'billing address state': `${payment_info.billing.state}`,
                'billing address city': `${payment_info.billing.city}`,
                'billing address country': `${payment_info.billing.country}`,
                'billing address zipcode': `${payment_info.billing.zipcode}`,
            }
            console.table(payment_table);
            console.log(`To view the payment, go to https://app.dodopayments.com/transactions/payments/${payment_info.payment_id}`)
        } catch (e) {
            if (isDodoPaymentsAPIError(e) && e.error.code === 'NOT_FOUND') {
                console.log("Incorrect payment ID!")
            } else {
                console.error(e);
            }
        }
    } else {
        usage.payments!.forEach(e => {
            console.log(`dodo payments ${e.command} - ${e.description}`)
        });
    }
} else if (category === 'customers') {
    if (subCommand === 'list') {
        const page = await input({
            message: 'Enter page:',
            default: "1",
            validate: (e => e.trim() !== '')
        });
        console.table((await DodoClient.customers.list({ page_number: parseInt(page) - 1, page_size: 100 })).items, ['customer_id', 'name', 'email', 'phone_number']);
    } else if (subCommand === 'create') {
        const name = await input({
            message: "Enter Name: ",
            validate: (e => e.trim() !== '')
        });

        const email = await input({
            message: "Enter Email: ",
            validate: (e => e.trim() !== '')
        });

        const phone = await input({
            message: "Enter Phone Number: ",
        });

        const creation = await DodoClient.customers.create({
            name,
            email,
            phone_number: phone.trim() !== '' ? phone : null
        });

        console.log('Customer Successfully Created!');
        console.table([creation], ['customer_id', 'name', 'email', 'phone_number']);
    } else if (subCommand === 'update') {
        const customer_id = await input({
            message: "Enter customer ID:",
            validate: (e => e.startsWith('cus_') || 'Please enter a valid customer ID!')
        });

        try {
            const existingInfo = await DodoClient.customers.retrieve(customer_id);
            const name = await input({
                message: "Enter customer name:",
                default: existingInfo.name
            });

            const phone = await input({
                message: "Enter customer phone:",
                default: existingInfo.phone_number?.toString()
            });

            const updated = await DodoClient.customers.update(customer_id, {
                name: name,
                phone_number: phone.trim() !== '' ? phone : null
            });

            console.table([updated], ['customer_id', 'name', 'email', 'phone_number']);
        } catch (e) {
            if (isDodoPaymentsAPIError(e) && e.error.code === "NOT_FOUND") {
                console.log("Incorrect customer ID!");
            } else {
                console.error(e);
            }
        }
    } else {
        usage.customers!.forEach(e => {
            console.log(`dodo customers ${e.command} - ${e.description}`)
        });
    }
} else if (category === 'discounts') {
    if (subCommand === 'list') {
        const page = await input({
            message: 'Enter page:',
            default: "1",
            validate: (e => e.trim() !== '')
        });
        const discounts = await DodoClient.discounts.list({ page_number: parseInt(page) - 1, page_size: 100 });
        const discountsTable = discounts.items.map(e => (
            {
                name: e.name,
                code: e.code,
                'discount id': e.discount_id,
                'created at': new Date(e.created_at).toLocaleString(),
                ...e.type === 'percentage' ? {
                    amount: `${(e.amount * 0.01).toFixed(2)}%`
                } : {
                    // I just added this in case of a breaking change in the future
                    amount: e.amount
                },
            }
        ));

        console.table(discountsTable);
        console.log(`To view a discount, go to https://app.dodopayments.com/sales/discounts/edit?id={discount_id}`)
    } else if (subCommand === 'create') {
        const name = await input({
            message: "Enter discount name:",
            validate: (e => e.trim() !== '')
        });

        const percentage = await input({
            message: "Enter discount percentage:",
            // Make sure user enters valid value
            validate: (e => {
                const parsed = parseFloat(e);
                if (!Number.isNaN(parsed) && parsed > 0 && parsed <= 100) {
                    return true;
                } else {
                    return false;
                }
            })
        });

        const code = await input({
            message: "Enter discount code (Optional):"
        });

        const cycles = await input({
            message: "Enter discount cycles (Optional):"
        });

        const newDiscount = await DodoClient.discounts.create({
            name,
            code: code.trim() !== '' ? code : null,
            amount: parseFloat(percentage) * 100,
            type: 'percentage',
            // If the subscription cycles is provided
            ...cycles.trim() !== '' && {
                subscription_cycles: parseInt(cycles)
            }
        });

        console.log('Discount created successfully!');
        console.table({
            name: newDiscount.name,
            code: newDiscount.code,
            'discount id': newDiscount.discount_id,
            ...cycles.trim() !== '' && {
                'subscription cycles': newDiscount.subscription_cycles
            }
        });
    } else if (subCommand === 'delete') {
        await DodoClient.discounts.delete(await input({
            message: "Enter discount ID to be deleted:",
            validate: (e => e.startsWith('dsc_'))
        }));

        console.log("Successfully deleted discount!");
    } else {
        usage.discounts!.forEach(e => {
            console.log(`dodo discounts ${e.command} - ${e.description}`)
        });
    }
} else if (category === 'licences') {
    if (subCommand === 'list') {
        const page = await input({
            message: 'Enter page:',
            default: "1",
            validate: (e => e.trim() !== '')
        });
        const licences = await DodoClient.licenseKeys.list({ page_number: parseInt(page) - 1, page_size: 100 });
        console.log(licences.items);
    } else {
        usage.licences!.forEach(e => {
            console.log(`dodo licences ${e.command} - ${e.description}`)
        });
    }
} else {
    RenderHelp();
}