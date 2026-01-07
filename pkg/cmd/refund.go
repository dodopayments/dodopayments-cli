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
		&requestflag.Flag[string]{
			Name:     "payment-id",
			Usage:    "The unique identifier of the payment to be refunded.",
			Required: true,
			BodyPath: "payment_id",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "item",
			Usage:    "Partially Refund an Individual Item",
			BodyPath: "items",
		},
		&requestflag.Flag[map[string]string]{
			Name:     "metadata",
			Usage:    "Additional metadata associated with the refund.",
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "reason",
			Usage:    "The reason for the refund, if any. Maximum length is 3000 characters. Optional.",
			BodyPath: "reason",
		},
	},
	Action:          handleRefundsCreate,
	HideHelpCommand: true,
}

var refundsRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Perform retrieve operation",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "refund-id",
			Required: true,
		},
	},
	Action:          handleRefundsRetrieve,
	HideHelpCommand: true,
}

var refundsList = cli.Command{
	Name:  "list",
	Usage: "Perform list operation",
	Flags: []cli.Flag{
		&requestflag.Flag[requestflag.DateTimeValue]{
			Name:      "created-at-gte",
			Usage:     "Get events after this created time",
			QueryPath: "created_at_gte",
		},
		&requestflag.Flag[requestflag.DateTimeValue]{
			Name:      "created-at-lte",
			Usage:     "Get events created before this time",
			QueryPath: "created_at_lte",
		},
		&requestflag.Flag[string]{
			Name:      "customer-id",
			Usage:     "Filter by customer_id",
			QueryPath: "customer_id",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "Page number default is 0",
			QueryPath: "page_number",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Page size default is 10 max is 100",
			QueryPath: "page_size",
		},
		&requestflag.Flag[string]{
			Name:      "status",
			Usage:     "Filter by status",
			QueryPath: "status",
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
		false,
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
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Refunds.Get(ctx, cmd.Value("refund-id").(string), options...)
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
		_, err = client.Refunds.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "refunds list", obj, format, transform)
	} else {
		iter := client.Refunds.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "refunds list", iter, format, transform)
	}
}
