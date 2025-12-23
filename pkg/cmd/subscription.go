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
		&requestflag.Flag[map[string]string]{
			Name:     "billing",
			BodyPath: "billing",
		},
		&requestflag.Flag[map[string]string]{
			Name:     "customer",
			BodyPath: "customer",
		},
		&requestflag.Flag[string]{
			Name:     "product-id",
			Usage:    "Unique identifier of the product to subscribe to",
			BodyPath: "product_id",
		},
		&requestflag.Flag[int64]{
			Name:     "quantity",
			Usage:    "Number of units to subscribe for. Must be at least 1.",
			BodyPath: "quantity",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "addon",
			Usage:    "Attach addons to this subscription",
			BodyPath: "addons",
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
			Usage:    "Discount Code to apply to the subscription",
			BodyPath: "discount_code",
		},
		&requestflag.Flag[bool]{
			Name:     "force-3ds",
			Usage:    "Override merchant default 3DS behaviour for this subscription",
			BodyPath: "force_3ds",
		},
		&requestflag.Flag[map[string]string]{
			Name:     "metadata",
			Usage:    "Additional metadata for the subscription\nDefaults to empty if not specified",
			BodyPath: "metadata",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "on-demand",
			BodyPath: "on_demand",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "one-time-product-cart",
			Usage:    "List of one time products that will be bundled with the first payment for this subscription",
			BodyPath: "one_time_product_cart",
		},
		&requestflag.Flag[bool]{
			Name:     "payment-link",
			Usage:    "If true, generates a payment link.\nDefaults to false if not specified.",
			BodyPath: "payment_link",
		},
		&requestflag.Flag[string]{
			Name:     "payment-method-id",
			Usage:    "Optional payment method ID to use for this subscription.\nIf provided, customer_id must also be provided (via AttachExistingCustomer).\nThe payment method will be validated for eligibility with the subscription's currency.",
			BodyPath: "payment_method_id",
		},
		&requestflag.Flag[bool]{
			Name:     "redirect-immediately",
			Usage:    "If true, redirects the customer immediately after payment completion\nFalse by default",
			BodyPath: "redirect_immediately",
		},
		&requestflag.Flag[string]{
			Name:     "return-url",
			Usage:    "Optional URL to redirect after successful subscription creation",
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
		&requestflag.Flag[int64]{
			Name:     "trial-period-days",
			Usage:    "Optional trial period in days\nIf specified, this value overrides the trial period set in the product's price\nMust be between 0 and 10000 days",
			BodyPath: "trial_period_days",
		},
	},
	Action:          handleSubscriptionsCreate,
	HideHelpCommand: true,
}

var subscriptionsRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Perform retrieve operation",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
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
		&requestflag.Flag[string]{
			Name: "subscription-id",
		},
		&requestflag.Flag[map[string]string]{
			Name:     "billing",
			BodyPath: "billing",
		},
		&requestflag.Flag[bool]{
			Name:     "cancel-at-next-billing-date",
			Usage:    "When set, the subscription will remain active until the end of billing period",
			BodyPath: "cancel_at_next_billing_date",
		},
		&requestflag.Flag[string]{
			Name:     "customer-name",
			BodyPath: "customer_name",
		},
		&requestflag.Flag[map[string]requestflag.DateTimeValue]{
			Name:     "disable-on-demand",
			BodyPath: "disable_on_demand",
		},
		&requestflag.Flag[map[string]string]{
			Name:     "metadata",
			BodyPath: "metadata",
		},
		&requestflag.Flag[requestflag.DateTimeValue]{
			Name:     "next-billing-date",
			BodyPath: "next_billing_date",
		},
		&requestflag.Flag[string]{
			Name:     "status",
			BodyPath: "status",
		},
		&requestflag.Flag[string]{
			Name:     "tax-id",
			BodyPath: "tax_id",
		},
	},
	Action:          handleSubscriptionsUpdate,
	HideHelpCommand: true,
}

var subscriptionsList = cli.Command{
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
	},
	Action:          handleSubscriptionsList,
	HideHelpCommand: true,
}

var subscriptionsChangePlan = cli.Command{
	Name:  "change-plan",
	Usage: "Perform change-plan operation",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "subscription-id",
		},
		&requestflag.Flag[string]{
			Name:     "product-id",
			Usage:    "Unique identifier of the product to subscribe to",
			BodyPath: "product_id",
		},
		&requestflag.Flag[string]{
			Name:     "proration-billing-mode",
			Usage:    "Proration Billing Mode",
			BodyPath: "proration_billing_mode",
		},
		&requestflag.Flag[int64]{
			Name:     "quantity",
			Usage:    "Number of units to subscribe for. Must be at least 1.",
			BodyPath: "quantity",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "addon",
			Usage:    "Addons for the new plan.\nNote : Leaving this empty would remove any existing addons",
			BodyPath: "addons",
		},
	},
	Action:          handleSubscriptionsChangePlan,
	HideHelpCommand: true,
}

var subscriptionsCharge = cli.Command{
	Name:  "charge",
	Usage: "Perform charge operation",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "subscription-id",
		},
		&requestflag.Flag[int64]{
			Name:     "product-price",
			Usage:    "The product price. Represented in the lowest denomination of the currency (e.g., cents for USD).\nFor example, to charge $1.00, pass `100`.",
			BodyPath: "product_price",
		},
		&requestflag.Flag[bool]{
			Name:     "adaptive-currency-fees-inclusive",
			Usage:    "Whether adaptive currency fees should be included in the product_price (true) or added on top (false).\nThis field is ignored if adaptive pricing is not enabled for the business.",
			BodyPath: "adaptive_currency_fees_inclusive",
		},
		&requestflag.Flag[map[string]bool]{
			Name:     "customer-balance-config",
			Usage:    "Specify how customer balance is used for the payment",
			BodyPath: "customer_balance_config",
		},
		&requestflag.Flag[map[string]string]{
			Name:     "metadata",
			Usage:    "Metadata for the payment. If not passed, the metadata of the subscription will be taken",
			BodyPath: "metadata",
		},
		&requestflag.Flag[string]{
			Name:     "product-currency",
			BodyPath: "product_currency",
		},
		&requestflag.Flag[string]{
			Name:     "product-description",
			Usage:    "Optional product description override for billing and line items.\nIf not specified, the stored description of the product will be used.",
			BodyPath: "product_description",
		},
	},
	Action:          handleSubscriptionsCharge,
	HideHelpCommand: true,
}

var subscriptionsPreviewChangePlan = cli.Command{
	Name:  "preview-change-plan",
	Usage: "Perform preview-change-plan operation",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "subscription-id",
		},
		&requestflag.Flag[string]{
			Name:     "product-id",
			Usage:    "Unique identifier of the product to subscribe to",
			BodyPath: "product_id",
		},
		&requestflag.Flag[string]{
			Name:     "proration-billing-mode",
			Usage:    "Proration Billing Mode",
			BodyPath: "proration_billing_mode",
		},
		&requestflag.Flag[int64]{
			Name:     "quantity",
			Usage:    "Number of units to subscribe for. Must be at least 1.",
			BodyPath: "quantity",
		},
		&requestflag.Flag[[]map[string]any]{
			Name:     "addon",
			Usage:    "Addons for the new plan.\nNote : Leaving this empty would remove any existing addons",
			BodyPath: "addons",
		},
	},
	Action:          handleSubscriptionsPreviewChangePlan,
	HideHelpCommand: true,
}

var subscriptionsRetrieveUsageHistory = cli.Command{
	Name:  "retrieve-usage-history",
	Usage: "Get detailed usage history for a subscription that includes usage-based billing\n(metered components). This endpoint provides insights into customer usage\npatterns and billing calculations over time.",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "subscription-id",
		},
		&requestflag.Flag[requestflag.DateTimeValue]{
			Name:      "end-date",
			Usage:     "Filter by end date (inclusive)",
			QueryPath: "end_date",
		},
		&requestflag.Flag[string]{
			Name:      "meter-id",
			Usage:     "Filter by specific meter ID",
			QueryPath: "meter_id",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "Page number (default: 0)",
			QueryPath: "page_number",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Page size (default: 10, max: 100)",
			QueryPath: "page_size",
		},
		&requestflag.Flag[requestflag.DateTimeValue]{
			Name:      "start-date",
			Usage:     "Filter by start date (inclusive)",
			QueryPath: "start_date",
		},
	},
	Action:          handleSubscriptionsRetrieveUsageHistory,
	HideHelpCommand: true,
}

var subscriptionsUpdatePaymentMethod = cli.Command{
	Name:  "update-payment-method",
	Usage: "Perform update-payment-method operation",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "subscription-id",
		},
		&requestflag.Flag[string]{
			Name:     "type",
			BodyPath: "type",
		},
		&requestflag.Flag[string]{
			Name:     "return-url",
			BodyPath: "return_url",
		},
		&requestflag.Flag[string]{
			Name:     "payment-method-id",
			BodyPath: "payment_method_id",
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
		false,
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
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Subscriptions.Get(ctx, cmd.Value("subscription-id").(string), options...)
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
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Subscriptions.Update(
		ctx,
		cmd.Value("subscription-id").(string),
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
		false,
	)
	if err != nil {
		return err
	}

	return client.Subscriptions.ChangePlan(
		ctx,
		cmd.Value("subscription-id").(string),
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
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Subscriptions.Charge(
		ctx,
		cmd.Value("subscription-id").(string),
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
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Subscriptions.PreviewChangePlan(
		ctx,
		cmd.Value("subscription-id").(string),
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
		_, err = client.Subscriptions.GetUsageHistory(
			ctx,
			cmd.Value("subscription-id").(string),
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
			cmd.Value("subscription-id").(string),
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
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Subscriptions.UpdatePaymentMethod(
		ctx,
		cmd.Value("subscription-id").(string),
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
