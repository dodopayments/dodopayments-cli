import { type baseArgs } from "../types/baseArgs"

const genBasePayoutData = ({
    business_id = 'bus_test'
}: baseArgs) => {
    return {
        "business_id": business_id,
        "data": {
            "amount": 100000,
            "business_id": business_id,
            "chargebacks": 0,
            "created_at": new Date().toISOString(),
            "currency": "USD",
            "fee": 2500,
            "name": "Test payout",
            "payload_type": "Payout",
            "payment_method": "bank_transfer",
            "payout_document_url": null,
            "payout_id": "pyt_test",
            "refunds": 0,
            "remarks": null,
            "status": "not_initiated",
            "tax": 0,
            "updated_at": new Date().toISOString()
        },
        "timestamp": new Date().toISOString(),
    }
}

// `payout.created` is emitted while the payout still reports `not_initiated`.
const genPayoutCreated = (args: baseArgs) => {
    return {
        ...genBasePayoutData(args),
        "type": "payout.created"
    }
}

const genPayoutInProgress = (args: baseArgs) => {
    const base = genBasePayoutData(args);
    return {
        ...base,
        "type": "payout.in_progress",
        data: {
            ...base.data,
            status: "in_progress"
        }
    }
}

const genPayoutOnHold = (args: baseArgs) => {
    const base = genBasePayoutData(args);
    return {
        ...base,
        "type": "payout.on_hold",
        data: {
            ...base.data,
            status: "on_hold",
            remarks: "Payout placed under review"
        }
    }
}

const genPayoutFailed = (args: baseArgs) => {
    const base = genBasePayoutData(args);
    return {
        ...base,
        "type": "payout.failed",
        data: {
            ...base.data,
            status: "failed",
            remarks: "Bank transfer rejected by the beneficiary bank"
        }
    }
}

const genPayoutSuccess = (args: baseArgs) => {
    const base = genBasePayoutData(args);
    return {
        ...base,
        "type": "payout.success",
        data: {
            ...base.data,
            status: "success",
            payout_document_url: "https://example.com/payouts/pyt_test.pdf"
        }
    }
}

export {
    genPayoutCreated,
    genPayoutInProgress,
    genPayoutOnHold,
    genPayoutFailed,
    genPayoutSuccess
}
