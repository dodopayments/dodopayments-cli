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

var paymentsCreate = cli.Command{
	Name:  "create",
	Usage: "Perform create operation",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name:  "billing.city",
			Usage: "City name",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "billing.city",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "billing.country",
			Usage: "Two-letter ISO country code (ISO 3166-1 alpha-2)",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "billing.country",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "billing.state",
			Usage: "State or province name",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "billing.state",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "billing.street",
			Usage: "Street address including house number and unit/apartment if applicable",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "billing.street",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "billing.zipcode",
			Usage: "Postal code or ZIP code",
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
			Name:  "product-cart.product_id",
			Usage: "List of products in the cart. Must contain at least 1 and at most 100 items.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "product_cart.#.product_id",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "product-cart.quantity",
			Usage: "List of products in the cart. Must contain at least 1 and at most 100 items.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "product_cart.#.quantity",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "product-cart.amount",
			Usage: "List of products in the cart. Must contain at least 1 and at most 100 items.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "product_cart.#.amount",
			},
		},
		&jsonflag.JSONAnyFlag{
			Name:  "+product-cart",
			Usage: "List of products in the cart. Must contain at least 1 and at most 100 items.",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "product_cart.-1",
				SetValue: map[string]interface{}{},
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "allowed-payment-method-types",
			Usage: "List of payment methods allowed during checkout.\n\nCustomers will **never** see payment methods that are **not** in this list.\nHowever, adding a method here **does not guarantee** customers will see it.\nAvailability still depends on other factors (e.g., customer location, merchant settings).",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "allowed_payment_method_types.#",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "+allowed-payment-method-type",
			Usage: "List of payment methods allowed during checkout.\n\nCustomers will **never** see payment methods that are **not** in this list.\nHowever, adding a method here **does not guarantee** customers will see it.\nAvailability still depends on other factors (e.g., customer location, merchant settings).",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "allowed_payment_method_types.-1",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "billing-currency",
			Usage: "Fix the currency in which the end customer is billed.\nIf Dodo Payments cannot support that currency for this transaction, it will not proceed",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "billing_currency",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "discount-code",
			Usage: "Discount Code to apply to the transaction",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "discount_code",
			},
		},
		&jsonflag.JSONBoolFlag{
			Name:  "payment-link",
			Usage: "Whether to generate a payment link. Defaults to false if not specified.",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "payment_link",
				SetValue: true,
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "return-url",
			Usage: "Optional URL to redirect the customer after payment.\nMust be a valid URL if provided.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "return_url",
			},
		},
		&jsonflag.JSONBoolFlag{
			Name:  "show-saved-payment-methods",
			Usage: "Display saved payment methods of a returning customer\nFalse by default",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "show_saved_payment_methods",
				SetValue: true,
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "tax-id",
			Usage: "Tax ID in case the payment is B2B. If tax id validation fails the payment creation will fail",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "tax_id",
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
		&cli.StringFlag{
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
		&jsonflag.JSONStringFlag{
			Name:  "brand-id",
			Usage: "filter by Brand id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "brand_id",
			},
		},
		&jsonflag.JSONDatetimeFlag{
			Name:  "created-at-gte",
			Usage: "Get events after this created time",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "created_at_gte",
			},
		},
		&jsonflag.JSONDatetimeFlag{
			Name:  "created-at-lte",
			Usage: "Get events created before this time",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "created_at_lte",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "customer-id",
			Usage: "Filter by customer id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "customer_id",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "page-number",
			Usage: "Page number default is 0",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "page_number",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "page-size",
			Usage: "Page size default is 10 max is 100",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "page_size",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "status",
			Usage: "Filter by status",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "status",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "subscription-id",
			Usage: "Filter by subscription id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "subscription_id",
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
		&cli.StringFlag{
			Name: "payment-id",
		},
	},
	Action:          handlePaymentsRetrieveLineItems,
	HideHelpCommand: true,
}

func handlePaymentsCreate(_ context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.PaymentNewParams{}
	var res []byte
	_, err := cc.client.Payments.New(
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
	return ShowJSON("payments create", json, format, transform)
}

func handlePaymentsRetrieve(_ context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("payment-id") && len(unusedArgs) > 0 {
		cmd.Set("payment-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	var res []byte
	_, err := cc.client.Payments.Get(
		context.TODO(),
		cmd.Value("payment-id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("payments retrieve", json, format, transform)
}

func handlePaymentsList(_ context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.PaymentListParams{}
	var res []byte
	_, err := cc.client.Payments.List(
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
	return ShowJSON("payments list", json, format, transform)
}

func handlePaymentsRetrieveLineItems(_ context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("payment-id") && len(unusedArgs) > 0 {
		cmd.Set("payment-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	var res []byte
	_, err := cc.client.Payments.GetLineItems(
		context.TODO(),
		cmd.Value("payment-id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("payments retrieve-line-items", json, format, transform)
}
