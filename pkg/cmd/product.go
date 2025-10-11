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

var productsCreate = cli.Command{
	Name:  "create",
	Usage: "Perform create operation",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name:  "price.currency",
			Usage: "The currency in which the payment is made.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.currency",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "price.discount",
			Usage: "Discount applied to the price, represented as a percentage (0 to 100).",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.discount",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "price.price",
			Usage: "The payment amount, in the smallest denomination of the currency (e.g., cents for USD).\nFor example, to charge $1.00, pass `100`.\n\nIf [`pay_what_you_want`](Self::pay_what_you_want) is set to `true`, this field represents\nthe **minimum** amount the customer must pay.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.price",
			},
		},
		&jsonflag.JSONBoolFlag{
			Name:  "price.purchasing_power_parity",
			Usage: "Indicates if purchasing power parity adjustments are applied to the price.\nPurchasing power parity feature is not available as of now.",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "price.purchasing_power_parity",
				SetValue: true,
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "price.type",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.type",
			},
		},
		&jsonflag.JSONBoolFlag{
			Name:  "price.pay_what_you_want",
			Usage: "Indicates whether the customer can pay any amount they choose.\nIf set to `true`, the [`price`](Self::price) field is the minimum amount.",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "price.pay_what_you_want",
				SetValue: true,
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "price.suggested_price",
			Usage: "A suggested price for the user to pay. This value is only considered if\n[`pay_what_you_want`](Self::pay_what_you_want) is `true`. Otherwise, it is ignored.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.suggested_price",
			},
		},
		&jsonflag.JSONBoolFlag{
			Name:  "price.tax_inclusive",
			Usage: "Indicates if the price is tax inclusive.",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "price.tax_inclusive",
				SetValue: true,
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "price.payment_frequency_count",
			Usage: "Number of units for the payment frequency.\nFor example, a value of `1` with a `payment_frequency_interval` of `month` represents monthly payments.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.payment_frequency_count",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "price.payment_frequency_interval",
			Usage: "The time interval for the payment frequency (e.g., day, month, year).",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.payment_frequency_interval",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "price.subscription_period_count",
			Usage: "Number of units for the subscription period.\nFor example, a value of `12` with a `subscription_period_interval` of `month` represents a one-year subscription.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.subscription_period_count",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "price.subscription_period_interval",
			Usage: "The time interval for the subscription period (e.g., day, month, year).",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.subscription_period_interval",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "price.trial_period_days",
			Usage: "Number of days for the trial period. A value of `0` indicates no trial period.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.trial_period_days",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "price.fixed_price",
			Usage: "The fixed payment amount. Represented in the lowest denomination of the currency (e.g., cents for USD).\nFor example, to charge $1.00, pass `100`.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.fixed_price",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "price.meters.meter_id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.meters.#.meter_id",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "price.meters.price_per_unit",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.meters.#.price_per_unit",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "price.meters.description",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.meters.#.description",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "price.meters.free_threshold",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.meters.#.free_threshold",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "price.meters.measurement_unit",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.meters.#.measurement_unit",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "price.meters.name",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.meters.#.name",
			},
		},
		&jsonflag.JSONAnyFlag{
			Name: "price.+meter",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "price.meters.-1",
				SetValue: map[string]interface{}{},
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "tax-category",
			Usage: "Tax category applied to this product",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "tax_category",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "addons",
			Usage: "Addons available for subscription product",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "addons.#",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "+addon",
			Usage: "Addons available for subscription product",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "addons.-1",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "brand-id",
			Usage: "Brand id for the product, if not provided will default to primary brand",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "brand_id",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "description",
			Usage: "Optional description of the product",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "description",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "digital-product-delivery.external_url",
			Usage: "External URL to digital product",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "digital_product_delivery.external_url",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "digital-product-delivery.instructions",
			Usage: "Instructions to download and use the digital product",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "digital_product_delivery.instructions",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "license-key-activation-message",
			Usage: "Optional message displayed during license key activation",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "license_key_activation_message",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "license-key-activations-limit",
			Usage: "The number of times the license key can be activated.\nMust be 0 or greater",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "license_key_activations_limit",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "license-key-duration.count",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "license_key_duration.count",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "license-key-duration.interval",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "license_key_duration.interval",
			},
		},
		&jsonflag.JSONBoolFlag{
			Name:  "license-key-enabled",
			Usage: "When true, generates and sends a license key to your customer.\nDefaults to false",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "license_key_enabled",
				SetValue: true,
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "name",
			Usage: "Optional name of the product",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "name",
			},
		},
	},
	Action:          handleProductsCreate,
	HideHelpCommand: true,
}

var productsRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Perform retrieve operation",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "id",
		},
	},
	Action:          handleProductsRetrieve,
	HideHelpCommand: true,
}

var productsUpdate = cli.Command{
	Name:  "update",
	Usage: "Perform update operation",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "id",
		},
		&jsonflag.JSONStringFlag{
			Name:  "addons",
			Usage: "Available Addons for subscription products",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "addons.#",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "+addon",
			Usage: "Available Addons for subscription products",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "addons.-1",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "brand-id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "brand_id",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "description",
			Usage: "Description of the product, optional and must be at most 1000 characters.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "description",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "digital-product-delivery.external_url",
			Usage: "External URL to digital product",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "digital_product_delivery.external_url",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "digital-product-delivery.files",
			Usage: "Uploaded files ids of digital product",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "digital_product_delivery.files.#",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "digital-product-delivery.+file",
			Usage: "Uploaded files ids of digital product",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "digital_product_delivery.files.-1",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "digital-product-delivery.instructions",
			Usage: "Instructions to download and use the digital product",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "digital_product_delivery.instructions",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "image-id",
			Usage: "Product image id after its uploaded to S3",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "image_id",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "license-key-activation-message",
			Usage: "Message sent to the customer upon license key activation.\n\nOnly applicable if `license_key_enabled` is `true`. This message contains instructions for\nactivating the license key.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "license_key_activation_message",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "license-key-activations-limit",
			Usage: "Limit for the number of activations for the license key.\n\nOnly applicable if `license_key_enabled` is `true`. Represents the maximum number of times\nthe license key can be activated.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "license_key_activations_limit",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "license-key-duration.count",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "license_key_duration.count",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "license-key-duration.interval",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "license_key_duration.interval",
			},
		},
		&jsonflag.JSONBoolFlag{
			Name:  "license-key-enabled",
			Usage: "Whether the product requires a license key.\n\nIf `true`, additional fields related to license key (duration, activations limit, activation message)\nbecome applicable.",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "license_key_enabled",
				SetValue: true,
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "name",
			Usage: "Name of the product, optional and must be at most 100 characters.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "name",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "price.currency",
			Usage: "The currency in which the payment is made.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.currency",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "price.discount",
			Usage: "Discount applied to the price, represented as a percentage (0 to 100).",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.discount",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "price.price",
			Usage: "The payment amount, in the smallest denomination of the currency (e.g., cents for USD).\nFor example, to charge $1.00, pass `100`.\n\nIf [`pay_what_you_want`](Self::pay_what_you_want) is set to `true`, this field represents\nthe **minimum** amount the customer must pay.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.price",
			},
		},
		&jsonflag.JSONBoolFlag{
			Name:  "price.purchasing_power_parity",
			Usage: "Indicates if purchasing power parity adjustments are applied to the price.\nPurchasing power parity feature is not available as of now.",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "price.purchasing_power_parity",
				SetValue: true,
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "price.type",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.type",
			},
		},
		&jsonflag.JSONBoolFlag{
			Name:  "price.pay_what_you_want",
			Usage: "Indicates whether the customer can pay any amount they choose.\nIf set to `true`, the [`price`](Self::price) field is the minimum amount.",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "price.pay_what_you_want",
				SetValue: true,
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "price.suggested_price",
			Usage: "A suggested price for the user to pay. This value is only considered if\n[`pay_what_you_want`](Self::pay_what_you_want) is `true`. Otherwise, it is ignored.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.suggested_price",
			},
		},
		&jsonflag.JSONBoolFlag{
			Name:  "price.tax_inclusive",
			Usage: "Indicates if the price is tax inclusive.",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "price.tax_inclusive",
				SetValue: true,
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "price.payment_frequency_count",
			Usage: "Number of units for the payment frequency.\nFor example, a value of `1` with a `payment_frequency_interval` of `month` represents monthly payments.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.payment_frequency_count",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "price.payment_frequency_interval",
			Usage: "The time interval for the payment frequency (e.g., day, month, year).",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.payment_frequency_interval",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "price.subscription_period_count",
			Usage: "Number of units for the subscription period.\nFor example, a value of `12` with a `subscription_period_interval` of `month` represents a one-year subscription.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.subscription_period_count",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "price.subscription_period_interval",
			Usage: "The time interval for the subscription period (e.g., day, month, year).",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.subscription_period_interval",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "price.trial_period_days",
			Usage: "Number of days for the trial period. A value of `0` indicates no trial period.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.trial_period_days",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "price.fixed_price",
			Usage: "The fixed payment amount. Represented in the lowest denomination of the currency (e.g., cents for USD).\nFor example, to charge $1.00, pass `100`.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.fixed_price",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "price.meters.meter_id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.meters.#.meter_id",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "price.meters.price_per_unit",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.meters.#.price_per_unit",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "price.meters.description",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.meters.#.description",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "price.meters.free_threshold",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.meters.#.free_threshold",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "price.meters.measurement_unit",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.meters.#.measurement_unit",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "price.meters.name",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.meters.#.name",
			},
		},
		&jsonflag.JSONAnyFlag{
			Name: "price.+meter",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "price.meters.-1",
				SetValue: map[string]interface{}{},
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "tax-category",
			Usage: "Tax category of the product.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "tax_category",
			},
		},
	},
	Action:          handleProductsUpdate,
	HideHelpCommand: true,
}

var productsList = cli.Command{
	Name:  "list",
	Usage: "Perform list operation",
	Flags: []cli.Flag{
		&jsonflag.JSONBoolFlag{
			Name:  "archived",
			Usage: "List archived products",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Query,
				Path:     "archived",
				SetValue: true,
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "brand-id",
			Usage: "filter by Brand id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "brand_id",
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
		&jsonflag.JSONBoolFlag{
			Name:  "recurring",
			Usage: "Filter products by pricing type:\n- `true`: Show only recurring pricing products (e.g. subscriptions)\n- `false`: Show only one-time price products\n- `null` or absent: Show both types of products",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Query,
				Path:     "recurring",
				SetValue: true,
			},
		},
	},
	Action:          handleProductsList,
	HideHelpCommand: true,
}

var productsArchive = cli.Command{
	Name:  "archive",
	Usage: "Perform archive operation",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "id",
		},
	},
	Action:          handleProductsArchive,
	HideHelpCommand: true,
}

var productsUnarchive = cli.Command{
	Name:  "unarchive",
	Usage: "Perform unarchive operation",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "id",
		},
	},
	Action:          handleProductsUnarchive,
	HideHelpCommand: true,
}

var productsUpdateFiles = cli.Command{
	Name:  "update-files",
	Usage: "Perform update-files operation",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "id",
		},
		&jsonflag.JSONStringFlag{
			Name: "file-name",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "file_name",
			},
		},
	},
	Action:          handleProductsUpdateFiles,
	HideHelpCommand: true,
}

func handleProductsCreate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.ProductNewParams{}
	var res []byte
	_, err := cc.client.Products.New(
		ctx,
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
	return ShowJSON("products create", json, format, transform)
}

func handleProductsRetrieve(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	var res []byte
	_, err := cc.client.Products.Get(
		ctx,
		cmd.Value("id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("products retrieve", json, format, transform)
}

func handleProductsUpdate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.ProductUpdateParams{}
	return cc.client.Products.Update(
		ctx,
		cmd.Value("id").(string),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
	)
}

func handleProductsList(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.ProductListParams{}
	var res []byte
	_, err := cc.client.Products.List(
		ctx,
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
	return ShowJSON("products list", json, format, transform)
}

func handleProductsArchive(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	return cc.client.Products.Archive(
		ctx,
		cmd.Value("id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
	)
}

func handleProductsUnarchive(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	return cc.client.Products.Unarchive(
		ctx,
		cmd.Value("id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
	)
}

func handleProductsUpdateFiles(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.ProductUpdateFilesParams{}
	var res []byte
	_, err := cc.client.Products.UpdateFiles(
		ctx,
		cmd.Value("id").(string),
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
	return ShowJSON("products update-files", json, format, transform)
}
