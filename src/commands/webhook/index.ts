import DodoPayments from 'dodopayments';
import { usage } from '../../utils/usage-help';
import WebhookListener from './listen';

export async function handleWebhook(client: DodoPayments, apiKey: string, subCommand?: string) {
    switch (subCommand) {
        case 'listen':
            WebhookListener({ API_KEY: apiKey, dodoClient: client });
            break;
        case 'trigger':
            import('./trigger');
            break;
        default:
            usage.wh!.forEach(e => console.log(`dodo wh ${e.command} - ${e.description}`));
    }
}