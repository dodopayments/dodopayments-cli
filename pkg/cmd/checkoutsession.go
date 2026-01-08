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

var checkoutSessionsCreate = requestflag.WithInnerFlags(cli.Command{
	Name:  "create",
	Usage: "Perform create operation",
	Flags: []cli.Flag{
		&requestflag.Flag[[]map[string]any]{
			Name:     "product-cart",
			Required: true,
			BodyPath: "product_cart",
		},
		&requestflag.Flag[[]string]{
			Name:     "allowed-payment-method-type",
			Usage:    "Customers will never see payment methods that are not in this list.\nHowever, adding a method here does not guarantee customers will see it.\nAvailability still depends on other factors (e.g., customer location, merchant settings).\n\nDisclaimar: Always provide 'credit' and 'debit' as a fallback.\nIf all payment methods are unavailable, checkout session will fail.",
			BodyPath: "allowed_payment_method_types",
		},
		&requestflag.Flag[map[string]any]{
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
		&requestflag.Flag[map[string]string]{
			Name:     "customer",
			BodyPath: "customer",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "customization",
			Usage:    "Customization for the checkout session page",
			BodyPath: "customization",
		},
		&requestflag.Flag[string]{
			Name:     "discount-code",
			BodyPath: "discount_code",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "feature-flags",
			BodyPath: "feature_flags",
		},
		&requestflag.Flag[bool]{
			Name:     "force-3ds",
			Usage:    "Override merchant default 3DS behaviour for this session",
			BodyPath: "force_3ds",
		},
		&requestflag.Flag[map[string]any]{
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
			Name:     "payment-method-id",
			Usage:    "Optional payment method ID to use for this checkout session.\nOnly allowed when `confirm` is true.\nIf provided, existing customer id must also be provided.",
			BodyPath: "payment_method_id",
		},
		&requestflag.Flag[string]{
			Name:     "return-url",
			Usage:    "The url to redirect after payment failure or success.",
			BodyPath: "return_url",
		},
		&requestflag.Flag[bool]{
			Name:     "short-link",
			Usage:    "If true, returns a shortened checkout URL.\nDefaults to false if not specified.",
			BodyPath: "short_link",
		},
		&requestflag.Flag[bool]{
			Name:     "show-saved-payment-methods",
			Usage:    "Display saved payment methods of a returning customer False by default",
			BodyPath: "show_saved_payment_methods",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "subscription-data",
			BodyPath: "subscription_data",
		},
	},
	Action:          handleCheckoutSessionsCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"product-cart": {
		&requestflag.InnerFlag[string]{
			Name:       "product-cart.product-id",
			Usage:      "unique id of the product",
			InnerField: "product_id",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "product-cart.quantity",
			InnerField: "quantity",
		},
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "product-cart.addons",
			Usage:      "only valid if product is a subscription",
			InnerField: "addons",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "product-cart.amount",
			Usage:      "Amount the customer pays if pay_what_you_want is enabled. If disabled then amount will be ignored\nRepresented in the lowest denomination of the currency (e.g., cents for USD).\nFor example, to charge $1.00, pass `100`.\nOnly applicable for one time payments\n\nIf amount is not set for pay_what_you_want product,\ncustomer is allowed to select the amount.",
			InnerField: "amount",
		},
	},
	"billing-address": {
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.country",
			Usage:      "ISO country code alpha2 variant",
			InnerField: "country",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.city",
			Usage:      "City name",
			InnerField: "city",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.state",
			Usage:      "State or province name",
			InnerField: "state",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.street",
			Usage:      "Street address including house number and unit/apartment if applicable",
			InnerField: "street",
		},
		&requestflag.InnerFlag[string]{
			Name:       "billing-address.zipcode",
			Usage:      "Postal code or ZIP code",
			InnerField: "zipcode",
		},
	},
	"customization": {
		&requestflag.InnerFlag[string]{
			Name:       "customization.force-language",
			Usage:      "Force the checkout interface to render in a specific language (e.g. `en`, `es`)",
			InnerField: "force_language",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "customization.show-on-demand-tag",
			Usage:      "Show on demand tag\n\nDefault is true",
			InnerField: "show_on_demand_tag",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "customization.show-order-details",
			Usage:      "Show order details by default\n\nDefault is true",
			InnerField: "show_order_details",
		},
		&requestflag.InnerFlag[string]{
			Name:       "customization.theme",
			Usage:      "Theme of the page\n\nDefault is `System`.",
			InnerField: "theme",
		},
	},
	"feature-flags": {
		&requestflag.InnerFlag[bool]{
			Name:       "feature-flags.allow-currency-selection",
			Usage:      "if customer is allowed to change currency, set it to true\n\nDefault is true",
			InnerField: "allow_currency_selection",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "feature-flags.allow-customer-editing-city",
			InnerField: "allow_customer_editing_city",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "feature-flags.allow-customer-editing-country",
			InnerField: "allow_customer_editing_country",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "feature-flags.allow-customer-editing-email",
			InnerField: "allow_customer_editing_email",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "feature-flags.allow-customer-editing-name",
			InnerField: "allow_customer_editing_name",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "feature-flags.allow-customer-editing-state",
			InnerField: "allow_customer_editing_state",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "feature-flags.allow-customer-editing-street",
			InnerField: "allow_customer_editing_street",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "feature-flags.allow-customer-editing-zipcode",
			InnerField: "allow_customer_editing_zipcode",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "feature-flags.allow-discount-code",
			Usage:      "If the customer is allowed to apply discount code, set it to true.\n\nDefault is true",
			InnerField: "allow_discount_code",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "feature-flags.allow-phone-number-collection",
			Usage:      "If phone number is collected from customer, set it to rue\n\nDefault is true",
			InnerField: "allow_phone_number_collection",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "feature-flags.allow-tax-id",
			Usage:      "If the customer is allowed to add tax id, set it to true\n\nDefault is true",
			InnerField: "allow_tax_id",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "feature-flags.always-create-new-customer",
			Usage:      "Set to true if a new customer object should be created.\nBy default email is used to find an existing customer to attach the session to\n\nDefault is false",
			InnerField: "always_create_new_customer",
		},
		&requestflag.InnerFlag[bool]{
			Name:       "feature-flags.redirect-immediately",
			Usage:      "If true, redirects the customer immediately after payment completion\n\nDefault is false",
			InnerField: "redirect_immediately",
		},
	},
	"subscription-data": {
		&requestflag.InnerFlag[map[string]any]{
			Name:       "subscription-data.on-demand",
			InnerField: "on_demand",
		},
		&requestflag.InnerFlag[int64]{
			Name:       "subscription-data.trial-period-days",
			Usage:      "Optional trial period in days If specified, this value overrides the trial period set in the product's price Must be between 0 and 10000 days",
			InnerField: "trial_period_days",
		},
	},
})

var checkoutSessionsRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Perform retrieve operation",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
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
		false,
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
		false,
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
