// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/dodopayments/dodopayments-cli/internal/mocktest"
)

func TestRefundsCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"refunds", "create",
		"--payment-id", "payment_id",
		"--item", "{item_id: item_id, amount: 0, tax_inclusive: true}",
		"--metadata", "{foo: string}",
		"--reason", "reason",
	)
}

func TestRefundsRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"refunds", "retrieve",
		"--refund-id", "refund_id",
	)
}

func TestRefundsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"refunds", "list",
		"--created-at-gte", "2019-12-27T18:11:19.117Z",
		"--created-at-lte", "2019-12-27T18:11:19.117Z",
		"--customer-id", "customer_id",
		"--page-number", "0",
		"--page-size", "0",
		"--status", "succeeded",
	)
}
