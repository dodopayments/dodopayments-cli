// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/dodopayments/dodopayments-cli/internal/mocktest"
	"github.com/dodopayments/dodopayments-cli/internal/requestflag"
)

func TestCheckoutSessionsCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"checkout-sessions", "create",
		"--product-cart", "{product_id: product_id, quantity: 0, addons: [{addon_id: addon_id, quantity: 0}], amount: 0}",
		"--allowed-payment-method-type", "credit",
		"--billing-address", "{country: AF, city: city, state: state, street: street, zipcode: zipcode}",
		"--billing-currency", "AED",
		"--confirm=true",
		"--custom-field", "{field_type: text, key: key, label: label, options: [string], placeholder: placeholder, required: true}",
		"--customer", "{customer_id: customer_id}",
		"--customization", "{force_language: force_language, show_on_demand_tag: true, show_order_details: true, theme: dark}",
		"--discount-code", "discount_code",
		"--feature-flags", "{allow_currency_selection: true, allow_customer_editing_city: true, allow_customer_editing_country: true, allow_customer_editing_email: true, allow_customer_editing_name: true, allow_customer_editing_state: true, allow_customer_editing_street: true, allow_customer_editing_zipcode: true, allow_discount_code: true, allow_phone_number_collection: true, allow_tax_id: true, always_create_new_customer: true, redirect_immediately: true}",
		"--force-3ds=true",
		"--metadata", "{foo: string}",
		"--minimal-address=true",
		"--payment-method-id", "payment_method_id",
		"--product-collection-id", "product_collection_id",
		"--return-url", "return_url",
		"--short-link=true",
		"--show-saved-payment-methods=true",
		"--subscription-data", "{on_demand: {mandate_only: true, adaptive_currency_fees_inclusive: true, product_currency: AED, product_description: product_description, product_price: 0}, trial_period_days: 0}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(checkoutSessionsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"checkout-sessions", "create",
		"--product-cart.product_id", "product_id",
		"--product-cart.quantity", "0",
		"--product-cart.addons", "[{addon_id: addon_id, quantity: 0}]",
		"--product-cart.amount", "0",
		"--allowed-payment-method-type", "credit",
		"--billing-address.country", "AF",
		"--billing-address.city", "city",
		"--billing-address.state", "state",
		"--billing-address.street", "street",
		"--billing-address.zipcode", "zipcode",
		"--billing-currency", "AED",
		"--confirm=true",
		"--custom-field.field_type", "text",
		"--custom-field.key", "key",
		"--custom-field.label", "label",
		"--custom-field.options", "[string]",
		"--custom-field.placeholder", "placeholder",
		"--custom-field.required=true",
		"--customer", "{customer_id: customer_id}",
		"--customization.force_language", "force_language",
		"--customization.show_on_demand_tag=true",
		"--customization.show_order_details=true",
		"--customization.theme", "dark",
		"--discount-code", "discount_code",
		"--feature-flags.allow_currency_selection=true",
		"--feature-flags.allow_customer_editing_city=true",
		"--feature-flags.allow_customer_editing_country=true",
		"--feature-flags.allow_customer_editing_email=true",
		"--feature-flags.allow_customer_editing_name=true",
		"--feature-flags.allow_customer_editing_state=true",
		"--feature-flags.allow_customer_editing_street=true",
		"--feature-flags.allow_customer_editing_zipcode=true",
		"--feature-flags.allow_discount_code=true",
		"--feature-flags.allow_phone_number_collection=true",
		"--feature-flags.allow_tax_id=true",
		"--feature-flags.always_create_new_customer=true",
		"--feature-flags.redirect_immediately=true",
		"--force-3ds=true",
		"--metadata", "{foo: string}",
		"--minimal-address=true",
		"--payment-method-id", "payment_method_id",
		"--product-collection-id", "product_collection_id",
		"--return-url", "return_url",
		"--short-link=true",
		"--show-saved-payment-methods=true",
		"--subscription-data.on_demand", "{mandate_only: true, adaptive_currency_fees_inclusive: true, product_currency: AED, product_description: product_description, product_price: 0}",
		"--subscription-data.trial_period_days", "0",
	)
}

func TestCheckoutSessionsRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"checkout-sessions", "retrieve",
		"--id", "id",
	)
}

func TestCheckoutSessionsPreview(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"checkout-sessions", "preview",
		"--product-cart", "{product_id: product_id, quantity: 0, addons: [{addon_id: addon_id, quantity: 0}], amount: 0}",
		"--allowed-payment-method-type", "credit",
		"--billing-address", "{country: AF, city: city, state: state, street: street, zipcode: zipcode}",
		"--billing-currency", "AED",
		"--confirm=true",
		"--custom-field", "{field_type: text, key: key, label: label, options: [string], placeholder: placeholder, required: true}",
		"--customer", "{customer_id: customer_id}",
		"--customization", "{force_language: force_language, show_on_demand_tag: true, show_order_details: true, theme: dark}",
		"--discount-code", "discount_code",
		"--feature-flags", "{allow_currency_selection: true, allow_customer_editing_city: true, allow_customer_editing_country: true, allow_customer_editing_email: true, allow_customer_editing_name: true, allow_customer_editing_state: true, allow_customer_editing_street: true, allow_customer_editing_zipcode: true, allow_discount_code: true, allow_phone_number_collection: true, allow_tax_id: true, always_create_new_customer: true, redirect_immediately: true}",
		"--force-3ds=true",
		"--metadata", "{foo: string}",
		"--minimal-address=true",
		"--payment-method-id", "payment_method_id",
		"--product-collection-id", "product_collection_id",
		"--return-url", "return_url",
		"--short-link=true",
		"--show-saved-payment-methods=true",
		"--subscription-data", "{on_demand: {mandate_only: true, adaptive_currency_fees_inclusive: true, product_currency: AED, product_description: product_description, product_price: 0}, trial_period_days: 0}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(checkoutSessionsPreview)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"checkout-sessions", "preview",
		"--product-cart.product_id", "product_id",
		"--product-cart.quantity", "0",
		"--product-cart.addons", "[{addon_id: addon_id, quantity: 0}]",
		"--product-cart.amount", "0",
		"--allowed-payment-method-type", "credit",
		"--billing-address.country", "AF",
		"--billing-address.city", "city",
		"--billing-address.state", "state",
		"--billing-address.street", "street",
		"--billing-address.zipcode", "zipcode",
		"--billing-currency", "AED",
		"--confirm=true",
		"--custom-field.field_type", "text",
		"--custom-field.key", "key",
		"--custom-field.label", "label",
		"--custom-field.options", "[string]",
		"--custom-field.placeholder", "placeholder",
		"--custom-field.required=true",
		"--customer", "{customer_id: customer_id}",
		"--customization.force_language", "force_language",
		"--customization.show_on_demand_tag=true",
		"--customization.show_order_details=true",
		"--customization.theme", "dark",
		"--discount-code", "discount_code",
		"--feature-flags.allow_currency_selection=true",
		"--feature-flags.allow_customer_editing_city=true",
		"--feature-flags.allow_customer_editing_country=true",
		"--feature-flags.allow_customer_editing_email=true",
		"--feature-flags.allow_customer_editing_name=true",
		"--feature-flags.allow_customer_editing_state=true",
		"--feature-flags.allow_customer_editing_street=true",
		"--feature-flags.allow_customer_editing_zipcode=true",
		"--feature-flags.allow_discount_code=true",
		"--feature-flags.allow_phone_number_collection=true",
		"--feature-flags.allow_tax_id=true",
		"--feature-flags.always_create_new_customer=true",
		"--feature-flags.redirect_immediately=true",
		"--force-3ds=true",
		"--metadata", "{foo: string}",
		"--minimal-address=true",
		"--payment-method-id", "payment_method_id",
		"--product-collection-id", "product_collection_id",
		"--return-url", "return_url",
		"--short-link=true",
		"--show-saved-payment-methods=true",
		"--subscription-data.on_demand", "{mandate_only: true, adaptive_currency_fees_inclusive: true, product_currency: AED, product_description: product_description, product_price: 0}",
		"--subscription-data.trial_period_days", "0",
	)
}
