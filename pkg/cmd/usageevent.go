// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/dodopayments/dodopayments-cli/internal/apiquery"
	"github.com/dodopayments/dodopayments-cli/internal/requestflag"
	"github.com/dodopayments/dodopayments-go"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var usageEventsRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Fetch detailed information about a single event using its unique event ID. This\nendpoint is useful for:",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name: "event-id",
		},
	},
	Action:          handleUsageEventsRetrieve,
	HideHelpCommand: true,
}

var usageEventsList = cli.Command{
	Name:  "list",
	Usage: "Fetch events from your account with powerful filtering capabilities. This\nendpoint is ideal for:",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name:  "customer-id",
			Usage: "Filter events by customer ID",
			Config: requestflag.RequestConfig{
				QueryPath: "customer_id",
			},
		},
		&requestflag.DateTimeFlag{
			Name:  "end",
			Usage: "Filter events created before this timestamp",
			Config: requestflag.RequestConfig{
				QueryPath: "end",
			},
		},
		&requestflag.StringFlag{
			Name:  "event-name",
			Usage: "Filter events by event name. If both event_name and meter_id are provided, they must match the meter's configured event_name",
			Config: requestflag.RequestConfig{
				QueryPath: "event_name",
			},
		},
		&requestflag.StringFlag{
			Name:  "meter-id",
			Usage: "Filter events by meter ID. When provided, only events that match the meter's event_name and filter criteria will be returned",
			Config: requestflag.RequestConfig{
				QueryPath: "meter_id",
			},
		},
		&requestflag.IntFlag{
			Name:  "page-number",
			Usage: "Page number (0-based, default: 0)",
			Config: requestflag.RequestConfig{
				QueryPath: "page_number",
			},
		},
		&requestflag.IntFlag{
			Name:  "page-size",
			Usage: "Number of events to return per page (default: 10)",
			Config: requestflag.RequestConfig{
				QueryPath: "page_size",
			},
		},
		&requestflag.DateTimeFlag{
			Name:  "start",
			Usage: "Filter events created after this timestamp",
			Config: requestflag.RequestConfig{
				QueryPath: "start",
			},
		},
	},
	Action:          handleUsageEventsList,
	HideHelpCommand: true,
}

var usageEventsIngest = cli.Command{
	Name:  "ingest",
	Usage: "This endpoint allows you to ingest custom events that can be used for:",
	Flags: []cli.Flag{
		&requestflag.YAMLSliceFlag{
			Name:  "event",
			Usage: "List of events to be pushed",
			Config: requestflag.RequestConfig{
				BodyPath: "events",
			},
		},
	},
	Action:          handleUsageEventsIngest,
	HideHelpCommand: true,
}

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
		ApplicationJSON,
	)
	if err != nil {
		return err
	}
	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.UsageEvents.Get(
		ctx,
		requestflag.CommandRequestValue[string](cmd, "event-id"),
		options...,
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("usage-events retrieve", json, format, transform)
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
		ApplicationJSON,
	)
	if err != nil {
		return err
	}
	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.UsageEvents.List(
		ctx,
		params,
		options...,
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("usage-events list", json, format, transform)
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
	)
	if err != nil {
		return err
	}
	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.UsageEvents.Ingest(
		ctx,
		params,
		options...,
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("usage-events ingest", json, format, transform)
}
