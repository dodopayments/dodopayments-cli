// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/dodopayments/dodopayments-cli/internal/mocktest"
)

func TestInvoicesPaymentsRetrieve(t *testing.T) {
	t.Skip("Prism doesn't support application/pdf responses")
	mocktest.TestRunMockTestWithFlags(
		t,
		"invoices:payments", "retrieve",
		"--payment-id", "payment_id",
	)
}

func TestInvoicesPaymentsRetrieveRefund(t *testing.T) {
	t.Skip("Prism doesn't support application/pdf responses")
	mocktest.TestRunMockTestWithFlags(
		t,
		"invoices:payments", "retrieve-refund",
		"--refund-id", "refund_id",
	)
}
