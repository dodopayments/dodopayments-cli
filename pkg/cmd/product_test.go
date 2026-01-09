// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/dodopayments/dodopayments-cli/internal/mocktest"
	"github.com/dodopayments/dodopayments-cli/internal/requestflag"
)

func TestProductsCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"products", "create",
		"--name", "name",
		"--price", "{currency: AED, discount: 0, price: 0, purchasing_power_parity: true, type: one_time_price, pay_what_you_want: true, suggested_price: 0, tax_inclusive: true}",
		"--tax-category", "digital_products",
		"--addon", "string",
		"--brand-id", "brand_id",
		"--description", "description",
		"--digital-product-delivery", "{external_url: external_url, instructions: instructions}",
		"--license-key-activation-message", "license_key_activation_message",
		"--license-key-activations-limit", "0",
		"--license-key-duration", "{count: 0, interval: Day}",
		"--license-key-enabled=true",
		"--metadata", "{foo: string}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(productsCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"products", "create",
		"--name", "name",
		"--price", "{currency: AED, discount: 0, price: 0, purchasing_power_parity: true, type: one_time_price, pay_what_you_want: true, suggested_price: 0, tax_inclusive: true}",
		"--tax-category", "digital_products",
		"--addon", "string",
		"--brand-id", "brand_id",
		"--description", "description",
		"--digital-product-delivery.external_url", "external_url",
		"--digital-product-delivery.instructions", "instructions",
		"--license-key-activation-message", "license_key_activation_message",
		"--license-key-activations-limit", "0",
		"--license-key-duration.count", "0",
		"--license-key-duration.interval", "Day",
		"--license-key-enabled=true",
		"--metadata", "{foo: string}",
	)
}

func TestProductsRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"products", "retrieve",
		"--id", "id",
	)
}

func TestProductsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"products", "list",
		"--archived=true",
		"--brand-id", "brand_id",
		"--page-number", "0",
		"--page-size", "0",
		"--recurring=true",
	)
}

func TestProductsUpdateFiles(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"products", "update-files",
		"--id", "id",
		"--file-name", "file_name",
	)
}
