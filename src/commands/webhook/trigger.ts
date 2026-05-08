import type { baseArgs } from './types/baseArgs';
import {
  supportedEvents,
  type SupportedEvent,
} from './functions/supported-events';
import type { CommandContext } from '../../ui/ink/CommandContext';

import {
  genSubscriptionActive,
  genSubscriptionUpdated,
  genSubscriptionOnHold,
  genSubscriptionRenewed,
  genSubscriptionPlanChanged,
  genSubscriptionCancelled,
  genSubscriptionFailed,
  genSubscriptionExpired,
} from './functions/generate-subscription-data';

import {
  genPaymentSuccess,
  genPaymentFailed,
  genPaymentProcessing,
  genPaymentCancelled,
} from './functions/generate-payment-data';

import { genRefundSuccess, genRefundFailed } from './functions/generate-refund-data';

import {
  genDisputeOpened,
  genDisputeExpired,
  genDisputeAccepted,
  genDisputeCancelled,
  genDisputeChallenged,
  genDisputeWon,
  genDisputeLost,
} from './functions/generate-dispute-data';

import { genLicenceCreated } from './functions/generate-licence-data';

type PayloadGenerator = (args: baseArgs) => unknown;

const eventGenerators: Record<SupportedEvent, PayloadGenerator> = {
  'subscription.active': genSubscriptionActive,
  'subscription.updated': genSubscriptionUpdated,
  'subscription.on_hold': genSubscriptionOnHold,
  'subscription.renewed': genSubscriptionRenewed,
  'subscription.plan_changed': genSubscriptionPlanChanged,
  'subscription.cancelled': genSubscriptionCancelled,
  'subscription.failed': genSubscriptionFailed,
  'subscription.expired': genSubscriptionExpired,
  'payment.success': genPaymentSuccess,
  'payment.failed': genPaymentFailed,
  'payment.processing': genPaymentProcessing,
  'payment.cancelled': genPaymentCancelled,
  'licence.created': genLicenceCreated,
  'refund.success': genRefundSuccess,
  'refund.failed': genRefundFailed,
  'dispute.opened': genDisputeOpened,
  'dispute.expired': genDisputeExpired,
  'dispute.accepted': genDisputeAccepted,
  'dispute.cancelled': genDisputeCancelled,
  'dispute.challenged': genDisputeChallenged,
  'dispute.won': genDisputeWon,
  'dispute.lost': genDisputeLost,
};

function parseMetadata(metadataInput: string): Record<string, unknown> {
  try {
    const parsed = JSON.parse(metadataInput);

    if (typeof parsed !== 'object' || parsed === null || Array.isArray(parsed)) {
      throw new Error('INVALID_METADATA');
    }

    return parsed as Record<string, unknown>;
  } catch {
    throw new Error('INVALID_METADATA');
  }
}

export async function handleWebhookTrigger(ctx: CommandContext): Promise<void> {
  let endpoint = '';
  while (!endpoint) {
    const value = await ctx.promptInput('Endpoint URL');
    if (value.startsWith('http://') || value.startsWith('https://')) {
      endpoint = value;
    } else {
      ctx.addBlock({ type: 'error', message: 'URL must start with http:// or https://' });
    }
  }

  const businessIdRaw = await ctx.promptInput('Business ID (optional)');
  const businessId = businessIdRaw.trim() || 'bus_test';

  const productIdRaw = await ctx.promptInput('Product ID (optional)');
  const productId = productIdRaw.trim() || 'pdt_test';

  let metadata = {};
  while (true) {
    const metadataInputRaw = await ctx.promptInput('Metadata (optional, JSON)');
    const metadataInput = metadataInputRaw.trim() || '{}';
    try {
      metadata = parseMetadata(metadataInput);
      break;
    } catch {
      ctx.addBlock({ type: 'error', message: 'Must be a valid JSON object.' });
    }
  }

  const emailRaw = await ctx.promptInput('Customer email (optional)');
  const email = emailRaw.trim() || 'john.doe@example.com';

  const customerIdRaw = await ctx.promptInput('Customer ID (optional)');
  const customerId = customerIdRaw.trim() || 'cus_test';

  const eventChoices: Array<{ label: string; value: SupportedEvent | 'exit' }> = [
    ...supportedEvents.map((value) => ({
      label: value,
      value,
    })),
    {
      label: 'exit',
      value: 'exit',
    },
  ];

  while (true) {
    const event = await ctx.promptSelect('Event to send', eventChoices) as SupportedEvent | 'exit';

    if (event === 'exit') {
      ctx.addBlock({ type: 'success', message: 'Webhook trigger closed.' });
      return;
    }

    const generator = eventGenerators[event];
    const data = generator({
      business_id: businessId,
      product_id: productId,
      metadata,
      email,
      customer_id: customerId,
    });

    try {
      const response = await fetch(endpoint, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(data),
      });

      const responseBody = await response.text();

      ctx.addBlock({ type: 'success', message: 'Webhook event sent.' });
      ctx.addBlock({
        type: 'detail',
        data: {
          'Status': response.status,
          'Response': responseBody
        }
      });
    } catch (error: any) {
      ctx.addBlock({ type: 'error', message: `Webhook event failed. ${error.message}` });
    }
  }
}
