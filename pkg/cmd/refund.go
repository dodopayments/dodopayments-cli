// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"

	"github.com/dodopayments/dodopayments-cli/pkg/jsonflag"
	"github.com/dodopayments/dodopayments-go"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var refundsCreate = cli.Command{
	Name:  "create",
	Usage: "Perform create operation",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name: "payment-id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "payment_id",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "items.item_id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "items.#.item_id",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "items.amount",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "items.#.amount",
			},
		},
		&jsonflag.JSONBoolFlag{
			Name: "items.tax_inclusive",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "items.#.tax_inclusive",
				SetValue: true,
			},
		},
		&jsonflag.JSONAnyFlag{
			Name: "+item",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "items.-1",
				SetValue: map[string]interface{}{},
			},
			Value: map[string]interface{}{},
		},
		&jsonflag.JSONStringFlag{
			Name: "reason",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "reason",
			},
		},
	},
	Action:          handleRefundsCreate,
	HideHelpCommand: true,
}

var refundsRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Perform retrieve operation",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "refund-id",
		},
	},
	Action:          handleRefundsRetrieve,
	HideHelpCommand: true,
}

var refundsList = cli.Command{
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
		&jsonflag.JSONStringFlag{
			Name: "status",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "status",
			},
		},
	},
	Action:          handleRefundsList,
	HideHelpCommand: true,
}

func handleRefundsCreate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := dodopayments.RefundNewParams{}
	var res []byte
	_, err := cc.client.Refunds.New(
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
	return ShowJSON("refunds create", json, format, transform)
}

func handleRefundsRetrieve(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	var res []byte
	_, err := cc.client.Refunds.Get(
		context.TODO(),
		cmd.Value("refund-id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("refunds retrieve", json, format, transform)
}

func handleRefundsList(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := dodopayments.RefundListParams{}
	var res []byte
	_, err := cc.client.Refunds.List(
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
	return ShowJSON("refunds list", json, format, transform)
}
