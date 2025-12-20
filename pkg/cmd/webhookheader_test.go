// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/dodopayments/dodopayments-cli/internal/mocktest"
)

func TestWebhooksHeadersRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"webhooks:headers", "retrieve",
		"--webhook-id", "webhook_id",
	)
}

func TestWebhooksHeadersUpdate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"webhooks:headers", "update",
		"--webhook-id", "webhook_id",
		"--headers", "{foo: string}",
	)
}
