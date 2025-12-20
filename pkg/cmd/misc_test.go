// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/dodopayments/dodopayments-cli/internal/mocktest"
)

func TestMiscListSupportedCountries(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"misc", "list-supported-countries",
	)
}
