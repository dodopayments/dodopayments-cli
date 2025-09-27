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
			Name: "price.currency",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.currency",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "price.discount",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.discount",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "price.price",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.price",
			},
		},
		&jsonflag.JSONBoolFlag{
			Name: "price.purchasing_power_parity",
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
			Name: "price.pay_what_you_want",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "price.pay_what_you_want",
				SetValue: true,
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "price.suggested_price",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.suggested_price",
			},
		},
		&jsonflag.JSONBoolFlag{
			Name: "price.tax_inclusive",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "price.tax_inclusive",
				SetValue: true,
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "price.payment_frequency_count",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.payment_frequency_count",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "price.payment_frequency_interval",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.payment_frequency_interval",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "price.subscription_period_count",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.subscription_period_count",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "price.subscription_period_interval",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.subscription_period_interval",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "price.trial_period_days",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.trial_period_days",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "price.fixed_price",
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
			Value: map[string]interface{}{},
		},
		&jsonflag.JSONStringFlag{
			Name: "tax-category",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "tax_category",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "addons",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "addons.#",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "+addon",
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
			Name: "description",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "description",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "digital-product-delivery.external_url",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "digital_product_delivery.external_url",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "digital-product-delivery.instructions",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "digital_product_delivery.instructions",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "license-key-activation-message",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "license_key_activation_message",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "license-key-activations-limit",
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
			Name: "license-key-enabled",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "license_key_enabled",
				SetValue: true,
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "name",
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
			Name: "addons",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "addons.#",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "+addon",
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
			Name: "description",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "description",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "digital-product-delivery.external_url",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "digital_product_delivery.external_url",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "digital-product-delivery.files",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "digital_product_delivery.files.#",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "digital-product-delivery.+file",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "digital_product_delivery.files.-1",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "digital-product-delivery.instructions",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "digital_product_delivery.instructions",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "image-id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "image_id",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "license-key-activation-message",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "license_key_activation_message",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "license-key-activations-limit",
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
			Name: "license-key-enabled",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "license_key_enabled",
				SetValue: true,
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "name",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "name",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "price.currency",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.currency",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "price.discount",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.discount",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "price.price",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.price",
			},
		},
		&jsonflag.JSONBoolFlag{
			Name: "price.purchasing_power_parity",
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
			Name: "price.pay_what_you_want",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "price.pay_what_you_want",
				SetValue: true,
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "price.suggested_price",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.suggested_price",
			},
		},
		&jsonflag.JSONBoolFlag{
			Name: "price.tax_inclusive",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "price.tax_inclusive",
				SetValue: true,
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "price.payment_frequency_count",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.payment_frequency_count",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "price.payment_frequency_interval",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.payment_frequency_interval",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "price.subscription_period_count",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.subscription_period_count",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "price.subscription_period_interval",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.subscription_period_interval",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "price.trial_period_days",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price.trial_period_days",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "price.fixed_price",
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
			Value: map[string]interface{}{},
		},
		&jsonflag.JSONStringFlag{
			Name: "tax-category",
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
			Name: "archived",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Query,
				Path:     "archived",
				SetValue: true,
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "brand-id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "brand_id",
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
		&jsonflag.JSONBoolFlag{
			Name: "recurring",
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
		context.TODO(),
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
		context.TODO(),
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
		context.TODO(),
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
		context.TODO(),
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
		context.TODO(),
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
