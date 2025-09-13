// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"

	"github.com/dodopayments/dodopayments-cli/pkg/jsonflag"
	"github.com/dodopayments/dodopayments-go"
	"github.com/dodopayments/dodopayments-go/option"
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
	res := []byte{}
	_, err := cc.client.UsageEvents.Get(
		context.TODO(),
		cmd.Value("event-id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("usage-events retrieve", string(res), format)
}

func handleUsageEventsList(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := dodopayments.UsageEventListParams{}
	res := []byte{}
	_, err := cc.client.UsageEvents.List(
		context.TODO(),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("usage-events list", string(res), format)
}

func handleUsageEventsIngest(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := dodopayments.UsageEventIngestParams{}
	res := []byte{}
	_, err := cc.client.UsageEvents.Ingest(
		context.TODO(),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("usage-events ingest", string(res), format)
}
