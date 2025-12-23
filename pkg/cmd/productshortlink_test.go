// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/dodopayments/dodopayments-cli/internal/mocktest"
)

func TestProductsShortLinksCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"products:short-links", "create",
		"--id", "id",
		"--slug", "slug",
		"--static-checkout-params", "{foo: string}",
	)
}

func TestProductsShortLinksList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"products:short-links", "list",
		"--page-number", "0",
		"--page-size", "0",
		"--product-id", "product_id",
	)
}
