// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/dodopayments/dodopayments-cli/pkg/jsonflag"
	"github.com/dodopayments/dodopayments-go"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var usageEventsRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Fetch detailed information about a single event using its unique event ID. This\nendpoint is useful for:",
	Flags: []cli.Flag{
		&cli.StringFlag{
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
		&jsonflag.JSONStringFlag{
			Name: "customer-id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "customer_id",
			},
		},
		&jsonflag.JSONDatetimeFlag{
			Name: "end",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "end",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "event-name",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "event_name",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "meter-id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "meter_id",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "page-number",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "page_number",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "page-size",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "page_size",
			},
		},
		&jsonflag.JSONDatetimeFlag{
			Name: "start",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "start",
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
		&jsonflag.JSONStringFlag{
			Name: "events.customer_id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "events.#.customer_id",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "events.event_id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "events.#.event_id",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "events.event_name",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "events.#.event_name",
			},
		},
		&jsonflag.JSONDatetimeFlag{
			Name: "events.timestamp",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "events.#.timestamp",
			},
		},
		&jsonflag.JSONAnyFlag{
			Name: "+event",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "events.-1",
				SetValue: map[string]interface{}{},
			},
			Value: map[string]interface{}{},
		},
	},
	Action:          handleUsageEventsIngest,
	HideHelpCommand: true,
}

func handleUsageEventsRetrieve(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("event-id") && len(unusedArgs) > 0 {
		cmd.Set("event-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	var res []byte
	_, err := cc.client.UsageEvents.Get(
		context.TODO(),
		cmd.Value("event-id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
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
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.UsageEventListParams{}
	var res []byte
	_, err := cc.client.UsageEvents.List(
		context.TODO(),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
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
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.UsageEventIngestParams{}
	var res []byte
	_, err := cc.client.UsageEvents.Ingest(
		context.TODO(),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("usage-events ingest", json, format, transform)
}
