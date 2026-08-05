import { type baseArgs } from "../types/baseArgs"

// Grant status is capitalized on the wire: EntitlementGrantStatus is
// Pending | Delivered | Failed | Revoked, unlike every other status enum here.
const genBaseEntitlementGrantData = ({
    business_id = 'bus_test',
    metadata,
    customer_id
}: baseArgs) => {
    return {
        "business_id": business_id,
        "data": {
            "brand_id": "brd_test",
            "business_id": business_id,
            "created_at": new Date().toISOString(),
            "customer_id": customer_id,
            "delivered_at": null,
            "digital_product_delivery": null,
            "entitlement_id": "ent_test",
            "error_code": null,
            "error_message": null,
            "feature": null,
            "id": "grant_test",
            "integration_type": "license_key",
            "license_key": null,
            "metadata": metadata ?? {},
            "oauth_expires_at": null,
            "oauth_url": null,
            "payload_type": "EntitlementGrant",
            "payment_id": "pay_test",
            "revocation_reason": null,
            "revoked_at": null,
            "status": "Pending",
            "subscription_id": null,
            "updated_at": new Date().toISOString()
        },
        "timestamp": new Date().toISOString(),
    }
}

const genEntitlementGrantCreated = (args: baseArgs) => {
    return {
        ...genBaseEntitlementGrantData(args),
        "type": "entitlement_grant.created"
    }
}

const genEntitlementGrantDelivered = (args: baseArgs) => {
    const base = genBaseEntitlementGrantData(args);
    return {
        ...base,
        "type": "entitlement_grant.delivered",
        data: {
            ...base.data,
            status: "Delivered",
            delivered_at: new Date().toISOString(),
            license_key: {
                "key": "PRO-AAAA-BBBB-CCCC-DDDD",
                "expires_at": new Date(new Date().setFullYear(new Date().getFullYear() + 1)).toISOString(),
                "activations_used": 0,
                "activations_limit": 5
            }
        }
    }
}

const genEntitlementGrantFailed = (args: baseArgs) => {
    const base = genBaseEntitlementGrantData(args);
    return {
        ...base,
        "type": "entitlement_grant.failed",
        data: {
            ...base.data,
            status: "Failed",
            error_code: "license_key_generation_failed",
            error_message: "The license key could not be generated for this grant."
        }
    }
}

const genEntitlementGrantRevoked = (args: baseArgs) => {
    const base = genBaseEntitlementGrantData(args);
    return {
        ...base,
        "type": "entitlement_grant.revoked",
        data: {
            ...base.data,
            status: "Revoked",
            delivered_at: new Date().toISOString(),
            revoked_at: new Date().toISOString(),
            revocation_reason: "subscription_cancelled"
        }
    }
}

export {
    genEntitlementGrantCreated,
    genEntitlementGrantDelivered,
    genEntitlementGrantFailed,
    genEntitlementGrantRevoked
}
