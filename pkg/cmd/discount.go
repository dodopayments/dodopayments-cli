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

var discountsCreate = cli.Command{
	Name:  "create",
	Usage: "POST /discounts If `code` is omitted or empty, a random 16-char uppercase code\nis generated.",
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:     "amount",
			Usage:    "The discount amount.\n\n- If `discount_type` is **not** `percentage`, `amount` is in **USD cents**. For example, `100` means `$1.00`.\n  Only USD is allowed.\n- If `discount_type` **is** `percentage`, `amount` is in **basis points**. For example, `540` means `5.4%`.\n\nMust be at least 1.",
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "code",
			Usage:    "Optionally supply a code (will be uppercased).\n- Must be at least 3 characters if provided.\n- If omitted, a random 16-character code is generated.",
			BodyPath: "code",
		},
		&requestflag.Flag[requestflag.DateTimeValue]{
			Name:     "expires-at",
			Usage:    "When the discount expires, if ever.",
			BodyPath: "expires_at",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[[]string]{
			Name:     "restricted-to",
			Usage:    "List of product IDs to restrict usage (if any).",
			BodyPath: "restricted_to",
		},
		&requestflag.Flag[int64]{
			Name:     "subscription-cycles",
			Usage:    "Number of subscription billing cycles this discount is valid for.\nIf not provided, the discount will be applied indefinitely to\nall recurring payments related to the subscription.",
			BodyPath: "subscription_cycles",
		},
		&requestflag.Flag[int64]{
			Name:     "usage-limit",
			Usage:    "How many times this discount can be used (if any).\nMust be >= 1 if provided.",
			BodyPath: "usage_limit",
		},
	},
	Action:          handleDiscountsCreate,
	HideHelpCommand: true,
}

var discountsRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "GET /discounts/{discount_id}",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "discount-id",
		},
	},
	Action:          handleDiscountsRetrieve,
	HideHelpCommand: true,
}

var discountsUpdate = cli.Command{
	Name:  "update",
	Usage: "PATCH /discounts/{discount_id}",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "discount-id",
		},
		&requestflag.Flag[int64]{
			Name:     "amount",
			Usage:    "If present, update the discount amount:\n- If `discount_type` is `percentage`, this represents **basis points** (e.g., `540` = `5.4%`).\n- Otherwise, this represents **USD cents** (e.g., `100` = `$1.00`).\n\nMust be at least 1 if provided.",
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:     "code",
			Usage:    "If present, update the discount code (uppercase).",
			BodyPath: "code",
		},
		&requestflag.Flag[requestflag.DateTimeValue]{
			Name:     "expires-at",
			BodyPath: "expires_at",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			BodyPath: "name",
		},
		&requestflag.Flag[[]string]{
			Name:     "restricted-to",
			Usage:    "If present, replaces all restricted product IDs with this new set.\nTo remove all restrictions, send empty array",
			BodyPath: "restricted_to",
		},
		&requestflag.Flag[int64]{
			Name:     "subscription-cycles",
			Usage:    "Number of subscription billing cycles this discount is valid for.\nIf not provided, the discount will be applied indefinitely to\nall recurring payments related to the subscription.",
			BodyPath: "subscription_cycles",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			BodyPath: "type",
		},
		&requestflag.Flag[int64]{
			Name:     "usage-limit",
			BodyPath: "usage_limit",
		},
	},
	Action:          handleDiscountsUpdate,
	HideHelpCommand: true,
}

var discountsList = cli.Command{
	Name:  "list",
	Usage: "GET /discounts",
	Flags: []cli.Flag{
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "Page number (default = 0).",
			QueryPath: "page_number",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Page size (default = 10, max = 100).",
			QueryPath: "page_size",
		},
	},
	Action:          handleDiscountsList,
	HideHelpCommand: true,
}

var discountsDelete = cli.Command{
	Name:  "delete",
	Usage: "DELETE /discounts/{discount_id}",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "discount-id",
		},
	},
	Action:          handleDiscountsDelete,
	HideHelpCommand: true,
}

func handleDiscountsCreate(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := dodopayments.DiscountNewParams{}

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
	_, err = client.Discounts.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "discounts create", obj, format, transform)
}

func handleDiscountsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("discount-id") && len(unusedArgs) > 0 {
		cmd.Set("discount-id", unusedArgs[0])
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
	_, err = client.Discounts.Get(ctx, cmd.Value("discount-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "discounts retrieve", obj, format, transform)
}

func handleDiscountsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("discount-id") && len(unusedArgs) > 0 {
		cmd.Set("discount-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := dodopayments.DiscountUpdateParams{}

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
	_, err = client.Discounts.Update(
		ctx,
		cmd.Value("discount-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "discounts update", obj, format, transform)
}

func handleDiscountsList(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := dodopayments.DiscountListParams{}

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
		_, err = client.Discounts.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "discounts list", obj, format, transform)
	} else {
		iter := client.Discounts.ListAutoPaging(ctx, params, options...)
		return streamOutput("discounts list", func(w *os.File) error {
			for iter.Next() {
				item := iter.Current()
				obj := gjson.Parse(item.RawJSON())
				if err := ShowJSON(w, "discounts list", obj, format, transform); err != nil {
					return err
				}
			}
			return iter.Err()
		})
	}
}

func handleDiscountsDelete(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("discount-id") && len(unusedArgs) > 0 {
		cmd.Set("discount-id", unusedArgs[0])
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

	return client.Discounts.Delete(ctx, cmd.Value("discount-id").(string), options...)
}
