// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/dodopayments/dodopayments-cli/internal/mocktest"
)

func TestCustomersWalletsLedgerEntriesCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"customers:wallets:ledger-entries", "create",
		"--customer-id", "customer_id",
		"--amount", "0",
		"--currency", "AED",
		"--entry-type", "credit",
		"--idempotency-key", "idempotency_key",
		"--reason", "reason",
	)
}

func TestCustomersWalletsLedgerEntriesList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"customers:wallets:ledger-entries", "list",
		"--customer-id", "customer_id",
		"--currency", "AED",
		"--page-number", "0",
		"--page-size", "0",
	)
}
