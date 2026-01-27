import { type baseArgs } from "../types/baseArgs"

const genBaseRefundData = ({
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
                "phone_number": null
            },
            "is_partial": false,
            "payload_type": "Refund",
            "payment_id": "pay_test",
            "reason": "Testing success refund",
            "refund_id": "ref_test",
            "status": "succeeded"
        },
        "timestamp": new Date().toISOString(),
    }
}

const genRefundSuccess = (args: baseArgs) => {
    return {
        ...genBaseRefundData(args),
        "type": "refund.succeeded"
    }
}

const genRefundFailed = (args: baseArgs) => {
    return {
        ...genBaseRefundData(args),
        "type": "refund.failed"
    }
}

export { genRefundSuccess, genRefundFailed }