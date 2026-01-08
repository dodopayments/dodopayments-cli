// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/dodopayments/dodopayments-cli/internal/mocktest"
	"github.com/dodopayments/dodopayments-cli/internal/requestflag"
)

func TestPaymentsCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"payments", "create",
		"--billing", "{country: AF, city: city, state: state, street: street, zipcode: zipcode}",
		"--customer", "{customer_id: customer_id}",
		"--product-cart", "{product_id: product_id, quantity: 0, amount: 0}",
		"--allowed-payment-method-type", "ach",
		"--billing-currency", "AED",
		"--discount-code", "discount_code",
		"--force-3ds=true",
		"--metadata", "{foo: string}",
		"--payment-link=true",
		"--payment-method-id", "payment_method_id",
		"--redirect-immediately=true",
		"--return-url", "return_url",
		"--short-link=true",
		"--show-saved-payment-methods=true",
		"--tax-id", "tax_id",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(paymentsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"payments", "create",
		"--billing.country", "AF",
		"--billing.city", "city",
		"--billing.state", "state",
		"--billing.street", "street",
		"--billing.zipcode", "zipcode",
		"--customer", "{customer_id: customer_id}",
		"--product-cart.product_id", "product_id",
		"--product-cart.quantity", "0",
		"--product-cart.amount", "0",
		"--allowed-payment-method-type", "ach",
		"--billing-currency", "AED",
		"--discount-code", "discount_code",
		"--force-3ds=true",
		"--metadata", "{foo: string}",
		"--payment-link=true",
		"--payment-method-id", "payment_method_id",
		"--redirect-immediately=true",
		"--return-url", "return_url",
		"--short-link=true",
		"--show-saved-payment-methods=true",
		"--tax-id", "tax_id",
	)
}

func TestPaymentsRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"payments", "retrieve",
		"--payment-id", "payment_id",
	)
}

func TestPaymentsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"payments", "list",
		"--brand-id", "brand_id",
		"--created-at-gte", "2019-12-27T18:11:19.117Z",
		"--created-at-lte", "2019-12-27T18:11:19.117Z",
		"--customer-id", "customer_id",
		"--page-number", "0",
		"--page-size", "0",
		"--status", "succeeded",
		"--subscription-id", "subscription_id",
	)
}

func TestPaymentsRetrieveLineItems(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"payments", "retrieve-line-items",
		"--payment-id", "payment_id",
	)
}
