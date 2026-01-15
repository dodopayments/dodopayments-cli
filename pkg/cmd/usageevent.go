// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

	"github.com/dodopayments/dodopayments-cli/internal/apiquery"
	"github.com/dodopayments/dodopayments-cli/internal/requestflag"
	"github.com/dodopayments/dodopayments-go"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var usageEventsRetrieve = cli.Command{
	Name:    "retrieve",
	Usage:   "Fetch detailed information about a single event using its unique event ID. This\nendpoint is useful for:",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "event-id",
			Required: true,
		},
	},
	Action:          handleUsageEventsRetrieve,
	HideHelpCommand: true,
}

var usageEventsList = cli.Command{
	Name:    "list",
	Usage:   "Fetch events from your account with powerful filtering capabilities. This\nendpoint is ideal for:",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "customer-id",
			Usage:     "Filter events by customer ID",
			QueryPath: "customer_id",
		},
		&requestflag.Flag[any]{
			Name:      "end",
			Usage:     "Filter events created before this timestamp",
			QueryPath: "end",
		},
		&requestflag.Flag[string]{
			Name:      "event-name",
			Usage:     "Filter events by event name. If both event_name and meter_id are provided, they must match the meter's configured event_name",
			QueryPath: "event_name",
		},
		&requestflag.Flag[string]{
			Name:      "meter-id",
			Usage:     "Filter events by meter ID. When provided, only events that match the meter's event_name and filter criteria will be returned",
			QueryPath: "meter_id",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "Page number (0-based, default: 0)",
			QueryPath: "page_number",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Number of events to return per page (default: 10)",
			QueryPath: "page_size",
		},
		&requestflag.Flag[any]{
			Name:      "start",
			Usage:     "Filter events created after this timestamp",
			QueryPath: "start",
		},
	},
	Action:          handleUsageEventsList,
	HideHelpCommand: true,
}

var usageEventsIngest = requestflag.WithInnerFlags(cli.Command{
	Name:    "ingest",
	Usage:   "This endpoint allows you to ingest custom events that can be used for:",
	Suggest: true,
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "event",
			Usage:    "List of events to be pushed",
			Required: true,
			BodyPath: "events",
		},
	},
	Action:          handleUsageEventsIngest,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"event": {
		&requestflag.InnerFlag[string]{
			Name:       "event.customer-id",
			Usage:      "customer_id of the customer whose usage needs to be tracked",
			InnerField: "customer_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "event.event-id",
			Usage:      "Event Id acts as an idempotency key. Any subsequent requests with the same event_id will be ignored",
			InnerField: "event_id",
		},
		&requestflag.InnerFlag[string]{
			Name:       "event.event-name",
			Usage:      "Name of the event",
			InnerField: "event_name",
		},
		&requestflag.InnerFlag[map[string]any]{
			Name:       "event.metadata",
			Usage:      "Custom metadata. Only key value pairs are accepted, objects or arrays submitted will be rejected.",
			InnerField: "metadata",
		},
		&requestflag.InnerFlag[any]{
			Name:       "event.timestamp",
			Usage:      "Custom Timestamp. Defaults to current timestamp in UTC.\nTimestamps that are older that 1 hour or after 5 mins, from current timestamp, will be rejected.",
			InnerField: "timestamp",
		},
	},
})

func handleUsageEventsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("event-id") && len(unusedArgs) > 0 {
		cmd.Set("event-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.UsageEvents.Get(ctx, cmd.Value("event-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "usage-events retrieve", obj, format, transform)
}

func handleUsageEventsList(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := dodopayments.UsageEventListParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.UsageEvents.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "usage-events list", obj, format, transform)
	} else {
		iter := client.UsageEvents.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "usage-events list", iter, format, transform)
	}
}

func handleUsageEventsIngest(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := dodopayments.UsageEventIngestParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.UsageEvents.Ingest(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "usage-events ingest", obj, format, transform)
}
