// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"

	"github.com/dodopayments/dodopayments-go"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/stainless-sdks/dodo-payments-cli/pkg/jsonflag"
	"github.com/urfave/cli/v3"
)

var checkoutSessionsCreate = cli.Command{
	Name:  "create",
	Usage: "Perform create operation",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name: "product-cart.product_id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "product_cart.#.product_id",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "product-cart.quantity",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "product_cart.#.quantity",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "product-cart.addons.addon_id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "product_cart.#.addons.#.addon_id",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "product-cart.addons.quantity",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "product_cart.#.addons.#.quantity",
			},
		},
		&jsonflag.JSONAnyFlag{
			Name: "product-cart.+addon",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "product_cart.#.addons.-1",
				SetValue: map[string]interface{}{},
			},
			Value: map[string]interface{}{},
		},
		&jsonflag.JSONIntFlag{
			Name: "product-cart.amount",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "product_cart.#.amount",
			},
		},
		&jsonflag.JSONAnyFlag{
			Name: "+product_cart",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "product_cart.-1",
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
			Name: "billing-address.country",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "billing_address.country",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "billing-address.city",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "billing_address.city",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "billing-address.state",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "billing_address.state",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "billing-address.street",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "billing_address.street",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "billing-address.zipcode",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "billing_address.zipcode",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "billing-currency",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "billing_currency",
			},
		},
		&jsonflag.JSONBoolFlag{
			Name: "confirm",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "confirm",
				SetValue: true,
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
		&jsonflag.JSONBoolFlag{
			Name: "customization.show_on_demand_tag",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "customization.show_on_demand_tag",
				SetValue: true,
			},
		},
		&jsonflag.JSONBoolFlag{
			Name: "customization.show_order_details",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "customization.show_order_details",
				SetValue: true,
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "customization.theme",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "customization.theme",
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
			Name: "feature-flags.allow_currency_selection",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "feature_flags.allow_currency_selection",
				SetValue: true,
			},
		},
		&jsonflag.JSONBoolFlag{
			Name: "feature-flags.allow_discount_code",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "feature_flags.allow_discount_code",
				SetValue: true,
			},
		},
		&jsonflag.JSONBoolFlag{
			Name: "feature-flags.allow_phone_number_collection",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "feature_flags.allow_phone_number_collection",
				SetValue: true,
			},
		},
		&jsonflag.JSONBoolFlag{
			Name: "feature-flags.allow_tax_id",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "feature_flags.allow_tax_id",
				SetValue: true,
			},
		},
		&jsonflag.JSONBoolFlag{
			Name: "feature-flags.always_create_new_customer",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "feature_flags.always_create_new_customer",
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
		&jsonflag.JSONBoolFlag{
			Name: "subscription-data.on_demand.mandate_only",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "subscription_data.on_demand.mandate_only",
				SetValue: true,
			},
		},
		&jsonflag.JSONBoolFlag{
			Name: "subscription-data.on_demand.adaptive_currency_fees_inclusive",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "subscription_data.on_demand.adaptive_currency_fees_inclusive",
				SetValue: true,
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "subscription-data.on_demand.product_currency",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "subscription_data.on_demand.product_currency",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "subscription-data.on_demand.product_description",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "subscription_data.on_demand.product_description",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "subscription-data.on_demand.product_price",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "subscription_data.on_demand.product_price",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "subscription-data.trial_period_days",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "subscription_data.trial_period_days",
			},
		},
	},
	Action:          handleCheckoutSessionsCreate,
	HideHelpCommand: true,
}

func handleCheckoutSessionsCreate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := dodopayments.CheckoutSessionNewParams{}
	res := []byte{}
	_, err := cc.client.CheckoutSessions.New(
		context.TODO(),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("checkout-sessions create", string(res), format)
}
