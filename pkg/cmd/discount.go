// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"

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
			Name: "amount",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "amount",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "type",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "type",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "code",
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
			Name: "restricted-to",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "restricted_to.#",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "+restricted_to",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "restricted_to.-1",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "subscription-cycles",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "subscription_cycles",
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
			Name: "amount",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "amount",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "code",
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
			Name: "restricted-to",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "restricted_to.#",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "+restricted_to",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "restricted_to.-1",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "subscription-cycles",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "subscription_cycles",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "type",
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
	params := dodopayments.DiscountNewParams{}
	var res []byte
	_, err := cc.client.Discounts.New(
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
	return ShowJSON("discounts create", json, format, transform)
}

func handleDiscountsRetrieve(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	var res []byte
	_, err := cc.client.Discounts.Get(
		context.TODO(),
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
	params := dodopayments.DiscountUpdateParams{}
	var res []byte
	_, err := cc.client.Discounts.Update(
		context.TODO(),
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
	params := dodopayments.DiscountListParams{}
	var res []byte
	_, err := cc.client.Discounts.List(
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
	return ShowJSON("discounts list", json, format, transform)
}

func handleDiscountsDelete(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	return cc.client.Discounts.Delete(
		context.TODO(),
		cmd.Value("discount-id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
	)
}
