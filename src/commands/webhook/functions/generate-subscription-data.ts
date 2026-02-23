import { type baseArgs } from '../types/baseArgs';

// For subscription data
const genBaseSubscriptionData = ({
    business_id = 'bus_test',
    product_id = 'pdt_test',
    subscription_id = 'sub_test',
    qty = 1,
    metadata
}: baseArgs) => {
    return {
        "business_id": business_id,
        "data": {
            "addons": [],
            "billing": {
                "city": "New York",
                "country": "US",
                "state": "New York",
                "street": "11th Main",
                "zipcode": "12345"
            },
            "cancel_at_next_billing_date": false,
            "cancelled_at": null,
            "created_at": new Date().toISOString(),
            "currency": "USD",
            "customer": {
                "customer_id": "cus_123456",
                "email": "test@acme.com",
                "name": "Test user",
                "phone_number": null
            },
            "discount_cycles_remaining": null,
            "discount_id": null,
            "expires_at": null,
            "metadata": metadata,
            "meters": [],
            "next_billing_date": new Date(new Date().setMonth(new Date().getMonth() + 1)),
            "on_demand": false,
            "payload_type": "Subscription",
            "payment_frequency_count": 1,
            "payment_frequency_interval": "Month",
            "previous_billing_date": new Date().toISOString(),
            "product_id": product_id,
            "quantity": qty,
            "recurring_pre_tax_amount": 400,
            "status": "active",
            "subscription_id": subscription_id,
            "subscription_period_count": 10,
            "subscription_period_interval": "Year",
            "tax_inclusive": false,
            "trial_period_days": 0
        },
        "timestamp": new Date().toISOString()
    }
}

const genSubscriptionActive = (args: baseArgs) => {
    return {
        ...genBaseSubscriptionData(args),
        "type": "subscription.active"
    };
}

const genSubscriptionUpdated = (args: baseArgs) => {
    const base = genBaseSubscriptionData(args);
    return {
        ...base,
        "type": "subscription.updated",
        data: {
            ...base.data,
            status: "active"
        }
    };
}

const genSubscriptionOnHold = (args: baseArgs) => {
    const base = genBaseSubscriptionData(args);
    return {
        ...base,
        "type": "subscription.on_hold",
        data: {
            ...base.data,
            status: "on_hold"
        }
    };
}

const genSubscriptionRenewed = (args: baseArgs) => {
    return {
        ...genBaseSubscriptionData(args),
        "type": "subscription.renewed"
    };
}

const genSubscriptionPlanChanged = (args: baseArgs) => {
    const base = genBaseSubscriptionData(args);
    return {
        ...base,
        "type": "subscription.plan_changed",
        data: {
            ...base.data,
            status: "active"
        }
    };
}

const genSubscriptionCancelled = (args: baseArgs) => {
    const base = genBaseSubscriptionData(args);
    return {
        ...base,
        "type": "subscription.cancelled",
        data: {
            ...base.data,
            cancel_at_next_billing_date: true,
            status: 'cancelled'
        }
    };
}

const genSubscriptionFailed = (args: baseArgs) => {
    const base = genBaseSubscriptionData(args);
    return {
        ...base,
        "type": "subscription.failed",
        data: {
            ...base.data,
            status: "failed"
        }
    };
}

const genSubscriptionExpired = (args: baseArgs) => {
    const base = genBaseSubscriptionData(args);
    return {
        ...base,
        "type": "subscription.expired",
        data: {
            ...base.data,
            status: "expired"
        }
    };
}

export {
    genSubscriptionActive,
    genSubscriptionUpdated,
    genSubscriptionOnHold,
    genSubscriptionRenewed,
    genSubscriptionPlanChanged,
    genSubscriptionCancelled,
    genSubscriptionFailed,
    genSubscriptionExpired
};