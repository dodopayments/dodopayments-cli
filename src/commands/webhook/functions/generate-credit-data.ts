import { type baseArgs } from "../types/baseArgs"

const genBaseCreditLedgerData = ({
    business_id = 'bus_test',
    subscription_id = 'sub_test',
    metadata,
    customer_id
}: baseArgs) => {
    return {
        "business_id": business_id,
        "data": {
            "amount": "100",
            "balance_after": "100",
            "balance_before": "0",
            "brand_id": "brd_test",
            "business_id": business_id,
            "created_at": new Date().toISOString(),
            "credit_entitlement_id": "cent_test",
            "customer_id": customer_id,
            "description": "Credits granted from subscription",
            "grant_id": "cgr_test",
            "id": "cle_test",
            "is_credit": true,
            "metadata": metadata ?? {},
            "overage_after": "0",
            "overage_before": "0",
            "payload_type": "CreditLedgerEntry",
            "reference_id": subscription_id,
            "reference_type": "subscription",
            "transaction_type": "credit_added"
        },
        "timestamp": new Date().toISOString(),
    }
}

const genCreditAdded = (args: baseArgs) => {
    return {
        ...genBaseCreditLedgerData(args),
        "type": "credit.added"
    }
}

const genCreditDeducted = (args: baseArgs) => {
    const base = genBaseCreditLedgerData(args);
    return {
        ...base,
        "type": "credit.deducted",
        data: {
            ...base.data,
            amount: "25",
            balance_before: "100",
            balance_after: "75",
            description: "Credits consumed by usage",
            is_credit: false,
            transaction_type: "credit_deducted"
        }
    }
}

const genCreditExpired = (args: baseArgs) => {
    const base = genBaseCreditLedgerData(args);
    return {
        ...base,
        "type": "credit.expired",
        data: {
            ...base.data,
            amount: "75",
            balance_before: "75",
            balance_after: "0",
            description: "Unused credits expired at the end of the billing cycle",
            is_credit: false,
            transaction_type: "credit_expired"
        }
    }
}

const genCreditRolledOver = (args: baseArgs) => {
    const base = genBaseCreditLedgerData(args);
    return {
        ...base,
        "type": "credit.rolled_over",
        data: {
            ...base.data,
            amount: "40",
            balance_before: "100",
            balance_after: "140",
            description: "Unused credits rolled over into the new billing cycle",
            transaction_type: "credit_rolled_over"
        }
    }
}

const genCreditRolloverForfeited = (args: baseArgs) => {
    const base = genBaseCreditLedgerData(args);
    return {
        ...base,
        "type": "credit.rollover_forfeited",
        data: {
            ...base.data,
            amount: "10",
            balance_before: "10",
            balance_after: "0",
            description: "Rolled-over credits forfeited past the rollover limit",
            is_credit: false,
            transaction_type: "rollover_forfeited"
        }
    }
}

const genCreditOverageCharged = (args: baseArgs) => {
    const base = genBaseCreditLedgerData(args);
    return {
        ...base,
        "type": "credit.overage_charged",
        data: {
            ...base.data,
            amount: "15",
            balance_before: "0",
            balance_after: "0",
            description: "Overage charged beyond the included credit allowance",
            is_credit: false,
            overage_before: "0",
            overage_after: "15",
            transaction_type: "overage_charged"
        }
    }
}

const genCreditOverageReset = (args: baseArgs) => {
    const base = genBaseCreditLedgerData(args);
    return {
        ...base,
        "type": "credit.overage_reset",
        data: {
            ...base.data,
            amount: "15",
            balance_before: "0",
            balance_after: "0",
            description: "Overage counter reset at the start of the billing cycle",
            overage_before: "15",
            overage_after: "0",
            transaction_type: "overage_reset"
        }
    }
}

const genCreditManualAdjustment = (args: baseArgs) => {
    const base = genBaseCreditLedgerData(args);
    return {
        ...base,
        "type": "credit.manual_adjustment",
        data: {
            ...base.data,
            amount: "50",
            balance_before: "75",
            balance_after: "125",
            description: "Manual goodwill credit adjustment",
            grant_id: null,
            reference_id: null,
            reference_type: null,
            transaction_type: "manual_adjustment"
        }
    }
}

// credit.balance_low uses CreditBalanceLowPayload, not the ledger entry payload.
const genCreditBalanceLow = ({
    business_id = 'bus_test',
    subscription_id = 'sub_test',
    customer_id
}: baseArgs) => {
    return {
        "business_id": business_id,
        "data": {
            "available_balance": "15",
            "brand_id": "brd_test",
            "credit_entitlement_id": "cent_test",
            "credit_entitlement_name": "API Credits",
            "customer_id": customer_id,
            "payload_type": "CreditBalanceLow",
            "subscription_credits_amount": "100",
            "subscription_id": subscription_id,
            "threshold_amount": "20",
            "threshold_percent": 20
        },
        "timestamp": new Date().toISOString(),
        "type": "credit.balance_low"
    }
}

export {
    genCreditAdded,
    genCreditDeducted,
    genCreditExpired,
    genCreditRolledOver,
    genCreditRolloverForfeited,
    genCreditOverageCharged,
    genCreditOverageReset,
    genCreditManualAdjustment,
    genCreditBalanceLow
}
