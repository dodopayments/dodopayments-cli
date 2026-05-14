import type DodoPayments from 'dodopayments';
import WebhookListener from './listen';
import { handleWebhookTrigger } from './trigger';
import type { CommandContext } from '../../tui/CommandContext';
import { unknownSubcommand } from '../../utils/usage-help';

type AuthenticatedWebhookContext = {
  apiKey: string;
  client: DodoPayments;
};

export async function handleWebhook(
  subCommand: string | undefined,
  ctx: CommandContext,
  context?: AuthenticatedWebhookContext,
  extraArgs: string[] = [],
) {
  switch (subCommand) {
    case 'listen':
      if (!context) {
        ctx.addBlock({ type: 'error', message: 'Sign in first. Run /login to get started.' });
        process.exit(1);
      }

      await WebhookListener({
        API_KEY: context.apiKey,
        dodoClient: context.client,
        ctx,
        endpoint: extraArgs[0],
      });
      break;
    case 'trigger':
      await handleWebhookTrigger(ctx, extraArgs);
      break;
    default:
      unknownSubcommand(ctx, 'wh', subCommand);
  }
}
