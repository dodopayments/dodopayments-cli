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

var discountsCreate = cli.Command{
	Name:  "create",
	Usage: "POST /discounts If `code` is omitted or empty, a random 16-char uppercase code\nis generated.",
	Flags: []cli.Flag{
		&jsonflag.JSONIntFlag{
			Name:  "amount",
			Usage: "The discount amount.\n\n- If `discount_type` is **not** `percentage`, `amount` is in **USD cents**. For example, `100` means `$1.00`.\n  Only USD is allowed.\n- If `discount_type` **is** `percentage`, `amount` is in **basis points**. For example, `540` means `5.4%`.\n\nMust be at least 1.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "amount",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "type",
			Usage: "The discount type (e.g. `percentage`, `flat`, or `flat_per_unit`).",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "type",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "code",
			Usage: "Optionally supply a code (will be uppercased).\n- Must be at least 3 characters if provided.\n- If omitted, a random 16-character code is generated.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "code",
			},
		},
		&jsonflag.JSONDatetimeFlag{
			Name:  "expires-at",
			Usage: "When the discount expires, if ever.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "expires_at",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "name",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "name",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "restricted-to",
			Usage: "List of product IDs to restrict usage (if any).",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "restricted_to.#",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "+restricted-to",
			Usage: "List of product IDs to restrict usage (if any).",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "restricted_to.-1",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "subscription-cycles",
			Usage: "Number of subscription billing cycles this discount is valid for.\nIf not provided, the discount will be applied indefinitely to\nall recurring payments related to the subscription.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "subscription_cycles",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "usage-limit",
			Usage: "How many times this discount can be used (if any).\nMust be >= 1 if provided.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "usage_limit",
			},
		},
	},
	Action:          handleDiscountsCreate,
	HideHelpCommand: true,
}

var discountsRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "GET /discounts/{discount_id}",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "discount-id",
		},
	},
	Action:          handleDiscountsRetrieve,
	HideHelpCommand: true,
}

var discountsUpdate = cli.Command{
	Name:  "update",
	Usage: "PATCH /discounts/{discount_id}",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "discount-id",
		},
		&jsonflag.JSONIntFlag{
			Name:  "amount",
			Usage: "If present, update the discount amount:\n- If `discount_type` is `percentage`, this represents **basis points** (e.g., `540` = `5.4%`).\n- Otherwise, this represents **USD cents** (e.g., `100` = `$1.00`).\n\nMust be at least 1 if provided.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "amount",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "code",
			Usage: "If present, update the discount code (uppercase).",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "code",
			},
		},
		&jsonflag.JSONDatetimeFlag{
			Name: "expires-at",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "expires_at",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "name",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "name",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "restricted-to",
			Usage: "If present, replaces all restricted product IDs with this new set.\nTo remove all restrictions, send empty array",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "restricted_to.#",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "+restricted-to",
			Usage: "If present, replaces all restricted product IDs with this new set.\nTo remove all restrictions, send empty array",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "restricted_to.-1",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "subscription-cycles",
			Usage: "Number of subscription billing cycles this discount is valid for.\nIf not provided, the discount will be applied indefinitely to\nall recurring payments related to the subscription.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "subscription_cycles",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "type",
			Usage: "If present, update the discount type.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "type",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "usage-limit",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "usage_limit",
			},
		},
	},
	Action:          handleDiscountsUpdate,
	HideHelpCommand: true,
}

var discountsList = cli.Command{
	Name:  "list",
	Usage: "GET /discounts",
	Flags: []cli.Flag{
		&jsonflag.JSONIntFlag{
			Name:  "page-number",
			Usage: "Page number (default = 0).",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "page_number",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "page-size",
			Usage: "Page size (default = 10, max = 100).",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "page_size",
			},
		},
	},
	Action:          handleDiscountsList,
	HideHelpCommand: true,
}

var discountsDelete = cli.Command{
	Name:  "delete",
	Usage: "DELETE /discounts/{discount_id}",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "discount-id",
		},
	},
	Action:          handleDiscountsDelete,
	HideHelpCommand: true,
}

func handleDiscountsCreate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.DiscountNewParams{}
	var res []byte
	_, err := cc.client.Discounts.New(
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
	return ShowJSON("discounts create", json, format, transform)
}

func handleDiscountsRetrieve(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("discount-id") && len(unusedArgs) > 0 {
		cmd.Set("discount-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	var res []byte
	_, err := cc.client.Discounts.Get(
		ctx,
		cmd.Value("discount-id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("discounts retrieve", json, format, transform)
}

func handleDiscountsUpdate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("discount-id") && len(unusedArgs) > 0 {
		cmd.Set("discount-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.DiscountUpdateParams{}
	var res []byte
	_, err := cc.client.Discounts.Update(
		ctx,
		cmd.Value("discount-id").(string),
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
	return ShowJSON("discounts update", json, format, transform)
}

func handleDiscountsList(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.DiscountListParams{}
	var res []byte
	_, err := cc.client.Discounts.List(
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
	return ShowJSON("discounts list", json, format, transform)
}

func handleDiscountsDelete(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("discount-id") && len(unusedArgs) > 0 {
		cmd.Set("discount-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	return cc.client.Discounts.Delete(
		ctx,
		cmd.Value("discount-id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
	)
}
