// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/dodopayments/dodopayments-cli/internal/mocktest"
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
		"--force-3ds",
		"--metadata", "{foo: string}",
		"--payment-link",
		"--payment-method-id", "payment_method_id",
		"--redirect-immediately",
		"--return-url", "return_url",
		"--short-link",
		"--show-saved-payment-methods",
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
