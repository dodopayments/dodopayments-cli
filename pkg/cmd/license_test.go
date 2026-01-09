// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/dodopayments/dodopayments-cli/internal/mocktest"
)

func TestLicensesActivate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"licenses", "activate",
		"--license-key", "license_key",
		"--name", "name",
	)
}

func TestLicensesValidate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"licenses", "validate",
		"--license-key", "2b1f8e2d-c41e-4e8f-b2d3-d9fd61c38f43",
		"--license-key-instance-id", "lki_123",
	)
}
