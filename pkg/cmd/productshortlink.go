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

var productsShortLinksCreate = cli.Command{
	Name:  "create",
	Usage: "Gives a Short Checkout URL with custom slug for a product. Uses a Static\nCheckout URL under the hood.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "id",
		},
		&requestflag.Flag[string]{
			Name:     "slug",
			Usage:    "Slug for the short link.",
			BodyPath: "slug",
		},
		&requestflag.Flag[map[string]string]{
			Name:     "static-checkout-params",
			Usage:    "Static Checkout URL parameters to apply to the resulting\nshort URL.",
			BodyPath: "static_checkout_params",
		},
	},
	Action:          handleProductsShortLinksCreate,
	HideHelpCommand: true,
}

var productsShortLinksList = cli.Command{
	Name:  "list",
	Usage: "Lists all short links created by the business.",
	Flags: []cli.Flag{
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
			Name:      "product-id",
			Usage:     "Filter by product ID",
			QueryPath: "product_id",
		},
	},
	Action:          handleProductsShortLinksList,
	HideHelpCommand: true,
}

func handleProductsShortLinksCreate(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := dodopayments.ProductShortLinkNewParams{}

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
	_, err = client.Products.ShortLinks.New(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "products:short-links create", obj, format, transform)
}

func handleProductsShortLinksList(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := dodopayments.ProductShortLinkListParams{}

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
		_, err = client.Products.ShortLinks.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "products:short-links list", obj, format, transform)
	} else {
		iter := client.Products.ShortLinks.ListAutoPaging(ctx, params, options...)
		return streamOutput("products:short-links list", func(w *os.File) error {
			for iter.Next() {
				item := iter.Current()
				obj := gjson.Parse(item.JSON.RawJSON())
				if err := ShowJSON(w, "products:short-links list", obj, format, transform); err != nil {
					return err
				}
			}
			return iter.Err()
		})
	}
}
