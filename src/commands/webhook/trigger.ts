import { input, password as passwordPrompt, select } from '@inquirer/prompts';
import { supportedEvents } from './functions/supported-events';

// Subscription Events
import {
    genSubscriptionActive,
    genSubscriptionUpdated,
    genSubscriptionOnHold,
    genSubscriptionRenewed,
    genSubscriptionPlanChanged,
    genSubscriptionCancelled,
    genSubscriptionFailed,
    genSubscriptionExpired
} from './functions/generate-subscription-data';

// Payment events
import {
    genPaymentSuccess,
    genPaymentFailed,
    genPaymentProcessing,
    genPaymentCancelled
} from './functions/generate-payment-data';

// Refund events
import {
    genRefundSuccess,
    genRefundFailed
} from './functions/generate-refund-data';

// Dispute events
import {
    genDisputeOpened,
    genDisputeExpired,
    genDisputeAccepted,
    genDisputeCancelled,
    genDisputeChallenged,
    genDisputeWon,
    genDisputeLost
} from './functions/generate-dispute-data';

// Licence events
import {
    genLicenceCreated
} from './functions/generate-licence-data';

const endpoint = await input({
    message: 'What is your endpoint?',
    validate: (input) => {
        if (!input.startsWith('http')) {
            return 'Please enter a valid URL starting with http or https';
        }
        return true;
    }
});

// const webhook_secret = await passwordPrompt({
//     message: 'What is your Dodo Payments webhook secret? (Optional)',
//     mask: true
// });

const business_id = await input({
    message: 'What is your Dodo Payments business ID? (Optional)',
    default: 'bus_test'
});

const product_id = await input({
    message: 'What is your product ID? (Optional)',
    default: 'pdt_test'
});

const metadata = await input({
    message: 'What is your metadata? (JSON stringified, Optional)',
    default: '{}'
});

while (true) {
    const event = await select({
        message: 'Select an event to send:',
        choices: [...supportedEvents, 'exit'].map(e => ({ value: e, name: e })),
        loop: false,
    });

    if (event === 'exit') {
        console.log('Exiting...');
        process.exit();
    }

    let data;
    const metadataParsed: { [key: string]: string } = JSON.parse(metadata);

    // Parameters based on subscription event
    if (event.startsWith('subscription.')) {
        const params = {
            business_id,
            product_id,
            metadata: metadataParsed
        };

        if (event === 'subscription.active') {
            data = genSubscriptionActive(params);
        } else if (event === 'subscription.updated') {
            data = genSubscriptionUpdated(params);
        } else if (event === 'subscription.on_hold') {
            data = genSubscriptionOnHold(params);
        } else if (event === 'subscription.renewed') {
            data = genSubscriptionRenewed(params);
        } else if (event === 'subscription.plan_changed') {
            data = genSubscriptionPlanChanged(params);
        } else if (event === 'subscription.cancelled') {
            data = genSubscriptionCancelled(params);
        } else if (event === 'subscription.failed') {
            data = genSubscriptionFailed(params);
        } else if (event === 'subscription.expired') {
            data = genSubscriptionExpired(params);
        }
    }

    // Parameters based on payment event
    if (event.startsWith('payment.')) {
        const params = {
            business_id,
            product_id,
            metadata: metadataParsed
        };

        if (event === 'payment.success') {
            data = genPaymentSuccess(params);
        } else if (event === 'payment.failed') {
            data = genPaymentFailed(params);
        } else if (event === 'payment.processing') {
            data = genPaymentProcessing(params);
        } else if (event === 'payment.cancelled') {
            data = genPaymentCancelled(params);
        }
    }

    // Parameters based on refund event
    if (event.startsWith('refund.')) {
        const params = {
            business_id
        };

        if (event === 'refund.success') {
            data = genRefundSuccess(params);
        } else if (event === 'refund.failed') {
            data = genRefundFailed(params);
        }
    }

    // Parameters based on dispute event
    if (event.startsWith('dispute.')) {
        const params = {
            business_id
        };

        if (event === 'dispute.opened') {
            data = genDisputeOpened(params);
        } else if (event === 'dispute.expired') {
            data = genDisputeExpired(params);
        } else if (event === 'dispute.accepted') {
            data = genDisputeAccepted(params);
        } else if (event === 'dispute.cancelled') {
            data = genDisputeCancelled(params);
        } else if (event === 'dispute.challenged') {
            data = genDisputeChallenged(params);
        } else if (event === 'dispute.won') {
            data = genDisputeWon(params);
        } else if (event === 'dispute.lost') {
            data = genDisputeLost(params);
        }
    }

    // Parameters based on licence event
    if (event.startsWith('licence.')) {
        const params = {
            business_id,
            product_id
        };

        if (event === 'licence.created') {
            data = genLicenceCreated(params);
        }
    }

    // const timestamp = new Date();
    // const webhookId = crypto.randomUUID();
    // const webhookInstance = new Webhook(webhook_secret);
    // const payloadString = JSON.stringify(data);
    // const signature = webhookInstance.sign(webhookId, timestamp, payloadString);

    await fetch(endpoint, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
            // "webhook-id": webhookId,
            // "webhook-timestamp": Math.floor(timestamp.getTime() / 1000).toString(),
            // "webhook-signature": signature,
        },
        body: JSON.stringify(data)
    }).then(async res => {
        console.log(`✅ Webhook Event Sent!\nReceived Response Code: ${res.status}\nReceived Response Body: ${await res.text()}`);
        console.log();
    }).catch(error => {
        console.error('❌ Webhook Event Failed:', error);
        console.log();
    });
}