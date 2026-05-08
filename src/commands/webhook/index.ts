import type DodoPayments from 'dodopayments';
import { usage } from '../../utils/usage-help';
import WebhookListener from './listen';
import { handleWebhookTrigger } from './trigger';
import type { CommandContext } from '../../ui/ink/CommandContext';

type AuthenticatedWebhookContext = {
  apiKey: string;
  client: DodoPayments;
};

export async function handleWebhook(
  subCommand: string | undefined,
  ctx: CommandContext,
  context?: AuthenticatedWebhookContext,
) {
  switch (subCommand) {
    case 'listen':
      if (!context) {
        ctx.addBlock({ type: 'error', message: 'Please run `dodo login` first.' });
        process.exit(1);
      }

      await WebhookListener({
        API_KEY: context.apiKey,
        dodoClient: context.client,
        ctx,
      });
      break;
    case 'trigger':
      await handleWebhookTrigger(ctx);
      break;
    default:
      ctx.addBlock({ type: 'help' });
  }
}
