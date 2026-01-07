// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/dodopayments/dodopayments-cli/internal/mocktest"
)

func TestCheckoutSessionsCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"checkout-sessions", "create",
		"--product-cart", "{product_id: product_id, quantity: 0, addons: [{addon_id: addon_id, quantity: 0}], amount: 0}",
		"--allowed-payment-method-type", "ach",
		"--billing-address", "{country: AF, city: city, state: state, street: street, zipcode: zipcode}",
		"--billing-currency", "AED",
		"--confirm",
		"--customer", "{customer_id: customer_id}",
		"--customization", "{force_language: force_language, show_on_demand_tag: true, show_order_details: true, theme: dark}",
		"--discount-code", "discount_code",
		"--feature-flags", "{allow_currency_selection: true, allow_customer_editing_city: true, allow_customer_editing_country: true, allow_customer_editing_email: true, allow_customer_editing_name: true, allow_customer_editing_state: true, allow_customer_editing_street: true, allow_customer_editing_zipcode: true, allow_discount_code: true, allow_phone_number_collection: true, allow_tax_id: true, always_create_new_customer: true, redirect_immediately: true}",
		"--force-3ds",
		"--metadata", "{foo: string}",
		"--minimal-address",
		"--payment-method-id", "payment_method_id",
		"--return-url", "return_url",
		"--short-link",
		"--show-saved-payment-methods",
		"--subscription-data", "{on_demand: {mandate_only: true, adaptive_currency_fees_inclusive: true, product_currency: AED, product_description: product_description, product_price: 0}, trial_period_days: 0}",
	)
}

func TestCheckoutSessionsRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"checkout-sessions", "retrieve",
		"--id", "id",
	)
}
