import type DodoPayments from 'dodopayments';
import { usage } from '../../utils/usage-help';
import WebhookListener from './listen';
import { handleWebhookTrigger } from './trigger';

type AuthenticatedWebhookContext = {
  apiKey: string;
  client: DodoPayments;
};

export async function handleWebhook(
  subCommand?: string,
  context?: AuthenticatedWebhookContext,
) {
  switch (subCommand) {
    case 'listen':
      if (!context) {
        console.error('Please run `dodo login` first.');
        process.exit(1);
      }

      await WebhookListener({
        API_KEY: context.apiKey,
        dodoClient: context.client,
      });
      break;
    case 'trigger':
      await handleWebhookTrigger();
      break;
    default:
      usage.wh!.forEach((e) =>
        console.log(`dodo wh ${e.command} - ${e.description}`),
      );
  }
}
