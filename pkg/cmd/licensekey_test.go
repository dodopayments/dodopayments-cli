// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/dodopayments/dodopayments-cli/internal/mocktest"
)

func TestLicenseKeysRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"license-keys", "retrieve",
		"--id", "lic_123",
	)
}

func TestLicenseKeysUpdate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"license-keys", "update",
		"--id", "lic_123",
		"--activations-limit", "0",
		"--disabled=true",
		"--expires-at", "2019-12-27T18:11:19.117Z",
	)
}

func TestLicenseKeysList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"license-keys", "list",
		"--customer-id", "customer_id",
		"--page-number", "0",
		"--page-size", "0",
		"--product-id", "product_id",
		"--status", "active",
	)
}
