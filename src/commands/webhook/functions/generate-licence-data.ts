import { type baseArgs } from "../types/baseArgs";

const genBaseLicenceData = ({
    business_id = 'bus_test',
    product_id = 'pdt_test',
    subscription_id
}: baseArgs) => {
    return {
        "business_id": business_id,
        "data": {
            "activations_limit": null,
            "business_id": business_id,
            "created_at": new Date().toISOString(),
            "customer_id": "cus_test",
            "expires_at": null,
            "id": "lic_test",
            "instances_count": 0,
            "key": "db44b22c-fe9b-4a68-bf0d-b0e0d6c6c8c0",
            "payload_type": "LicenseKey",
            "payment_id": "pay_test",
            "product_id": product_id,
            "status": "active",
            "subscription_id": subscription_id
        },
        "timestamp": new Date().toISOString()
    }
}


const genLicenceCreated = (args: baseArgs) => {
    return {
        ...genBaseLicenceData(args),
        "type": "license_key.created"
    }
}

export {
    genLicenceCreated
};