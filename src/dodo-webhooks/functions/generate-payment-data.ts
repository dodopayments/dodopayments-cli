import { type baseArgs } from "../types/baseArgs";

// For payment events
const genBasePaymentData = ({
    business_id = 'bus_test',
    product_id = 'pdt_test',
    qty = 1
}: baseArgs) => {
    return {
        "business_id": business_id,
        "data": {
            "billing": {
                "city": "New York",
                "country": "US",
                "state": "New York",
                "street": "New York, New York",
                "zipcode": "12345"
            },
            "brand_id": business_id,
            "business_id": business_id,
            "card_issuing_country": "GB",
            "card_last_four": "4242",
            "card_network": "VISA",
            "card_type": "CREDIT",
            "checkout_session_id": "cks_123",
            "created_at": new Date().toISOString(),
            "currency": "USD",
            "customer": {
                "customer_id": "cus_test",
                "email": "test@acme.com",
                "metadata": {},
                "name": "Test user",
                "phone_number": "+15555550100"
            },
            "digital_products_delivered": false,
            "discount_id": null,
            "disputes": [],
            "error_code": null,
            "error_message": null,
            "invoice_id": "inv_test",
            "metadata": {},
            "payload_type": "Payment",
            "payment_id": "pay_test",
            "payment_link": "https://test.checkout.dodopayments.com/cbq",
            "payment_method": "card",
            "payment_method_type": "card",
            "product_cart": [
                {
                    "product_id": product_id,
                    "quantity": qty
                }
            ],
            "refunds": [],
            "settlement_amount": 400,
            "settlement_currency": "USD",
            "settlement_tax": null,
            "status": "succeeded",
            "subscription_id": null,
            "tax": null,
            "total_amount": 400,
            "updated_at": null
        },
        "timestamp": new Date().toISOString()
    }
}

const genPaymentSuccess = (args: baseArgs) => {
    return {
        ...genBasePaymentData(args),
        "type": "payment.succeeded"
    }
}

const genPaymentFailed = (args: baseArgs) => {
    return {
        ...genBasePaymentData(args),
        "type": "payment.failed"
    }
}

const genPaymentProcessing = (args: baseArgs) => {
    return {
        ...genBasePaymentData(args),
        "type": "payment.processing"
    }
}

const genPaymentCancelled = (args: baseArgs) => {
    return {
        ...genBasePaymentData(args),
        "type": "payment.cancelled"
    }
}


export {
    genPaymentSuccess,
    genPaymentFailed,
    genPaymentProcessing,
    genPaymentCancelled
};