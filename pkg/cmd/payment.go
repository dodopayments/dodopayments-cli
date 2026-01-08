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

var paymentsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:  "create",
	Usage: "Perform create operation",
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "billing",
			Required: true,
			BodyPath: "billing",
		},
		&requestflag.Flag[map[string]string]{
			Name:     "customer",
			Required: true,
			BodyPath: "customer",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "product-cart",
			Usage:    "List of products in the cart. Must contain at least 1 and at most 100 items.",
			Required: true,
			BodyPath: "product_cart",
		},
		&requestflag.Flag[[]string]{
			Name:     "allowed-payment-method-type",
			Usage:    "List of payment methods allowed during checkout.\n\nCustomers will **never** see payment methods that are **not** in this list.\nHowever, adding a method here **does not guarantee** customers will see it.\nAvailability still depends on other factors (e.g., customer location, merchant settings).",
			BodyPath: "allowed_payment_method_types",
		},
		&requestflag.Flag[string]{
			Name:     "billing-currency",
			BodyPath: "billing_currency",
		},
		&requestflag.Flag[string]{
			Name:     "discount-code",
			Usage:    "Discount Code to apply to the transaction",
			BodyPath: "discount_code",
		},
		&requestflag.Flag[bool]{
			Name:     "force-3ds",
			Usage:    "Override merchant default 3DS behaviour for this payment",
			BodyPath: "force_3ds",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "metadata",
			Usage:    "Additional metadata associated with the payment.\nDefaults to empty if not provided.",
			BodyPath: "metadata",
		},
		&requestflag.Flag[bool]{
			Name:     "payment-link",
			Usage:    "Whether to generate a payment link. Defaults to false if not specified.",
			BodyPath: "payment_link",
		},
		&requestflag.Flag[string]{
			Name:     "payment-method-id",
			Usage:    "Optional payment method ID to use for this payment.\nIf provided, customer_id must also be provided.\nThe payment method will be validated for eligibility with the payment's currency.",
			BodyPath: "payment_method_id",
		},
		&requestflag.Flag[bool]{
			Name:     "redirect-immediately",
			Usage:    "If true, redirects the customer immediately after payment completion\nFalse by default",
			BodyPath: "redirect_immediately",
		},
		&requestflag.Flag[string]{
			Name:     "return-url",
			Usage:    "Optional URL to redirect the customer after payment.\nMust be a valid URL if provided.",
			BodyPath: "return_url",
		},
		&requestflag.Flag[bool]{
			Name:     "short-link",
			Usage:    "If true, returns a shortened payment link.\nDefaults to false if not specified.",
			BodyPath: "short_link",
		},
		&requestflag.Flag[bool]{
			Name:     "show-saved-payment-methods",
			Usage:    "Display saved payment methods of a returning customer\nFalse by default",
			BodyPath: "show_saved_payment_methods",
		},
		&requestflag.Flag[string]{
			Name:     "tax-id",
			Usage:    "Tax ID in case the payment is B2B. If tax id validation fails the payment creation will fail",
			BodyPath: "tax_id",
		},
	},
	Action:          handlePaymentsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"billing": {
		&requestflag.InnerFlag[string]{
			Name:       "billing.country",
			Usage:      "ISO country code alpha2 variant",
			InnerField: "country",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing.city",
			Usage:      "City name",
			InnerField: "city",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing.state",
			Usage:      "State or province name",
			InnerField: "state",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing.street",
			Usage:      "Street address including house number and unit/apartment if applicable",
			InnerField: "street",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing.zipcode",
			Usage:      "Postal code or ZIP code",
			InnerField: "zipcode",
		},
	},
	"product-cart": {
		&requestflag.InnerFlag[string]{
			Name:       "product-cart.product-id",
			InnerField: "product_id",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "product-cart.quantity",
			InnerField: "quantity",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "product-cart.amount",
			Usage:      "Amount the customer pays if pay_what_you_want is enabled. If disabled then amount will be ignored\nRepresented in the lowest denomination of the currency (e.g., cents for USD).\nFor example, to charge $1.00, pass `100`.",
			InnerField: "amount",
		},
	},
})

var paymentsRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Perform retrieve operation",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "payment-id",
			Required: true,
		},
	},
	Action:          handlePaymentsRetrieve,
	HideHelpCommand: true,
}

var paymentsList = cli.Command{
	Name:  "list",
	Usage: "Perform list operation",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:      "brand-id",
			Usage:     "filter by Brand id",
			QueryPath: "brand_id",
		},
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
			Usage:     "Filter by customer id",
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
		&requestflag.Flag[string]{
			Name:      "subscription-id",
			Usage:     "Filter by subscription id",
			QueryPath: "subscription_id",
		},
	},
	Action:          handlePaymentsList,
	HideHelpCommand: true,
}

var paymentsRetrieveLineItems = cli.Command{
	Name:  "retrieve-line-items",
	Usage: "Perform retrieve-line-items operation",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "payment-id",
			Required: true,
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
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Payments.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "payments create", obj, format, transform)
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
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Payments.Get(ctx, cmd.Value("payment-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "payments retrieve", obj, format, transform)
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
		_, err = client.Payments.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "payments list", obj, format, transform)
	} else {
		iter := client.Payments.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "payments list", iter, format, transform)
	}
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
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Payments.GetLineItems(ctx, cmd.Value("payment-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "payments retrieve-line-items", obj, format, transform)
}
