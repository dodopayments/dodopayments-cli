// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/dodopayments/dodopayments-cli/internal/mocktest"
)

func TestUsageEventsRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"usage-events", "retrieve",
		"--event-id", "event_id",
	)
}

func TestUsageEventsList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"usage-events", "list",
		"--customer-id", "customer_id",
		"--end", "2019-12-27T18:11:19.117Z",
		"--event-name", "event_name",
		"--meter-id", "meter_id",
		"--page-number", "0",
		"--page-size", "0",
		"--start", "2019-12-27T18:11:19.117Z",
	)
}

func TestUsageEventsIngest(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"usage-events", "ingest",
		"--event", "{customer_id: customer_id, event_id: event_id, event_name: event_name, metadata: {foo: string}, timestamp: '2019-12-27T18:11:19.117Z'}",
	)
}
