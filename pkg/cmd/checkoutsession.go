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

var checkoutSessionsCreate = cli.Command{
	Name:  "create",
	Usage: "Perform create operation",
	Flags: []cli.Flag{
		&requestflag.YAMLSliceFlag{
			Name: "product-cart",
			Config: requestflag.RequestConfig{
				BodyPath: "product_cart",
			},
		},
		&requestflag.StringSliceFlag{
			Name:  "allowed-payment-method-type",
			Usage: "Customers will never see payment methods that are not in this list.\nHowever, adding a method here does not guarantee customers will see it.\nAvailability still depends on other factors (e.g., customer location, merchant settings).\n\nDisclaimar: Always provide 'credit' and 'debit' as a fallback.\nIf all payment methods are unavailable, checkout session will fail.",
			Config: requestflag.RequestConfig{
				BodyPath: "allowed_payment_method_types",
			},
		},
		&requestflag.YAMLFlag{
			Name:  "billing-address",
			Usage: "Billing address information for the session",
			Config: requestflag.RequestConfig{
				BodyPath: "billing_address",
			},
		},
		&requestflag.StringFlag{
			Name: "billing-currency",
			Config: requestflag.RequestConfig{
				BodyPath: "billing_currency",
			},
		},
		&requestflag.BoolFlag{
			Name:  "confirm",
			Usage: "If confirm is true, all the details will be finalized. If required data is missing, an API error is thrown.",
			Config: requestflag.RequestConfig{
				BodyPath: "confirm",
			},
		},
		&requestflag.YAMLFlag{
			Name: "customer",
			Config: requestflag.RequestConfig{
				BodyPath: "customer",
			},
		},
		&requestflag.YAMLFlag{
			Name:  "customization",
			Usage: "Customization for the checkout session page",
			Config: requestflag.RequestConfig{
				BodyPath: "customization",
			},
		},
		&requestflag.StringFlag{
			Name: "discount-code",
			Config: requestflag.RequestConfig{
				BodyPath: "discount_code",
			},
		},
		&requestflag.YAMLFlag{
			Name: "feature-flags",
			Config: requestflag.RequestConfig{
				BodyPath: "feature_flags",
			},
		},
		&requestflag.BoolFlag{
			Name:  "force-3ds",
			Usage: "Override merchant default 3DS behaviour for this session",
			Config: requestflag.RequestConfig{
				BodyPath: "force_3ds",
			},
		},
		&requestflag.YAMLFlag{
			Name:  "metadata",
			Usage: "Additional metadata associated with the payment. Defaults to empty if not provided.",
			Config: requestflag.RequestConfig{
				BodyPath: "metadata",
			},
		},
		&requestflag.BoolFlag{
			Name:  "minimal-address",
			Usage: "If true, only zipcode is required when confirm is true; other address fields remain optional",
			Config: requestflag.RequestConfig{
				BodyPath: "minimal_address",
			},
		},
		&requestflag.StringFlag{
			Name:  "return-url",
			Usage: "The url to redirect after payment failure or success.",
			Config: requestflag.RequestConfig{
				BodyPath: "return_url",
			},
		},
		&requestflag.BoolFlag{
			Name:  "show-saved-payment-methods",
			Usage: "Display saved payment methods of a returning customer False by default",
			Config: requestflag.RequestConfig{
				BodyPath: "show_saved_payment_methods",
			},
		},
		&requestflag.YAMLFlag{
			Name: "subscription-data",
			Config: requestflag.RequestConfig{
				BodyPath: "subscription_data",
			},
		},
	},
	Action:          handleCheckoutSessionsCreate,
	HideHelpCommand: true,
}

var checkoutSessionsRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Perform retrieve operation",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name: "id",
		},
	},
	Action:          handleCheckoutSessionsRetrieve,
	HideHelpCommand: true,
}

func handleCheckoutSessionsCreate(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.CheckoutSessionNewParams{}

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
	_, err = client.CheckoutSessions.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "checkout-sessions create", obj, format, transform)
}

func handleCheckoutSessionsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
		ApplicationJSON,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.CheckoutSessions.Get(ctx, requestflag.CommandRequestValue[string](cmd, "id"), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "checkout-sessions retrieve", obj, format, transform)
}
