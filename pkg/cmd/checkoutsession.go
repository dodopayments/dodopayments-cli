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
		&requestflag.Flag[[]any]{
			Name:     "product-cart",
			BodyPath: "product_cart",
		},
		&requestflag.Flag[[]string]{
			Name:     "allowed-payment-method-type",
			Usage:    "Customers will never see payment methods that are not in this list.\nHowever, adding a method here does not guarantee customers will see it.\nAvailability still depends on other factors (e.g., customer location, merchant settings).\n\nDisclaimar: Always provide 'credit' and 'debit' as a fallback.\nIf all payment methods are unavailable, checkout session will fail.",
			BodyPath: "allowed_payment_method_types",
		},
		&requestflag.Flag[any]{
			Name:     "billing-address",
			Usage:    "Billing address information for the session",
			BodyPath: "billing_address",
		},
		&requestflag.Flag[string]{
			Name:     "billing-currency",
			BodyPath: "billing_currency",
		},
		&requestflag.Flag[bool]{
			Name:     "confirm",
			Usage:    "If confirm is true, all the details will be finalized. If required data is missing, an API error is thrown.",
			BodyPath: "confirm",
		},
		&requestflag.Flag[any]{
			Name:     "customer",
			BodyPath: "customer",
		},
		&requestflag.Flag[any]{
			Name:     "customization",
			Usage:    "Customization for the checkout session page",
			BodyPath: "customization",
		},
		&requestflag.Flag[string]{
			Name:     "discount-code",
			BodyPath: "discount_code",
		},
		&requestflag.Flag[any]{
			Name:     "feature-flags",
			BodyPath: "feature_flags",
		},
		&requestflag.Flag[bool]{
			Name:     "force-3ds",
			Usage:    "Override merchant default 3DS behaviour for this session",
			BodyPath: "force_3ds",
		},
		&requestflag.Flag[any]{
			Name:     "metadata",
			Usage:    "Additional metadata associated with the payment. Defaults to empty if not provided.",
			BodyPath: "metadata",
		},
		&requestflag.Flag[bool]{
			Name:     "minimal-address",
			Usage:    "If true, only zipcode is required when confirm is true; other address fields remain optional",
			BodyPath: "minimal_address",
		},
		&requestflag.Flag[string]{
			Name:     "return-url",
			Usage:    "The url to redirect after payment failure or success.",
			BodyPath: "return_url",
		},
		&requestflag.Flag[bool]{
			Name:     "show-saved-payment-methods",
			Usage:    "Display saved payment methods of a returning customer False by default",
			BodyPath: "show_saved_payment_methods",
		},
		&requestflag.Flag[any]{
			Name:     "subscription-data",
			BodyPath: "subscription_data",
		},
	},
	Action:          handleCheckoutSessionsCreate,
	HideHelpCommand: true,
}

var checkoutSessionsRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Perform retrieve operation",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
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
		EmptyBody,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.CheckoutSessions.Get(ctx, cmd.Value("id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "checkout-sessions retrieve", obj, format, transform)
}
