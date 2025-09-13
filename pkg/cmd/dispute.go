// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"

	"github.com/dodopayments/dodopayments-go"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/stainless-sdks/dodo-payments-cli/pkg/jsonflag"
	"github.com/urfave/cli/v3"
)

var disputesRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Perform retrieve operation",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "dispute-id",
		},
	},
	Action:          handleDisputesRetrieve,
	HideHelpCommand: true,
}

var disputesList = cli.Command{
	Name:  "list",
	Usage: "Perform list operation",
	Flags: []cli.Flag{
		&jsonflag.JSONDatetimeFlag{
			Name: "created-at-gte",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "created_at_gte",
			},
		},
		&jsonflag.JSONDatetimeFlag{
			Name: "created-at-lte",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "created_at_lte",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "customer-id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "customer_id",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "dispute-stage",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "dispute_stage",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "dispute-status",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "dispute_status",
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
	},
	Action:          handleDisputesList,
	HideHelpCommand: true,
}

func handleDisputesRetrieve(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	res := []byte{}
	_, err := cc.client.Disputes.Get(
		context.TODO(),
		cmd.Value("dispute-id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("disputes retrieve", string(res), format)
}

func handleDisputesList(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := dodopayments.DisputeListParams{}
	res := []byte{}
	_, err := cc.client.Disputes.List(
		context.TODO(),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("disputes list", string(res), format)
}
