// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/dodopayments/dodopayments-cli/internal/mocktest"
)

func TestCustomersCustomerPortalCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"customers:customer-portal", "create",
		"--customer-id", "customer_id",
		"--send-email=true",
	)
}
