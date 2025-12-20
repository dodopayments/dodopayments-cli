// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/dodopayments/dodopayments-cli/internal/mocktest"
)

func TestProductsImagesUpdate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"products:images", "update",
		"--id", "id",
		"--force-update",
	)
}
