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

var subscriptionsCreate = cli.Command{
	Name:  "create",
	Usage: "Perform create operation",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name: "billing.city",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "billing.city",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "billing.country",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "billing.country",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "billing.state",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "billing.state",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "billing.street",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "billing.street",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "billing.zipcode",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "billing.zipcode",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "customer.customer_id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "customer.customer_id",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "customer.email",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "customer.email",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "customer.name",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "customer.name",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "customer.phone_number",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "customer.phone_number",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "product-id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "product_id",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "quantity",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "quantity",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "addons.addon_id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "addons.#.addon_id",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "addons.quantity",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "addons.#.quantity",
			},
		},
		&jsonflag.JSONAnyFlag{
			Name: "+addon",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "addons.-1",
				SetValue: map[string]interface{}{},
			},
			Value: map[string]interface{}{},
		},
		&jsonflag.JSONStringFlag{
			Name: "allowed-payment-method-types",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "allowed_payment_method_types.#",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "+allowed_payment_method_type",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "allowed_payment_method_types.-1",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "billing-currency",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "billing_currency",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "discount-code",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "discount_code",
			},
		},
		&jsonflag.JSONBoolFlag{
			Name: "on-demand.mandate_only",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "on_demand.mandate_only",
				SetValue: true,
			},
		},
		&jsonflag.JSONBoolFlag{
			Name: "on-demand.adaptive_currency_fees_inclusive",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "on_demand.adaptive_currency_fees_inclusive",
				SetValue: true,
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "on-demand.product_currency",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "on_demand.product_currency",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "on-demand.product_description",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "on_demand.product_description",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "on-demand.product_price",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "on_demand.product_price",
			},
		},
		&jsonflag.JSONBoolFlag{
			Name: "payment-link",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "payment_link",
				SetValue: true,
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "return-url",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "return_url",
			},
		},
		&jsonflag.JSONBoolFlag{
			Name: "show-saved-payment-methods",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "show_saved_payment_methods",
				SetValue: true,
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "tax-id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "tax_id",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "trial-period-days",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "trial_period_days",
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
		&cli.StringFlag{
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
		&cli.StringFlag{
			Name: "subscription-id",
		},
		&jsonflag.JSONStringFlag{
			Name: "billing.city",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "billing.city",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "billing.country",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "billing.country",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "billing.state",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "billing.state",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "billing.street",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "billing.street",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "billing.zipcode",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "billing.zipcode",
			},
		},
		&jsonflag.JSONBoolFlag{
			Name: "cancel-at-next-billing-date",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "cancel_at_next_billing_date",
				SetValue: true,
			},
		},
		&jsonflag.JSONDatetimeFlag{
			Name: "disable-on-demand.next_billing_date",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "disable_on_demand.next_billing_date",
			},
		},
		&jsonflag.JSONDatetimeFlag{
			Name: "next-billing-date",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "next_billing_date",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "status",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "status",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "tax-id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "tax_id",
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
		&jsonflag.JSONStringFlag{
			Name: "brand-id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "brand_id",
			},
		},
		&jsonflag.JSONDatetimeFlag{
			Name: "created-at-gte",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "created_at_gte",
			},
		},
		&jsonflag.JSONDatetimeFlag{
			Name: "created-at-lte",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "created_at_lte",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "customer-id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "customer_id",
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
		&jsonflag.JSONStringFlag{
			Name: "status",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "status",
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
		&cli.StringFlag{
			Name: "subscription-id",
		},
		&jsonflag.JSONStringFlag{
			Name: "product-id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "product_id",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "proration-billing-mode",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "proration_billing_mode",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "quantity",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "quantity",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "addons.addon_id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "addons.#.addon_id",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "addons.quantity",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "addons.#.quantity",
			},
		},
		&jsonflag.JSONAnyFlag{
			Name: "+addon",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "addons.-1",
				SetValue: map[string]interface{}{},
			},
			Value: map[string]interface{}{},
		},
	},
	Action:          handleSubscriptionsChangePlan,
	HideHelpCommand: true,
}

var subscriptionsCharge = cli.Command{
	Name:  "charge",
	Usage: "Perform charge operation",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "subscription-id",
		},
		&jsonflag.JSONIntFlag{
			Name: "product-price",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "product_price",
			},
		},
		&jsonflag.JSONBoolFlag{
			Name: "adaptive-currency-fees-inclusive",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "adaptive_currency_fees_inclusive",
				SetValue: true,
			},
		},
		&jsonflag.JSONBoolFlag{
			Name: "customer-balance-config.allow_customer_credits_purchase",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "customer_balance_config.allow_customer_credits_purchase",
				SetValue: true,
			},
		},
		&jsonflag.JSONBoolFlag{
			Name: "customer-balance-config.allow_customer_credits_usage",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "customer_balance_config.allow_customer_credits_usage",
				SetValue: true,
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "product-currency",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "product_currency",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "product-description",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "product_description",
			},
		},
	},
	Action:          handleSubscriptionsCharge,
	HideHelpCommand: true,
}

var subscriptionsRetrieveUsageHistory = cli.Command{
	Name:  "retrieve-usage-history",
	Usage: "Get detailed usage history for a subscription that includes usage-based billing\n(metered components). This endpoint provides insights into customer usage\npatterns and billing calculations over time.",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "subscription-id",
		},
		&jsonflag.JSONDatetimeFlag{
			Name: "end-date",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "end_date",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "meter-id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "meter_id",
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
		&jsonflag.JSONDatetimeFlag{
			Name: "start-date",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "start_date",
			},
		},
	},
	Action:          handleSubscriptionsRetrieveUsageHistory,
	HideHelpCommand: true,
}

func handleSubscriptionsCreate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.SubscriptionNewParams{}
	var res []byte
	_, err := cc.client.Subscriptions.New(
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
	return ShowJSON("subscriptions create", json, format, transform)
}

func handleSubscriptionsRetrieve(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("subscription-id") && len(unusedArgs) > 0 {
		cmd.Set("subscription-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	var res []byte
	_, err := cc.client.Subscriptions.Get(
		context.TODO(),
		cmd.Value("subscription-id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("subscriptions retrieve", json, format, transform)
}

func handleSubscriptionsUpdate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("subscription-id") && len(unusedArgs) > 0 {
		cmd.Set("subscription-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.SubscriptionUpdateParams{}
	var res []byte
	_, err := cc.client.Subscriptions.Update(
		context.TODO(),
		cmd.Value("subscription-id").(string),
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
	return ShowJSON("subscriptions update", json, format, transform)
}

func handleSubscriptionsList(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.SubscriptionListParams{}
	var res []byte
	_, err := cc.client.Subscriptions.List(
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
	return ShowJSON("subscriptions list", json, format, transform)
}

func handleSubscriptionsChangePlan(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("subscription-id") && len(unusedArgs) > 0 {
		cmd.Set("subscription-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.SubscriptionChangePlanParams{}
	return cc.client.Subscriptions.ChangePlan(
		context.TODO(),
		cmd.Value("subscription-id").(string),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
	)
}

func handleSubscriptionsCharge(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("subscription-id") && len(unusedArgs) > 0 {
		cmd.Set("subscription-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.SubscriptionChargeParams{}
	var res []byte
	_, err := cc.client.Subscriptions.Charge(
		context.TODO(),
		cmd.Value("subscription-id").(string),
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
	return ShowJSON("subscriptions charge", json, format, transform)
}

func handleSubscriptionsRetrieveUsageHistory(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("subscription-id") && len(unusedArgs) > 0 {
		cmd.Set("subscription-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.SubscriptionGetUsageHistoryParams{}
	var res []byte
	_, err := cc.client.Subscriptions.GetUsageHistory(
		context.TODO(),
		cmd.Value("subscription-id").(string),
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
	return ShowJSON("subscriptions retrieve-usage-history", json, format, transform)
}
