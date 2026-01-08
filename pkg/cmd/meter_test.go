// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"testing"

	"github.com/dodopayments/dodopayments-cli/internal/mocktest"
	"github.com/dodopayments/dodopayments-cli/internal/requestflag"
)

func TestMetersCreate(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"meters", "create",
		"--aggregation", "{type: count, key: key}",
		"--event-name", "event_name",
		"--measurement-unit", "measurement_unit",
		"--name", "name",
		"--description", "description",
		"--filter", "{clauses: [{key: user_id, operator: equals, value: user123}, {key: amount, operator: greater_than, value: 100}], conjunction: and}",
	)

	// Check that inner flags have been set up correctly
	requestflag.CheckInnerFlags(metersCreate)

	// Alternative argument passing style using inner flags
	mocktest.TestRunMockTestWithFlags(
		t,
		"meters", "create",
		"--aggregation.type", "count",
		"--aggregation.key", "key",
		"--event-name", "event_name",
		"--measurement-unit", "measurement_unit",
		"--name", "name",
		"--description", "description",
		"--filter.clauses", "[{key: user_id, operator: equals, value: user123}, {key: amount, operator: greater_than, value: 100}]",
		"--filter.conjunction", "and",
	)
}

func TestMetersRetrieve(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"meters", "retrieve",
		"--id", "id",
	)
}

func TestMetersList(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"meters", "list",
		"--archived=true",
		"--page-number", "0",
		"--page-size", "0",
	)
}

func TestMetersArchive(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"meters", "archive",
		"--id", "id",
	)
}

func TestMetersUnarchive(t *testing.T) {
	mocktest.TestRunMockTestWithFlags(
		t,
		"meters", "unarchive",
		"--id", "id",
	)
}
