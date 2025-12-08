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

var subscriptionsCreate = cli.Command{
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
		&requestflag.StringFlag{
			Name:  "product-id",
			Usage: "Unique identifier of the product to subscribe to",
			Config: requestflag.RequestConfig{
				BodyPath: "product_id",
			},
		},
		&requestflag.IntFlag{
			Name:  "quantity",
			Usage: "Number of units to subscribe for. Must be at least 1.",
			Config: requestflag.RequestConfig{
				BodyPath: "quantity",
			},
		},
		&requestflag.YAMLSliceFlag{
			Name:  "addon",
			Usage: "Attach addons to this subscription",
			Config: requestflag.RequestConfig{
				BodyPath: "addons",
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
			Usage: "Discount Code to apply to the subscription",
			Config: requestflag.RequestConfig{
				BodyPath: "discount_code",
			},
		},
		&requestflag.BoolFlag{
			Name:  "force-3ds",
			Usage: "Override merchant default 3DS behaviour for this subscription",
			Config: requestflag.RequestConfig{
				BodyPath: "force_3ds",
			},
		},
		&requestflag.YAMLFlag{
			Name:  "metadata",
			Usage: "Additional metadata for the subscription\nDefaults to empty if not specified",
			Config: requestflag.RequestConfig{
				BodyPath: "metadata",
			},
		},
		&requestflag.YAMLFlag{
			Name: "on-demand",
			Config: requestflag.RequestConfig{
				BodyPath: "on_demand",
			},
		},
		&requestflag.BoolFlag{
			Name:  "payment-link",
			Usage: "If true, generates a payment link.\nDefaults to false if not specified.",
			Config: requestflag.RequestConfig{
				BodyPath: "payment_link",
			},
		},
		&requestflag.StringFlag{
			Name:  "return-url",
			Usage: "Optional URL to redirect after successful subscription creation",
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
		&requestflag.IntFlag{
			Name:  "trial-period-days",
			Usage: "Optional trial period in days\nIf specified, this value overrides the trial period set in the product's price\nMust be between 0 and 10000 days",
			Config: requestflag.RequestConfig{
				BodyPath: "trial_period_days",
			},
		},
	},
	Action:          handleSubscriptionsCreate,
	HideHelpCommand: true,
}

var subscriptionsRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Perform retrieve operation",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name: "subscription-id",
		},
	},
	Action:          handleSubscriptionsRetrieve,
	HideHelpCommand: true,
}

var subscriptionsUpdate = cli.Command{
	Name:  "update",
	Usage: "Perform update operation",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name: "subscription-id",
		},
		&requestflag.YAMLFlag{
			Name: "billing",
			Config: requestflag.RequestConfig{
				BodyPath: "billing",
			},
		},
		&requestflag.BoolFlag{
			Name:  "cancel-at-next-billing-date",
			Usage: "When set, the subscription will remain active until the end of billing period",
			Config: requestflag.RequestConfig{
				BodyPath: "cancel_at_next_billing_date",
			},
		},
		&requestflag.StringFlag{
			Name: "customer-name",
			Config: requestflag.RequestConfig{
				BodyPath: "customer_name",
			},
		},
		&requestflag.YAMLFlag{
			Name: "disable-on-demand",
			Config: requestflag.RequestConfig{
				BodyPath: "disable_on_demand",
			},
		},
		&requestflag.YAMLFlag{
			Name: "metadata",
			Config: requestflag.RequestConfig{
				BodyPath: "metadata",
			},
		},
		&requestflag.DateTimeFlag{
			Name: "next-billing-date",
			Config: requestflag.RequestConfig{
				BodyPath: "next_billing_date",
			},
		},
		&requestflag.StringFlag{
			Name: "status",
			Config: requestflag.RequestConfig{
				BodyPath: "status",
			},
		},
		&requestflag.StringFlag{
			Name: "tax-id",
			Config: requestflag.RequestConfig{
				BodyPath: "tax_id",
			},
		},
	},
	Action:          handleSubscriptionsUpdate,
	HideHelpCommand: true,
}

var subscriptionsList = cli.Command{
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
	},
	Action:          handleSubscriptionsList,
	HideHelpCommand: true,
}

var subscriptionsChangePlan = cli.Command{
	Name:  "change-plan",
	Usage: "Perform change-plan operation",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name: "subscription-id",
		},
		&requestflag.StringFlag{
			Name:  "product-id",
			Usage: "Unique identifier of the product to subscribe to",
			Config: requestflag.RequestConfig{
				BodyPath: "product_id",
			},
		},
		&requestflag.StringFlag{
			Name:  "proration-billing-mode",
			Usage: "Proration Billing Mode",
			Config: requestflag.RequestConfig{
				BodyPath: "proration_billing_mode",
			},
		},
		&requestflag.IntFlag{
			Name:  "quantity",
			Usage: "Number of units to subscribe for. Must be at least 1.",
			Config: requestflag.RequestConfig{
				BodyPath: "quantity",
			},
		},
		&requestflag.YAMLSliceFlag{
			Name:  "addon",
			Usage: "Addons for the new plan.\nNote : Leaving this empty would remove any existing addons",
			Config: requestflag.RequestConfig{
				BodyPath: "addons",
			},
		},
	},
	Action:          handleSubscriptionsChangePlan,
	HideHelpCommand: true,
}

var subscriptionsCharge = cli.Command{
	Name:  "charge",
	Usage: "Perform charge operation",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name: "subscription-id",
		},
		&requestflag.IntFlag{
			Name:  "product-price",
			Usage: "The product price. Represented in the lowest denomination of the currency (e.g., cents for USD).\nFor example, to charge $1.00, pass `100`.",
			Config: requestflag.RequestConfig{
				BodyPath: "product_price",
			},
		},
		&requestflag.BoolFlag{
			Name:  "adaptive-currency-fees-inclusive",
			Usage: "Whether adaptive currency fees should be included in the product_price (true) or added on top (false).\nThis field is ignored if adaptive pricing is not enabled for the business.",
			Config: requestflag.RequestConfig{
				BodyPath: "adaptive_currency_fees_inclusive",
			},
		},
		&requestflag.YAMLFlag{
			Name:  "customer-balance-config",
			Usage: "Specify how customer balance is used for the payment",
			Config: requestflag.RequestConfig{
				BodyPath: "customer_balance_config",
			},
		},
		&requestflag.YAMLFlag{
			Name:  "metadata",
			Usage: "Metadata for the payment. If not passed, the metadata of the subscription will be taken",
			Config: requestflag.RequestConfig{
				BodyPath: "metadata",
			},
		},
		&requestflag.StringFlag{
			Name: "product-currency",
			Config: requestflag.RequestConfig{
				BodyPath: "product_currency",
			},
		},
		&requestflag.StringFlag{
			Name:  "product-description",
			Usage: "Optional product description override for billing and line items.\nIf not specified, the stored description of the product will be used.",
			Config: requestflag.RequestConfig{
				BodyPath: "product_description",
			},
		},
	},
	Action:          handleSubscriptionsCharge,
	HideHelpCommand: true,
}

var subscriptionsPreviewChangePlan = cli.Command{
	Name:  "preview-change-plan",
	Usage: "Perform preview-change-plan operation",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name: "subscription-id",
		},
		&requestflag.StringFlag{
			Name:  "product-id",
			Usage: "Unique identifier of the product to subscribe to",
			Config: requestflag.RequestConfig{
				BodyPath: "product_id",
			},
		},
		&requestflag.StringFlag{
			Name:  "proration-billing-mode",
			Usage: "Proration Billing Mode",
			Config: requestflag.RequestConfig{
				BodyPath: "proration_billing_mode",
			},
		},
		&requestflag.IntFlag{
			Name:  "quantity",
			Usage: "Number of units to subscribe for. Must be at least 1.",
			Config: requestflag.RequestConfig{
				BodyPath: "quantity",
			},
		},
		&requestflag.YAMLSliceFlag{
			Name:  "addon",
			Usage: "Addons for the new plan.\nNote : Leaving this empty would remove any existing addons",
			Config: requestflag.RequestConfig{
				BodyPath: "addons",
			},
		},
	},
	Action:          handleSubscriptionsPreviewChangePlan,
	HideHelpCommand: true,
}

var subscriptionsRetrieveUsageHistory = cli.Command{
	Name:  "retrieve-usage-history",
	Usage: "Get detailed usage history for a subscription that includes usage-based billing\n(metered components). This endpoint provides insights into customer usage\npatterns and billing calculations over time.",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name: "subscription-id",
		},
		&requestflag.DateTimeFlag{
			Name:  "end-date",
			Usage: "Filter by end date (inclusive)",
			Config: requestflag.RequestConfig{
				QueryPath: "end_date",
			},
		},
		&requestflag.StringFlag{
			Name:  "meter-id",
			Usage: "Filter by specific meter ID",
			Config: requestflag.RequestConfig{
				QueryPath: "meter_id",
			},
		},
		&requestflag.IntFlag{
			Name:  "page-number",
			Usage: "Page number (default: 0)",
			Config: requestflag.RequestConfig{
				QueryPath: "page_number",
			},
		},
		&requestflag.IntFlag{
			Name:  "page-size",
			Usage: "Page size (default: 10, max: 100)",
			Config: requestflag.RequestConfig{
				QueryPath: "page_size",
			},
		},
		&requestflag.DateTimeFlag{
			Name:  "start-date",
			Usage: "Filter by start date (inclusive)",
			Config: requestflag.RequestConfig{
				QueryPath: "start_date",
			},
		},
	},
	Action:          handleSubscriptionsRetrieveUsageHistory,
	HideHelpCommand: true,
}

var subscriptionsUpdatePaymentMethod = cli.Command{
	Name:  "update-payment-method",
	Usage: "Perform update-payment-method operation",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name: "subscription-id",
		},
		&requestflag.StringFlag{
			Name: "type",
			Config: requestflag.RequestConfig{
				BodyPath: "type",
			},
		},
		&requestflag.StringFlag{
			Name: "return-url",
			Config: requestflag.RequestConfig{
				BodyPath: "return_url",
			},
		},
		&requestflag.StringFlag{
			Name: "payment-method-id",
			Config: requestflag.RequestConfig{
				BodyPath: "payment_method_id",
			},
		},
	},
	Action:          handleSubscriptionsUpdatePaymentMethod,
	HideHelpCommand: true,
}

func handleSubscriptionsCreate(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.SubscriptionNewParams{}

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
	_, err = client.Subscriptions.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "subscriptions create", obj, format, transform)
}

func handleSubscriptionsRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("subscription-id") && len(unusedArgs) > 0 {
		cmd.Set("subscription-id", unusedArgs[0])
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
	_, err = client.Subscriptions.Get(ctx, requestflag.CommandRequestValue[string](cmd, "subscription-id"), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "subscriptions retrieve", obj, format, transform)
}

func handleSubscriptionsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("subscription-id") && len(unusedArgs) > 0 {
		cmd.Set("subscription-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.SubscriptionUpdateParams{}

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
	_, err = client.Subscriptions.Update(
		ctx,
		requestflag.CommandRequestValue[string](cmd, "subscription-id"),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "subscriptions update", obj, format, transform)
}

func handleSubscriptionsList(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.SubscriptionListParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Subscriptions.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "subscriptions list", obj, format, transform)
	} else {
		iter := client.Subscriptions.ListAutoPaging(ctx, params, options...)
		return streamOutput("subscriptions list", func(w *os.File) error {
			for iter.Next() {
				item := iter.Current()
				obj := gjson.Parse(item.JSON.RawJSON())
				if err := ShowJSON(w, "subscriptions list", obj, format, transform); err != nil {
					return err
				}
			}
			return iter.Err()
		})
	}
}

func handleSubscriptionsChangePlan(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("subscription-id") && len(unusedArgs) > 0 {
		cmd.Set("subscription-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.SubscriptionChangePlanParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}

	return client.Subscriptions.ChangePlan(
		ctx,
		requestflag.CommandRequestValue[string](cmd, "subscription-id"),
		params,
		options...,
	)
}

func handleSubscriptionsCharge(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("subscription-id") && len(unusedArgs) > 0 {
		cmd.Set("subscription-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.SubscriptionChargeParams{}

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
	_, err = client.Subscriptions.Charge(
		ctx,
		requestflag.CommandRequestValue[string](cmd, "subscription-id"),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "subscriptions charge", obj, format, transform)
}

func handleSubscriptionsPreviewChangePlan(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("subscription-id") && len(unusedArgs) > 0 {
		cmd.Set("subscription-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.SubscriptionPreviewChangePlanParams{}

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
	_, err = client.Subscriptions.PreviewChangePlan(
		ctx,
		requestflag.CommandRequestValue[string](cmd, "subscription-id"),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "subscriptions preview-change-plan", obj, format, transform)
}

func handleSubscriptionsRetrieveUsageHistory(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("subscription-id") && len(unusedArgs) > 0 {
		cmd.Set("subscription-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.SubscriptionGetUsageHistoryParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Subscriptions.GetUsageHistory(
			ctx,
			requestflag.CommandRequestValue[string](cmd, "subscription-id"),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "subscriptions retrieve-usage-history", obj, format, transform)
	} else {
		iter := client.Subscriptions.GetUsageHistoryAutoPaging(
			ctx,
			requestflag.CommandRequestValue[string](cmd, "subscription-id"),
			params,
			options...,
		)
		return streamOutput("subscriptions retrieve-usage-history", func(w *os.File) error {
			for iter.Next() {
				item := iter.Current()
				obj := gjson.Parse(item.JSON.RawJSON())
				if err := ShowJSON(w, "subscriptions retrieve-usage-history", obj, format, transform); err != nil {
					return err
				}
			}
			return iter.Err()
		})
	}
}

func handleSubscriptionsUpdatePaymentMethod(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("subscription-id") && len(unusedArgs) > 0 {
		cmd.Set("subscription-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.SubscriptionUpdatePaymentMethodParams{}

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
	_, err = client.Subscriptions.UpdatePaymentMethod(
		ctx,
		requestflag.CommandRequestValue[string](cmd, "subscription-id"),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "subscriptions update-payment-method", obj, format, transform)
}
