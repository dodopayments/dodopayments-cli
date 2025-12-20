// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/dodopayments/dodopayments-cli/internal/mocktest"
)

func TestPayoutsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"payouts", "list",
		"--created-at-gte", "2019-12-27T18:11:19.117Z",
		"--created-at-lte", "2019-12-27T18:11:19.117Z",
		"--page-number", "0",
		"--page-size", "0",
	)
}
