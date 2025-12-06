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

var refundsCreate = cli.Command{
	Name:  "create",
	Usage: "Perform create operation",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name:  "payment-id",
			Usage: "The unique identifier of the payment to be refunded.",
			Config: requestflag.RequestConfig{
				BodyPath: "payment_id",
			},
		},
		&requestflag.YAMLSliceFlag{
			Name:  "item",
			Usage: "Partially Refund an Individual Item",
			Config: requestflag.RequestConfig{
				BodyPath: "items",
			},
		},
		&requestflag.YAMLFlag{
			Name:  "metadata",
			Usage: "Additional metadata associated with the refund.",
			Config: requestflag.RequestConfig{
				BodyPath: "metadata",
			},
		},
		&requestflag.StringFlag{
			Name:  "reason",
			Usage: "The reason for the refund, if any. Maximum length is 3000 characters. Optional.",
			Config: requestflag.RequestConfig{
				BodyPath: "reason",
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
		&requestflag.StringFlag{
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
		&requestflag.DateTimeFlag{
			Name:  "created-at-gte",
			Usage: "Get events after this created time",
			Config: requestflag.RequestConfig{
				QueryPath: "created_at_gte",
			},
		},
		&requestflag.DateTimeFlag{
			Name:  "created-at-lte",
			Usage: "Get events created before this time",
			Config: requestflag.RequestConfig{
				QueryPath: "created_at_lte",
			},
		},
		&requestflag.StringFlag{
			Name:  "customer-id",
			Usage: "Filter by customer_id",
			Config: requestflag.RequestConfig{
				QueryPath: "customer_id",
			},
		},
		&requestflag.IntFlag{
			Name:  "page-number",
			Usage: "Page number default is 0",
			Config: requestflag.RequestConfig{
				QueryPath: "page_number",
			},
		},
		&requestflag.IntFlag{
			Name:  "page-size",
			Usage: "Page size default is 10 max is 100",
			Config: requestflag.RequestConfig{
				QueryPath: "page_size",
			},
		},
		&requestflag.StringFlag{
			Name:  "status",
			Usage: "Filter by status",
			Config: requestflag.RequestConfig{
				QueryPath: "status",
			},
		},
	},
	Action:          handleRefundsList,
	HideHelpCommand: true,
}

func handleRefundsCreate(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.RefundNewParams{}

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
	_, err = client.Refunds.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "refunds create", obj, format, transform)
}

func handleRefundsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("refund-id") && len(unusedArgs) > 0 {
		cmd.Set("refund-id", unusedArgs[0])
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
	_, err = client.Refunds.Get(ctx, requestflag.CommandRequestValue[string](cmd, "refund-id"), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "refunds retrieve", obj, format, transform)
}

func handleRefundsList(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.RefundListParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Refunds.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "refunds list", obj, format, transform)
	} else {
		iter := client.Refunds.ListAutoPaging(ctx, params, options...)
		return streamOutput("refunds list", func(w *os.File) error {
			for iter.Next() {
				item := iter.Current()
				obj := gjson.Parse(item.JSON.RawJSON())
				if err := ShowJSON(w, "refunds list", obj, format, transform); err != nil {
					return err
				}
			}
			return iter.Err()
		})
	}
}
