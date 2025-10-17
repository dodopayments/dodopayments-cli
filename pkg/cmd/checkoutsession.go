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
		},
		&jsonflag.JSONIntFlag{
			Name: "product-cart.amount",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "product_cart.#.amount",
			},
		},
		&jsonflag.JSONAnyFlag{
			Name: "+product-cart",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "product_cart.-1",
				SetValue: map[string]interface{}{},
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "allowed-payment-method-types",
			Usage: "Customers will never see payment methods that are not in this list.\nHowever, adding a method here does not guarantee customers will see it.\nAvailability still depends on other factors (e.g., customer location, merchant settings).\n\nDisclaimar: Always provide 'credit' and 'debit' as a fallback.\nIf all payment methods are unavailable, checkout session will fail.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "allowed_payment_method_types.#",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "+allowed-payment-method-type",
			Usage: "Customers will never see payment methods that are not in this list.\nHowever, adding a method here does not guarantee customers will see it.\nAvailability still depends on other factors (e.g., customer location, merchant settings).\n\nDisclaimar: Always provide 'credit' and 'debit' as a fallback.\nIf all payment methods are unavailable, checkout session will fail.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "allowed_payment_method_types.-1",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "billing-address.country",
			Usage: "Two-letter ISO country code (ISO 3166-1 alpha-2)",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "billing_address.country",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "billing-address.city",
			Usage: "City name",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "billing_address.city",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "billing-address.state",
			Usage: "State or province name",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "billing_address.state",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "billing-address.street",
			Usage: "Street address including house number and unit/apartment if applicable",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "billing_address.street",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "billing-address.zipcode",
			Usage: "Postal code or ZIP code",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "billing_address.zipcode",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "billing-currency",
			Usage: "This field is ingored if adaptive pricing is disabled",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "billing_currency",
			},
		},
		&jsonflag.JSONBoolFlag{
			Name:  "confirm",
			Usage: "If confirm is true, all the details will be finalized. If required data is missing, an API error is thrown.",
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
			Name:  "customer.email",
			Usage: "Email is required for creating a new customer",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "customer.email",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "customer.name",
			Usage: "Optional full name of the customer. If provided during session creation,\nit is persisted and becomes immutable for the session. If omitted here,\nit can be provided later via the confirm API.",
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
			Name:  "customization.force_language",
			Usage: "Force the checkout interface to render in a specific language (e.g. `en`, `es`)",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "customization.force_language",
			},
		},
		&jsonflag.JSONBoolFlag{
			Name:  "customization.show_on_demand_tag",
			Usage: "Show on demand tag\n\nDefault is true",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "customization.show_on_demand_tag",
				SetValue: true,
			},
		},
		&jsonflag.JSONBoolFlag{
			Name:  "customization.show_order_details",
			Usage: "Show order details by default\n\nDefault is true",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "customization.show_order_details",
				SetValue: true,
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "customization.theme",
			Usage: "Theme of the page\n\nDefault is `System`.",
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
			Name:  "feature-flags.allow_currency_selection",
			Usage: "if customer is allowed to change currency, set it to true\n\nDefault is true",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "feature_flags.allow_currency_selection",
				SetValue: true,
			},
		},
		&jsonflag.JSONBoolFlag{
			Name:  "feature-flags.allow_discount_code",
			Usage: "If the customer is allowed to apply discount code, set it to true.\n\nDefault is true",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "feature_flags.allow_discount_code",
				SetValue: true,
			},
		},
		&jsonflag.JSONBoolFlag{
			Name:  "feature-flags.allow_phone_number_collection",
			Usage: "If phone number is collected from customer, set it to rue\n\nDefault is true",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "feature_flags.allow_phone_number_collection",
				SetValue: true,
			},
		},
		&jsonflag.JSONBoolFlag{
			Name:  "feature-flags.allow_tax_id",
			Usage: "If the customer is allowed to add tax id, set it to true\n\nDefault is true",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "feature_flags.allow_tax_id",
				SetValue: true,
			},
		},
		&jsonflag.JSONBoolFlag{
			Name:  "feature-flags.always_create_new_customer",
			Usage: "Set to true if a new customer object should be created.\nBy default email is used to find an existing customer to attach the session to\n\nDefault is false",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "feature_flags.always_create_new_customer",
				SetValue: true,
			},
		},
		&jsonflag.JSONBoolFlag{
			Name:  "force-3ds",
			Usage: "Override merchant default 3DS behaviour for this session",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "force_3ds",
				SetValue: true,
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "return-url",
			Usage: "The url to redirect after payment failure or success.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "return_url",
			},
		},
		&jsonflag.JSONBoolFlag{
			Name:  "show-saved-payment-methods",
			Usage: "Display saved payment methods of a returning customer False by default",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "show_saved_payment_methods",
				SetValue: true,
			},
		},
		&jsonflag.JSONBoolFlag{
			Name:  "subscription-data.on_demand.mandate_only",
			Usage: "If set as True, does not perform any charge and only authorizes payment method details for future use.",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "subscription_data.on_demand.mandate_only",
				SetValue: true,
			},
		},
		&jsonflag.JSONBoolFlag{
			Name:  "subscription-data.on_demand.adaptive_currency_fees_inclusive",
			Usage: "Whether adaptive currency fees should be included in the product_price (true) or added on top (false).\nThis field is ignored if adaptive pricing is not enabled for the business.",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "subscription_data.on_demand.adaptive_currency_fees_inclusive",
				SetValue: true,
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "subscription-data.on_demand.product_currency",
			Usage: "Optional currency of the product price. If not specified, defaults to the currency of the product.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "subscription_data.on_demand.product_currency",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "subscription-data.on_demand.product_description",
			Usage: "Optional product description override for billing and line items.\nIf not specified, the stored description of the product will be used.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "subscription_data.on_demand.product_description",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "subscription-data.on_demand.product_price",
			Usage: "Product price for the initial charge to customer\nIf not specified the stored price of the product will be used\nRepresented in the lowest denomination of the currency (e.g., cents for USD).\nFor example, to charge $1.00, pass `100`.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "subscription_data.on_demand.product_price",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "subscription-data.trial_period_days",
			Usage: "Optional trial period in days If specified, this value overrides the trial period set in the product's price Must be between 0 and 10000 days",
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
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.CheckoutSessionNewParams{}
	var res []byte
	_, err := cc.client.CheckoutSessions.New(
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
	return ShowJSON("checkout-sessions create", json, format, transform)
}
