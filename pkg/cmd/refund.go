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

var refundsCreate = cli.Command{
	Name:  "create",
	Usage: "Perform create operation",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name:  "payment-id",
			Usage: "The unique identifier of the payment to be refunded.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "payment_id",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "items.item_id",
			Usage: "Partially Refund an Individual Item",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "items.#.item_id",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "items.amount",
			Usage: "Partially Refund an Individual Item",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "items.#.amount",
			},
		},
		&jsonflag.JSONBoolFlag{
			Name:  "items.tax_inclusive",
			Usage: "Partially Refund an Individual Item",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "items.#.tax_inclusive",
				SetValue: true,
			},
		},
		&jsonflag.JSONAnyFlag{
			Name:  "+item",
			Usage: "Partially Refund an Individual Item",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "items.-1",
				SetValue: map[string]interface{}{},
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "reason",
			Usage: "The reason for the refund, if any. Maximum length is 3000 characters. Optional.",
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
			Name:  "created-at-gte",
			Usage: "Get events after this created time",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "created_at_gte",
			},
		},
		&jsonflag.JSONDatetimeFlag{
			Name:  "created-at-lte",
			Usage: "Get events created before this time",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "created_at_lte",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "customer-id",
			Usage: "Filter by customer_id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "customer_id",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "page-number",
			Usage: "Page number default is 0",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "page_number",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "page-size",
			Usage: "Page size default is 10 max is 100",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "page_size",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "status",
			Usage: "Filter by status",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "status",
			},
		},
	},
	Action:          handleRefundsList,
	HideHelpCommand: true,
}

func handleRefundsCreate(_ context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
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

func handleRefundsRetrieve(_ context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("refund-id") && len(unusedArgs) > 0 {
		cmd.Set("refund-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
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

func handleRefundsList(_ context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
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
