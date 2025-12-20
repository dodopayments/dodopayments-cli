// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/dodopayments/dodopayments-cli/internal/mocktest"
)

func TestAddonsCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"addons", "create",
		"--currency", "AED",
		"--name", "name",
		"--price", "0",
		"--tax-category", "digital_products",
		"--description", "description",
	)
}

func TestAddonsRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"addons", "retrieve",
		"--id", "id",
	)
}

func TestAddonsUpdate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"addons", "update",
		"--id", "id",
		"--currency", "AED",
		"--description", "description",
		"--image-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--name", "name",
		"--price", "0",
		"--tax-category", "digital_products",
	)
}

func TestAddonsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"addons", "list",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestAddonsUpdateImages(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"addons", "update-images",
		"--id", "id",
	)
}
