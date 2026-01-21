// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/dodopayments/dodopayments-cli/internal/mocktest"
)

func TestDiscountsCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"discounts", "create",
		"--amount", "0",
		"--type", "percentage",
		"--code", "code",
		"--expires-at", "2019-12-27T18:11:19.117Z",
		"--name", "name",
		"--restricted-to", "string",
		"--subscription-cycles", "0",
		"--usage-limit", "0",
	)
}

func TestDiscountsRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"discounts", "retrieve",
		"--discount-id", "discount_id",
	)
}

func TestDiscountsUpdate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"discounts", "update",
		"--discount-id", "discount_id",
		"--amount", "0",
		"--code", "code",
		"--expires-at", "2019-12-27T18:11:19.117Z",
		"--name", "name",
		"--restricted-to", "string",
		"--subscription-cycles", "0",
		"--type", "percentage",
		"--usage-limit", "0",
	)
}

func TestDiscountsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"discounts", "list",
		"--active=true",
		"--code", "code",
		"--discount-type", "percentage",
		"--page-number", "0",
		"--page-size", "0",
		"--product-id", "product_id",
	)
}

func TestDiscountsRetrieveByCode(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"discounts", "retrieve-by-code",
		"--code", "code",
	)
}
