// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/dodopayments/dodopayments-cli/internal/mocktest"
)

func TestWebhooksCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"webhooks", "create",
		"--url", "url",
		"--description", "description",
		"--disabled=true",
		"--filter-type", "payment.succeeded",
		"--headers", "{foo: string}",
		"--idempotency-key", "idempotency_key",
		"--metadata", "{foo: string}",
		"--rate-limit", "0",
	)
}

func TestWebhooksRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"webhooks", "retrieve",
		"--webhook-id", "webhook_id",
	)
}

func TestWebhooksUpdate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"webhooks", "update",
		"--webhook-id", "webhook_id",
		"--description", "description",
		"--disabled=true",
		"--filter-type", "payment.succeeded",
		"--metadata", "{foo: string}",
		"--rate-limit", "0",
		"--url", "url",
	)
}

func TestWebhooksList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"webhooks", "list",
		"--iterator", "iterator",
		"--limit", "0",
	)
}

func TestWebhooksRetrieveSecret(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"webhooks", "retrieve-secret",
		"--webhook-id", "webhook_id",
	)
}
