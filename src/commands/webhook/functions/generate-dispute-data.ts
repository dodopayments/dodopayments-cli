import { type baseArgs } from "../types/baseArgs";

const genBaseDisputeData = ({
    business_id = 'bus_test'
}: baseArgs) => {
    return {
        "business_id": business_id,
        "data": {
            "amount": 400,
            "business_id": business_id,
            "created_at": new Date().toISOString(),
            "currency": "USD",
            "customer": {
                "customer_id": "cus_test",
                "email": "test@acme.com",
                "name": "Test user",
                "phone_number": "+447700900123"
            },
            "dispute_id": "dp_test",
            "dispute_stage": "dispute",
            "dispute_status": "dispute_opened",
            "payload_type": "Dispute",
            "payment_id": "pay_test",
            "reason": "testing dispute",
            "remarks": null
        },
        "timestamp": new Date().toISOString(),
    }
}

const genDisputeOpened = (args: baseArgs) => {
    const baseData = genBaseDisputeData(args);
    return {
        ...baseData,
        "type": "dispute.opened",
        data: {
            ...baseData.data,
            dispute_status: "dispute_opened"
        }
    }
}

const genDisputeExpired = (args: baseArgs) => {
    const baseData = genBaseDisputeData(args);
    return {
        ...baseData,
        "type": "dispute.expired",
        data: {
            ...baseData.data,
            dispute_status: "dispute_expired"
        }
    }
}

const genDisputeAccepted = (args: baseArgs) => {
    const baseData = genBaseDisputeData(args);
    return {
        ...baseData,
        "type": "dispute.accepted",
        data: {
            ...baseData.data,
            dispute_status: "dispute_accepted"
        }
    }
}

const genDisputeCancelled = (args: baseArgs) => {
    const baseData = genBaseDisputeData(args);
    return {
        ...baseData,
        "type": "dispute.cancelled",
        data: {
            ...baseData.data,
            dispute_status: "dispute_cancelled"
        }
    }
}

const genDisputeChallenged = (args: baseArgs) => {
    const baseData = genBaseDisputeData(args);
    return {
        ...baseData,
        "type": "dispute.challenged",
        data: {
            ...baseData.data,
            dispute_status: "dispute_challenged"
        }
    }
}

const genDisputeWon = (args: baseArgs) => {
    const baseData = genBaseDisputeData(args);
    return {
        ...baseData,
        "type": "dispute.won",
        data: {
            ...baseData.data,
            dispute_status: "dispute_won"
        }
    }
}

const genDisputeLost = (args: baseArgs) => {
    const baseData = genBaseDisputeData(args);
    return {
        ...baseData,
        "type": "dispute.lost",
        data: {
            ...baseData.data,
            dispute_status: "dispute_lost"
        }
    }
}

export {
    genDisputeOpened,
    genDisputeExpired,
    genDisputeAccepted,
    genDisputeCancelled,
    genDisputeChallenged,
    genDisputeWon,
    genDisputeLost
};