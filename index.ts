import os from 'node:os';
import path from 'node:path';
import DodoPayments from 'dodopayments';
import { input, select } from '@inquirer/prompts';
import open from 'open';
import { CurrencyToSymbolMap } from './utils/currency-to-symbol-map';

type DodoPaymentsAPIError = {
    error: {
        code: string;
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

const link = (text: string, url: string) =>
    `\u001b]8;;${url}\u001b\\${text}\u001b]8;;\u001b\\`;

const usage = {
    products: [
        { command: 'list', description: 'List your products' },
        { command: 'create', description: 'Create a new product' }
    ]
}


const args = Bun.argv;
const category = args[2];
const subCommand = args[3];
const homedir = os.homedir();

// Added this to the top so that it can bypass all further auth that happens for the login route
if (category === 'login') {
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
        existingConfig = Object.create(await Bun.file(path.join(homedir, '.dodopayments', 'api-key')).json());
    } catch {
        existingConfig = {};
    }

    existingConfig[MODE] = API_KEY;
    await Bun.file(path.join(homedir, '.dodopayments', 'api-key')).write(JSON.stringify(existingConfig));

    // Mode will always be either test_mode or live_mode
    existingConfig[MODE] = API_KEY;
    console.log("Setup complete successfully!");
    process.exit(0);
}





// Normal functions which require the API key to be present start from here
// Authentication part
const ExistingAPIKeyConfigFile = Bun.file(path.join(homedir, '.dodopayments', 'api-key'));

// Read the API key config
if (!(await ExistingAPIKeyConfigFile.exists())) {
    console.log('Please login using `dodo login` command first!');
    process.exit(0);
}

// Parse the API key config
let existingAPIKeyConfigParsed;
try {
    existingAPIKeyConfigParsed = await ExistingAPIKeyConfigFile.json();
} catch {
    // Delete API config if something fails with parsing
    await Bun.file(path.join(homedir, '.dodopayments', 'api-key')).delete();
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
        console.table((await DodoClient.products.list({ page_number: parseInt(page) - 1, page_size: 100 })).items, ['name', 'product_id', 'updated_at', 'price']);
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
                description: info.description,
                created_at: info.created_at,
                ...info.is_recurring ? {
                    price: `${CurrencyToSymbolMap[info.price.currency] || (info.price.currency + ' ')}${info.price.price * 0.01}/${info.price.payment_frequency_interval}`,
                } : {
                    price: `${CurrencyToSymbolMap[info.price.currency] || (info.price.currency + ' ')}${info.price.price * 0.01} (One Time)`,
                },
                tax_category: info.tax_category,
                edit_url: link('CTRL + Click to open', `https://app.dodopayments.com/products/edit?id=${info.product_id}`)
            });
        } catch (e) {
            if (isDodoPaymentsAPIError(e) && e.error.code === "NOT_FOUND") {
                console.log("Incorrect product ID!");
            } else {
                console.error(e);
            }
        }
    } else {
        usage.products.forEach(e => {
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
                payment_id: payment.payment_id,
                created_at: new Date(payment.created_at).toLocaleString(),
                subscription_id: payment.subscription_id,
                total_amount: payment.total_amount,
                currency: payment.currency,
                status: payment.status,
                more_info: link('CTRL + Click to open', `https://app.dodopayments.com/transactions/payments/${payment.payment_id}`)
            };
        });
        console.table(paymentsTable);
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
    }
} else {
    // List all available methods
    // todo: Add more comments to make it clear what's being done
    Object.keys(usage).forEach(e => {
        console.log(`Category: ${e}`);
        (usage as any)[e].forEach((e: { command: string, description: string }) => {
            console.log(`dodo products ${e.command} - ${e.description}`)
        });
        // Blank space as a separator
        console.log();
    });
}