import DodoPayments from 'dodopayments';
import { input, select } from '@inquirer/prompts';
import { usage } from '../../utils/usage-help';
import { isDodoPaymentsAPIError } from '../../utils/error';

export async function handleCheckout(
  client: DodoPayments,
  subCommand?: string,
) {
  if (subCommand === 'new') {
    let config: DodoPayments.CheckoutSessions.CheckoutSessionCreateParams = {
      product_cart: [],
    };

    const product = await input({
      message: 'Enter product:',
      validate: (e) => e.startsWith('pdt_'),
    });

    config.product_cart = [{ product_id: product, quantity: 1 }];

    const useAdvanced = await select({
      message: 'Use advanced options?',
      choices: [
        { name: 'Yes', value: true },
        { name: 'No', value: false },
      ],
      default: false,
    });

    if (useAdvanced) {
      config.minimal_address = await select({
        message: 'Enable minimal address:',
        choices: [
          { name: 'Yes', value: true },
          { name: 'No', value: false },
        ],
        default: false,
      });

      const return_url = await input({
        message: 'Enter return URL (Optional):',
      });

      if (return_url.trim() !== '') {
        config.return_url = return_url;
      }

      config.force_3ds = await select({
        message: 'Force 3DS?',
        choices: [
          { name: 'Yes', value: true },
          { name: 'No', value: false },
        ],
      });

      const disc_code = await input({
        message: 'Enter discount code (Optional):',
      });

      if (disc_code.trim() !== '') {
        config.discount_code = disc_code;
      }

      const metadata = await input({
        message: 'Enter metadata (Optional, JSON stringified):',
      });

      if (metadata.trim() !== '') {
        config.metadata = JSON.parse(metadata);
      }
    }
    try {
      const session = await client.checkoutSessions.create(config);
      console.log('Checkout Session URL:', session.checkout_url);
    } catch (e) {
      // This is the only possible error here. I have used isDodoPaymentsAPIError() to infer types support.
      if (isDodoPaymentsAPIError(e)) {
        console.log(`Error: ${e.error.message}`);
      } else {
        console.error(e);
      }
    }
  } else {
    usage.checkout!.forEach((e) =>
      console.log(`dodo checkout ${e.command} - ${e.description}`),
    );
  }
}
