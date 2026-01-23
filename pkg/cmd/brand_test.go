// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/dodopayments/dodopayments-cli/internal/mocktest"
)

func TestBrandsCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"brands", "create",
		"--description", "description",
		"--name", "name",
		"--statement-descriptor", "statement_descriptor",
		"--support-email", "support_email",
		"--url", "url",
	)
}

func TestBrandsRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"brands", "retrieve",
		"--id", "id",
	)
}

func TestBrandsUpdate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"brands", "update",
		"--id", "id",
		"--description", "description",
		"--image-id", "182bd5e5-6e1a-4fe4-a799-aa6d9a6ab26e",
		"--name", "name",
		"--statement-descriptor", "statement_descriptor",
		"--support-email", "support_email",
		"--url", "url",
	)
}

func TestBrandsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"brands", "list",
	)
}

func TestBrandsUpdateImages(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"brands", "update-images",
		"--id", "id",
	)
}
