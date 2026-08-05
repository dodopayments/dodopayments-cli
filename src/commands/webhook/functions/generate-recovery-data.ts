import { type baseArgs } from "../types/baseArgs"

const genBaseAbandonedCheckoutData = ({
    business_id = 'bus_test',
    customer_id
}: baseArgs) => {
    return {
        "business_id": business_id,
        "data": {
            "abandoned_at": new Date().toISOString(),
            "abandonment_reason": "checkout_incomplete",
            "brand_id": "brd_test",
            "customer_id": customer_id,
            "payload_type": "AbandonedCheckout",
            "payment_id": "pay_test",
            "recovered_payment_id": null,
            "status": "abandoned"
        },
        "timestamp": new Date().toISOString(),
    }
}

const genAbandonedCheckoutDetected = (args: baseArgs) => {
    return {
        ...genBaseAbandonedCheckoutData(args),
        "type": "abandoned_checkout.detected"
    }
}

const genAbandonedCheckoutRecovered = (args: baseArgs) => {
    const base = genBaseAbandonedCheckoutData(args);
    return {
        ...base,
        "type": "abandoned_checkout.recovered",
        data: {
            ...base.data,
            status: "recovered",
            recovered_payment_id: "pay_test_recovered"
        }
    }
}

const genBaseDunningData = ({
    business_id = 'bus_test',
    subscription_id = 'sub_test',
    customer_id
}: baseArgs) => {
    return {
        "business_id": business_id,
        "data": {
            "brand_id": "brd_test",
            "created_at": new Date().toISOString(),
            "customer_id": customer_id,
            "payload_type": "DunningAttempt",
            "payment_id": null,
            "status": "recovering",
            "subscription_id": subscription_id,
            "trigger_state": "on_hold"
        },
        "timestamp": new Date().toISOString(),
    }
}

const genDunningStarted = (args: baseArgs) => {
    return {
        ...genBaseDunningData(args),
        "type": "dunning.started"
    }
}

const genDunningRecovered = (args: baseArgs) => {
    const base = genBaseDunningData(args);
    return {
        ...base,
        "type": "dunning.recovered",
        data: {
            ...base.data,
            status: "recovered",
            payment_id: "pay_test"
        }
    }
}

export {
    genAbandonedCheckoutDetected,
    genAbandonedCheckoutRecovered,
    genDunningStarted,
    genDunningRecovered
}
