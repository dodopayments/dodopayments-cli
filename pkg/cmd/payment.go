// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/dodopayments/dodopayments-cli/internal/apiquery"
	"github.com/dodopayments/dodopayments-cli/internal/requestflag"
	"github.com/dodopayments/dodopayments-go"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var paymentsCreate = cli.Command{
	Name:  "create",
	Usage: "Perform create operation",
	Flags: []cli.Flag{
		&requestflag.YAMLFlag{
			Name: "billing",
			Config: requestflag.RequestConfig{
				BodyPath: "billing",
			},
		},
		&requestflag.YAMLFlag{
			Name: "customer",
			Config: requestflag.RequestConfig{
				BodyPath: "customer",
			},
		},
		&requestflag.YAMLSliceFlag{
			Name:  "product-cart",
			Usage: "List of products in the cart. Must contain at least 1 and at most 100 items.",
			Config: requestflag.RequestConfig{
				BodyPath: "product_cart",
			},
		},
		&requestflag.StringSliceFlag{
			Name:  "allowed-payment-method-type",
			Usage: "List of payment methods allowed during checkout.\n\nCustomers will **never** see payment methods that are **not** in this list.\nHowever, adding a method here **does not guarantee** customers will see it.\nAvailability still depends on other factors (e.g., customer location, merchant settings).",
			Config: requestflag.RequestConfig{
				BodyPath: "allowed_payment_method_types",
			},
		},
		&requestflag.StringFlag{
			Name: "billing-currency",
			Config: requestflag.RequestConfig{
				BodyPath: "billing_currency",
			},
		},
		&requestflag.StringFlag{
			Name:  "discount-code",
			Usage: "Discount Code to apply to the transaction",
			Config: requestflag.RequestConfig{
				BodyPath: "discount_code",
			},
		},
		&requestflag.BoolFlag{
			Name:  "force-3ds",
			Usage: "Override merchant default 3DS behaviour for this payment",
			Config: requestflag.RequestConfig{
				BodyPath: "force_3ds",
			},
		},
		&requestflag.YAMLFlag{
			Name:  "metadata",
			Usage: "Additional metadata associated with the payment.\nDefaults to empty if not provided.",
			Config: requestflag.RequestConfig{
				BodyPath: "metadata",
			},
		},
		&requestflag.BoolFlag{
			Name:  "payment-link",
			Usage: "Whether to generate a payment link. Defaults to false if not specified.",
			Config: requestflag.RequestConfig{
				BodyPath: "payment_link",
			},
		},
		&requestflag.StringFlag{
			Name:  "return-url",
			Usage: "Optional URL to redirect the customer after payment.\nMust be a valid URL if provided.",
			Config: requestflag.RequestConfig{
				BodyPath: "return_url",
			},
		},
		&requestflag.BoolFlag{
			Name:  "show-saved-payment-methods",
			Usage: "Display saved payment methods of a returning customer\nFalse by default",
			Config: requestflag.RequestConfig{
				BodyPath: "show_saved_payment_methods",
			},
		},
		&requestflag.StringFlag{
			Name:  "tax-id",
			Usage: "Tax ID in case the payment is B2B. If tax id validation fails the payment creation will fail",
			Config: requestflag.RequestConfig{
				BodyPath: "tax_id",
			},
		},
	},
	Action:          handlePaymentsCreate,
	HideHelpCommand: true,
}

var paymentsRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Perform retrieve operation",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name: "payment-id",
		},
	},
	Action:          handlePaymentsRetrieve,
	HideHelpCommand: true,
}

var paymentsList = cli.Command{
	Name:  "list",
	Usage: "Perform list operation",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name:  "brand-id",
			Usage: "filter by Brand id",
			Config: requestflag.RequestConfig{
				QueryPath: "brand_id",
			},
		},
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
			Usage: "Filter by customer id",
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
		&requestflag.StringFlag{
			Name:  "subscription-id",
			Usage: "Filter by subscription id",
			Config: requestflag.RequestConfig{
				QueryPath: "subscription_id",
			},
		},
	},
	Action:          handlePaymentsList,
	HideHelpCommand: true,
}

var paymentsRetrieveLineItems = cli.Command{
	Name:  "retrieve-line-items",
	Usage: "Perform retrieve-line-items operation",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name: "payment-id",
		},
	},
	Action:          handlePaymentsRetrieveLineItems,
	HideHelpCommand: true,
}

func handlePaymentsCreate(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.PaymentNewParams{}

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
	_, err = client.Payments.New(
		ctx,
		params,
		options...,
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("payments create", json, format, transform)
}

func handlePaymentsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("payment-id") && len(unusedArgs) > 0 {
		cmd.Set("payment-id", unusedArgs[0])
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
	_, err = client.Payments.Get(
		ctx,
		requestflag.CommandRequestValue[string](cmd, "payment-id"),
		options...,
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("payments retrieve", json, format, transform)
}

func handlePaymentsList(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.PaymentListParams{}

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
	_, err = client.Payments.List(
		ctx,
		params,
		options...,
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("payments list", json, format, transform)
}

func handlePaymentsRetrieveLineItems(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("payment-id") && len(unusedArgs) > 0 {
		cmd.Set("payment-id", unusedArgs[0])
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
	_, err = client.Payments.GetLineItems(
		ctx,
		requestflag.CommandRequestValue[string](cmd, "payment-id"),
		options...,
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("payments retrieve-line-items", json, format, transform)
}
