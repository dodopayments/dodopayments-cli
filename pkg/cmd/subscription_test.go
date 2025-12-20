// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/dodopayments/dodopayments-cli/internal/mocktest"
)

func TestSubscriptionsCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"subscriptions", "create",
		"--billing", "{country: AF, city: city, state: state, street: street, zipcode: zipcode}",
		"--customer", "{customer_id: customer_id}",
		"--product-id", "product_id",
		"--quantity", "0",
		"--addon", "{addon_id: addon_id, quantity: 0}",
		"--allowed-payment-method-type", "credit",
		"--billing-currency", "AED",
		"--discount-code", "discount_code",
		"--force-3ds",
		"--metadata", "{foo: string}",
		"--on-demand", "{mandate_only: true, adaptive_currency_fees_inclusive: true, product_currency: AED, product_description: product_description, product_price: 0}",
		"--one-time-product-cart", "{product_id: product_id, quantity: 0, amount: 0}",
		"--payment-link",
		"--return-url", "return_url",
		"--show-saved-payment-methods",
		"--tax-id", "tax_id",
		"--trial-period-days", "0",
	)
}

func TestSubscriptionsRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"subscriptions", "retrieve",
		"--subscription-id", "subscription_id",
	)
}

func TestSubscriptionsUpdate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"subscriptions", "update",
		"--subscription-id", "subscription_id",
		"--billing", "{country: AF, city: city, state: state, street: street, zipcode: zipcode}",
		"--cancel-at-next-billing-date",
		"--customer-name", "customer_name",
		"--disable-on-demand", "{next_billing_date: '2019-12-27T18:11:19.117Z'}",
		"--metadata", "{foo: string}",
		"--next-billing-date", "2019-12-27T18:11:19.117Z",
		"--status", "pending",
		"--tax-id", "tax_id",
	)
}

func TestSubscriptionsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"subscriptions", "list",
		"--brand-id", "brand_id",
		"--created-at-gte", "2019-12-27T18:11:19.117Z",
		"--created-at-lte", "2019-12-27T18:11:19.117Z",
		"--customer-id", "customer_id",
		"--page-number", "0",
		"--page-size", "0",
		"--status", "pending",
	)
}

func TestSubscriptionsChangePlan(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"subscriptions", "change-plan",
		"--subscription-id", "subscription_id",
		"--product-id", "product_id",
		"--proration-billing-mode", "prorated_immediately",
		"--quantity", "0",
		"--addon", "{addon_id: addon_id, quantity: 0}",
	)
}

func TestSubscriptionsCharge(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"subscriptions", "charge",
		"--subscription-id", "subscription_id",
		"--product-price", "0",
		"--adaptive-currency-fees-inclusive",
		"--customer-balance-config", "{allow_customer_credits_purchase: true, allow_customer_credits_usage: true}",
		"--metadata", "{foo: string}",
		"--product-currency", "AED",
		"--product-description", "product_description",
	)
}

func TestSubscriptionsPreviewChangePlan(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"subscriptions", "preview-change-plan",
		"--subscription-id", "subscription_id",
		"--product-id", "product_id",
		"--proration-billing-mode", "prorated_immediately",
		"--quantity", "0",
		"--addon", "{addon_id: addon_id, quantity: 0}",
	)
}

func TestSubscriptionsRetrieveUsageHistory(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"subscriptions", "retrieve-usage-history",
		"--subscription-id", "subscription_id",
		"--end-date", "2019-12-27T18:11:19.117Z",
		"--meter-id", "meter_id",
		"--page-number", "0",
		"--page-size", "0",
		"--start-date", "2019-12-27T18:11:19.117Z",
	)
}

func TestSubscriptionsUpdatePaymentMethod(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"subscriptions", "update-payment-method",
		"--subscription-id", "subscription_id",
		"--type", "new",
		"--return-url", "return_url",
		"--payment-method-id", "payment_method_id",
	)
}
