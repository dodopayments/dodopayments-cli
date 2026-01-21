// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/dodopayments/dodopayments-cli/internal/mocktest"
)

func TestCustomersCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"customers", "create",
		"--email", "email",
		"--name", "name",
		"--metadata", "{foo: string}",
		"--phone-number", "phone_number",
	)
}

func TestCustomersRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"customers", "retrieve",
		"--customer-id", "customer_id",
	)
}

func TestCustomersUpdate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"customers", "update",
		"--customer-id", "customer_id",
		"--metadata", "{foo: string}",
		"--name", "name",
		"--phone-number", "phone_number",
	)
}

func TestCustomersList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"customers", "list",
		"--created-at-gte", "2019-12-27T18:11:19.117Z",
		"--created-at-lte", "2019-12-27T18:11:19.117Z",
		"--email", "email",
		"--name", "name",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestCustomersRetrievePaymentMethods(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"customers", "retrieve-payment-methods",
		"--customer-id", "customer_id",
	)
}
