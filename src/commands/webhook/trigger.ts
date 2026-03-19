import { input, select } from '@inquirer/prompts';
import type { baseArgs } from './types/baseArgs';
import {
  supportedEvents,
  type SupportedEvent,
} from './functions/supported-events';

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

export async function handleWebhookTrigger(): Promise<void> {
  const endpoint = await input({
    message: 'What is your endpoint?',
    validate: (value) => {
      if (!value.startsWith('http://') && !value.startsWith('https://')) {
        return 'Please enter a valid URL starting with http:// or https://';
      }

      return true;
    },
  });

  const businessId = await input({
    message: 'What is your Dodo Payments business ID? (Optional)',
    default: 'bus_test',
  });

  const productId = await input({
    message: 'What is your product ID? (Optional)',
    default: 'pdt_test',
  });

  const metadataInput = await input({
    message: 'What is your metadata? (JSON stringified, Optional)',
    default: '{}',
    validate: (value) => {
      try {
        parseMetadata(value);
        return true;
      } catch {
        return 'Please enter a valid JSON object.';
      }
    },
  });

  const email = await input({
    message: "What is the customer's email? (Optional)",
    default: 'john.doe@example.com',
  });

  const customerId = await input({
    message: "What is the customer's id? (Optional)",
    default: 'cus_test',
  });

  const metadata = parseMetadata(metadataInput);
  const eventChoices: Array<{ name: string; value: SupportedEvent | 'exit' }> = [
    ...supportedEvents.map((value) => ({
      name: value,
      value,
    })),
    {
      name: 'exit',
      value: 'exit',
    },
  ];

  while (true) {
    const event = await select<SupportedEvent | 'exit'>({
      message: 'Select an event to send:',
      choices: eventChoices,
      loop: false,
    });

    if (event === 'exit') {
      console.log('Exiting webhook trigger.');
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

      console.log('Webhook event sent successfully.');
      console.log(`Response status: ${response.status}`);
      console.log(`Response body: ${responseBody}`);
      console.log();
    } catch (error) {
      console.error('Webhook event failed:', error);
      console.log();
    }
  }
}
