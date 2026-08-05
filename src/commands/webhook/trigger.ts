import type { baseArgs } from './types/baseArgs';
import {
  supportedEvents,
  type SupportedEvent,
} from './functions/supported-events';
import type { CommandContext } from '../../tui/CommandContext';
import { HTTP_SCHEME_ERROR, isHttpUrl } from '../../utils/url';

import {
  genSubscriptionActive,
  genSubscriptionUpdated,
  genSubscriptionOnHold,
  genSubscriptionRenewed,
  genSubscriptionPlanChanged,
  genSubscriptionCancelled,
  genSubscriptionFailed,
  genSubscriptionExpired,
  genSubscriptionPaused,
  genSubscriptionUpdatePaymentMethod,
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

import {
  genPayoutCreated,
  genPayoutInProgress,
  genPayoutOnHold,
  genPayoutFailed,
  genPayoutSuccess,
} from './functions/generate-payout-data';

import {
  genCreditAdded,
  genCreditDeducted,
  genCreditExpired,
  genCreditRolledOver,
  genCreditRolloverForfeited,
  genCreditOverageCharged,
  genCreditOverageReset,
  genCreditManualAdjustment,
  genCreditBalanceLow,
} from './functions/generate-credit-data';

import {
  genAbandonedCheckoutDetected,
  genAbandonedCheckoutRecovered,
  genDunningStarted,
  genDunningRecovered,
} from './functions/generate-recovery-data';

import {
  genEntitlementGrantCreated,
  genEntitlementGrantDelivered,
  genEntitlementGrantFailed,
  genEntitlementGrantRevoked,
} from './functions/generate-entitlement-grant-data';

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
  'subscription.paused': genSubscriptionPaused,
  'subscription.update_payment_method': genSubscriptionUpdatePaymentMethod,
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
  'payout.created': genPayoutCreated,
  'payout.in_progress': genPayoutInProgress,
  'payout.on_hold': genPayoutOnHold,
  'payout.failed': genPayoutFailed,
  'payout.success': genPayoutSuccess,
  'credit.added': genCreditAdded,
  'credit.deducted': genCreditDeducted,
  'credit.expired': genCreditExpired,
  'credit.rolled_over': genCreditRolledOver,
  'credit.rollover_forfeited': genCreditRolloverForfeited,
  'credit.overage_charged': genCreditOverageCharged,
  'credit.overage_reset': genCreditOverageReset,
  'credit.manual_adjustment': genCreditManualAdjustment,
  'credit.balance_low': genCreditBalanceLow,
  'abandoned_checkout.detected': genAbandonedCheckoutDetected,
  'abandoned_checkout.recovered': genAbandonedCheckoutRecovered,
  'dunning.started': genDunningStarted,
  'dunning.recovered': genDunningRecovered,
  'entitlement_grant.created': genEntitlementGrantCreated,
  'entitlement_grant.delivered': genEntitlementGrantDelivered,
  'entitlement_grant.failed': genEntitlementGrantFailed,
  'entitlement_grant.revoked': genEntitlementGrantRevoked,
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

const USAGE = [
  'Usage: dodo wh trigger <event> <url>',
  '',
  'Supported events:',
  ...supportedEvents.map((e) => `  ${e}`),
].join('\n');

async function sendEvent(
  ctx: CommandContext,
  endpoint: string,
  event: SupportedEvent,
  args: baseArgs,
): Promise<void> {
  const data = eventGenerators[event](args);

  try {
    const response = await fetch(endpoint, {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(data),
    });

    const responseBody = await response.text();

    ctx.addBlock({ type: 'success', message: 'Webhook event sent.' });
    ctx.addBlock({
      type: 'detail',
      data: { Status: response.status, Response: responseBody },
    });
  } catch (error: any) {
    ctx.addBlock({ type: 'error', message: `Webhook event failed. ${error.message}` });
  }
}

async function triggerOnce(
  ctx: CommandContext,
  eventArg: string | undefined,
  urlArg: string | undefined,
): Promise<void> {
  if (!eventArg || !urlArg) {
    ctx.addBlock({ type: 'error', message: 'Event and URL are required.' });
    ctx.addBlock({ type: 'info', message: USAGE });
    return;
  }

  if (!(supportedEvents as readonly string[]).includes(eventArg)) {
    ctx.addBlock({ type: 'error', message: `Unsupported event '${eventArg}'.` });
    ctx.addBlock({ type: 'info', message: USAGE });
    return;
  }
  const event = eventArg as SupportedEvent;

  if (!isHttpUrl(urlArg)) {
    ctx.addBlock({ type: 'error', message: HTTP_SCHEME_ERROR });
    ctx.addBlock({ type: 'info', message: USAGE });
    return;
  }

  await sendEvent(ctx, urlArg, event, {
    business_id: 'bus_test',
    product_id: 'pdt_test',
    metadata: {},
    email: 'john.doe@example.com',
    customer_id: 'cus_test',
  });
}

export async function handleWebhookTrigger(
  ctx: CommandContext,
  extraArgs: readonly string[] = [],
): Promise<void> {
  const [eventArg, urlArg] = extraArgs;
  if (eventArg || urlArg) {
    await triggerOnce(ctx, eventArg, urlArg);
    return;
  }

  let endpoint = '';
  while (!endpoint) {
    let value: string;
    try {
      value = await ctx.promptInput('Endpoint URL');
    } catch {
      ctx.addBlock({ type: 'error', message: 'Event and URL are required.' });
      ctx.addBlock({ type: 'info', message: USAGE });
      return;
    }
    if (isHttpUrl(value)) {
      endpoint = value;
    } else {
      ctx.addBlock({ type: 'error', message: HTTP_SCHEME_ERROR });
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

    await sendEvent(ctx, endpoint, event, {
      business_id: businessId,
      product_id: productId,
      metadata,
      email,
      customer_id: customerId,
    });
  }
}
