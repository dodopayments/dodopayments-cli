// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/dodopayments/dodopayments-cli/internal/mocktest"
)

func TestCustomersWalletsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"customers:wallets", "list",
		"--customer-id", "customer_id",
	)
}
