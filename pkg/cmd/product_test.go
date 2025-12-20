// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/dodopayments/dodopayments-cli/internal/mocktest"
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
		"--license-key-enabled",
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

func TestProductsUpdate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"products", "update",
		"--id", "id",
		"--addon", "string",
		"--brand-id", "brand_id",
		"--description", "description",
		"--digital-product-delivery", "{external_url: external_url, files: [182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e], instructions: instructions}",
		"--image-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--license-key-activation-message", "license_key_activation_message",
		"--license-key-activations-limit", "0",
		"--license-key-duration", "{count: 0, interval: Day}",
		"--license-key-enabled",
		"--metadata", "{foo: string}",
		"--name", "name",
		"--price", "{currency: AED, discount: 0, price: 0, purchasing_power_parity: true, type: one_time_price, pay_what_you_want: true, suggested_price: 0, tax_inclusive: true}",
		"--tax-category", "digital_products",
	)
}

func TestProductsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"products", "list",
		"--archived",
		"--brand-id", "brand_id",
		"--page-number", "0",
		"--page-size", "0",
		"--recurring",
	)
}

func TestProductsArchive(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"products", "archive",
		"--id", "id",
	)
}

func TestProductsUnarchive(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"products", "unarchive",
		"--id", "id",
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
