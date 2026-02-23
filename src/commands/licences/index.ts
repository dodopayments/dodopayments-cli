import DodoPayments from 'dodopayments';
import { input } from '@inquirer/prompts';
import { usage } from '../../utils/usage-help';
import { isDodoPaymentsAPIError } from '../../utils/error';

export async function handleLicences(
  client: DodoPayments,
  subCommand?: string,
) {
  if (subCommand === 'list') {
    const page = await input({
      message: 'Enter page:',
      default: '1',
      validate: (e) => e.trim() !== '',
    });
    try {
      // WORK ON THIS
      const licences = await client.licenseKeys.list({
        page_number: parseInt(page) - 1,
        page_size: 100,
      });

      console.log(licences.items);
    } catch (e) {
      if (isDodoPaymentsAPIError(e)) {
        console.log('Failed to fetch licences:', e.error.message);
      } else {
        console.error(e);
      }
    }
  } else {
    usage.licences!.forEach((e) =>
      console.log(`dodo licences ${e.command} - ${e.description}`),
    );
  }
}
