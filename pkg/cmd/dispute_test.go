// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/dodopayments/dodopayments-cli/internal/mocktest"
)

func TestDisputesRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"disputes", "retrieve",
		"--dispute-id", "dispute_id",
	)
}

func TestDisputesList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"disputes", "list",
		"--created-at-gte", "2019-12-27T18:11:19.117Z",
		"--created-at-lte", "2019-12-27T18:11:19.117Z",
		"--customer-id", "customer_id",
		"--dispute-stage", "pre_dispute",
		"--dispute-status", "dispute_opened",
		"--page-number", "0",
		"--page-size", "0",
	)
}
