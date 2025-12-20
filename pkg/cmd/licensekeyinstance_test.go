// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/dodopayments/dodopayments-cli/internal/mocktest"
)

func TestLicenseKeyInstancesRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"license-key-instances", "retrieve",
		"--id", "lki_123",
	)
}

func TestLicenseKeyInstancesUpdate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"license-key-instances", "update",
		"--id", "lki_123",
		"--name", "name",
	)
}

func TestLicenseKeyInstancesList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"license-key-instances", "list",
		"--license-key-id", "license_key_id",
		"--page-number", "0",
		"--page-size", "0",
	)
}
