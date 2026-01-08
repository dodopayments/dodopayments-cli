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

var productsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:  "create",
	Usage: "Perform create operation",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the product",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "price",
			Usage:    "One-time price details.",
			Required: true,
			BodyPath: "price",
		},
		&requestflag.Flag[string]{
			Name:     "tax-category",
			Usage:    "Represents the different categories of taxation applicable to various products and services.",
			Required: true,
			BodyPath: "tax_category",
		},
		&requestflag.Flag[[]string]{
			Name:     "addon",
			Usage:    "Addons available for subscription product",
			BodyPath: "addons",
		},
		&requestflag.Flag[string]{
			Name:     "brand-id",
			Usage:    "Brand id for the product, if not provided will default to primary brand",
			BodyPath: "brand_id",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Optional description of the product",
			BodyPath: "description",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "digital-product-delivery",
			Usage:    "Choose how you would like you digital product delivered",
			BodyPath: "digital_product_delivery",
		},
		&requestflag.Flag[string]{
			Name:     "license-key-activation-message",
			Usage:    "Optional message displayed during license key activation",
			BodyPath: "license_key_activation_message",
		},
		&requestflag.Flag[int64]{
			Name:     "license-key-activations-limit",
			Usage:    "The number of times the license key can be activated.\nMust be 0 or greater",
			BodyPath: "license_key_activations_limit",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "license-key-duration",
			BodyPath: "license_key_duration",
		},
		&requestflag.Flag[bool]{
			Name:     "license-key-enabled",
			Usage:    "When true, generates and sends a license key to your customer.\nDefaults to false",
			BodyPath: "license_key_enabled",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "Additional metadata for the product",
			BodyPath: "metadata",
		},
	},
	Action:          handleProductsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"digital-product-delivery": {
		&requestflag.InnerFlag[string]{
			Name:       "digital-product-delivery.external-url",
			Usage:      "External URL to digital product",
			InnerField: "external_url",
		},
		&requestflag.InnerFlag[string]{
			Name:       "digital-product-delivery.instructions",
			Usage:      "Instructions to download and use the digital product",
			InnerField: "instructions",
		},
	},
	"license-key-duration": {
		&requestflag.InnerFlag[int64]{
			Name:       "license-key-duration.count",
			InnerField: "count",
		},
		&requestflag.InnerFlag[string]{
			Name:       "license-key-duration.interval",
			InnerField: "interval",
		},
	},
})

var productsRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Perform retrieve operation",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleProductsRetrieve,
	HideHelpCommand: true,
}

var productsUpdate = requestflag.WithInnerFlags(cli.Command{
	Name:  "update",
	Usage: "Perform update operation",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[[]string]{
			Name:     "addon",
			Usage:    "Available Addons for subscription products",
			BodyPath: "addons",
		},
		&requestflag.Flag[string]{
			Name:     "brand-id",
			BodyPath: "brand_id",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Description of the product, optional and must be at most 1000 characters.",
			BodyPath: "description",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "digital-product-delivery",
			Usage:    "Choose how you would like you digital product delivered",
			BodyPath: "digital_product_delivery",
		},
		&requestflag.Flag[string]{
			Name:     "image-id",
			Usage:    "Product image id after its uploaded to S3",
			BodyPath: "image_id",
		},
		&requestflag.Flag[string]{
			Name:     "license-key-activation-message",
			Usage:    "Message sent to the customer upon license key activation.\n\nOnly applicable if `license_key_enabled` is `true`. This message contains instructions for\nactivating the license key.",
			BodyPath: "license_key_activation_message",
		},
		&requestflag.Flag[int64]{
			Name:     "license-key-activations-limit",
			Usage:    "Limit for the number of activations for the license key.\n\nOnly applicable if `license_key_enabled` is `true`. Represents the maximum number of times\nthe license key can be activated.",
			BodyPath: "license_key_activations_limit",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "license-key-duration",
			BodyPath: "license_key_duration",
		},
		&requestflag.Flag[bool]{
			Name:     "license-key-enabled",
			Usage:    "Whether the product requires a license key.\n\nIf `true`, additional fields related to license key (duration, activations limit, activation message)\nbecome applicable.",
			BodyPath: "license_key_enabled",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "Additional metadata for the product",
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the product, optional and must be at most 100 characters.",
			BodyPath: "name",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "price",
			Usage:    "One-time price details.",
			BodyPath: "price",
		},
		&requestflag.Flag[string]{
			Name:     "tax-category",
			Usage:    "Represents the different categories of taxation applicable to various products and services.",
			BodyPath: "tax_category",
		},
	},
	Action:          handleProductsUpdate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"digital-product-delivery": {
		&requestflag.InnerFlag[string]{
			Name:       "digital-product-delivery.external-url",
			Usage:      "External URL to digital product",
			InnerField: "external_url",
		},
		&requestflag.InnerFlag[[]string]{
			Name:       "digital-product-delivery.files",
			Usage:      "Uploaded files ids of digital product",
			InnerField: "files",
		},
		&requestflag.InnerFlag[string]{
			Name:       "digital-product-delivery.instructions",
			Usage:      "Instructions to download and use the digital product",
			InnerField: "instructions",
		},
	},
	"license-key-duration": {
		&requestflag.InnerFlag[int64]{
			Name:       "license-key-duration.count",
			InnerField: "count",
		},
		&requestflag.InnerFlag[string]{
			Name:       "license-key-duration.interval",
			InnerField: "interval",
		},
	},
})

var productsList = cli.Command{
	Name:  "list",
	Usage: "Perform list operation",
	Flags: []cli.Flag{
		&requestflag.Flag[bool]{
			Name:      "archived",
			Usage:     "List archived products",
			QueryPath: "archived",
		},
		&requestflag.Flag[string]{
			Name:      "brand-id",
			Usage:     "filter by Brand id",
			QueryPath: "brand_id",
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
		&requestflag.Flag[bool]{
			Name:      "recurring",
			Usage:     "Filter products by pricing type:\n- `true`: Show only recurring pricing products (e.g. subscriptions)\n- `false`: Show only one-time price products\n- `null` or absent: Show both types of products",
			QueryPath: "recurring",
		},
	},
	Action:          handleProductsList,
	HideHelpCommand: true,
}

var productsArchive = cli.Command{
	Name:  "archive",
	Usage: "Perform archive operation",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleProductsArchive,
	HideHelpCommand: true,
}

var productsUnarchive = cli.Command{
	Name:  "unarchive",
	Usage: "Perform unarchive operation",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleProductsUnarchive,
	HideHelpCommand: true,
}

var productsUpdateFiles = cli.Command{
	Name:  "update-files",
	Usage: "Perform update-files operation",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
		&requestflag.Flag[string]{
			Name:     "file-name",
			Required: true,
			BodyPath: "file_name",
		},
	},
	Action:          handleProductsUpdateFiles,
	HideHelpCommand: true,
}

func handleProductsCreate(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := dodopayments.ProductNewParams{}

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
	_, err = client.Products.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "products create", obj, format, transform)
}

func handleProductsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
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
	_, err = client.Products.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "products retrieve", obj, format, transform)
}

func handleProductsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := dodopayments.ProductUpdateParams{}

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

	return client.Products.Update(
		ctx,
		cmd.Value("id").(string),
		params,
		options...,
	)
}

func handleProductsList(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := dodopayments.ProductListParams{}

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
		_, err = client.Products.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "products list", obj, format, transform)
	} else {
		iter := client.Products.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "products list", iter, format, transform)
	}
}

func handleProductsArchive(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
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

	return client.Products.Archive(ctx, cmd.Value("id").(string), options...)
}

func handleProductsUnarchive(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
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

	return client.Products.Unarchive(ctx, cmd.Value("id").(string), options...)
}

func handleProductsUpdateFiles(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}

	params := dodopayments.ProductUpdateFilesParams{}

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
	_, err = client.Products.UpdateFiles(
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
	return ShowJSON(os.Stdout, "products update-files", obj, format, transform)
}
