// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"

	"github.com/dodopayments/dodopayments-go"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/stainless-sdks/dodo-payments-cli/pkg/jsonflag"
	"github.com/urfave/cli/v3"
)

var addonsCreate = cli.Command{
	Name:  "create",
	Usage: "Perform create operation",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name: "currency",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "currency",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "name",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "name",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "price",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "tax-category",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "tax_category",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "description",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "description",
			},
		},
	},
	Action:          handleAddonsCreate,
	HideHelpCommand: true,
}

var addonsRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Perform retrieve operation",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "id",
		},
	},
	Action:          handleAddonsRetrieve,
	HideHelpCommand: true,
}

var addonsUpdate = cli.Command{
	Name:  "update",
	Usage: "Perform update operation",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "id",
		},
		&jsonflag.JSONStringFlag{
			Name: "currency",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "currency",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "description",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "description",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "image-id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "image_id",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "name",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "name",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "price",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "price",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "tax-category",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "tax_category",
			},
		},
	},
	Action:          handleAddonsUpdate,
	HideHelpCommand: true,
}

var addonsList = cli.Command{
	Name:  "list",
	Usage: "Perform list operation",
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
	Action:          handleAddonsList,
	HideHelpCommand: true,
}

var addonsUpdateImages = cli.Command{
	Name:  "update-images",
	Usage: "Perform update-images operation",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "id",
		},
	},
	Action:          handleAddonsUpdateImages,
	HideHelpCommand: true,
}

func handleAddonsCreate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := dodopayments.AddonNewParams{}
	res := []byte{}
	_, err := cc.client.Addons.New(
		context.TODO(),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("addons create", string(res), format)
}

func handleAddonsRetrieve(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	res := []byte{}
	_, err := cc.client.Addons.Get(
		context.TODO(),
		cmd.Value("id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("addons retrieve", string(res), format)
}

func handleAddonsUpdate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := dodopayments.AddonUpdateParams{}
	res := []byte{}
	_, err := cc.client.Addons.Update(
		context.TODO(),
		cmd.Value("id").(string),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("addons update", string(res), format)
}

func handleAddonsList(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := dodopayments.AddonListParams{}
	res := []byte{}
	_, err := cc.client.Addons.List(
		context.TODO(),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("addons list", string(res), format)
}

func handleAddonsUpdateImages(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	res := []byte{}
	_, err := cc.client.Addons.UpdateImages(
		context.TODO(),
		cmd.Value("id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("addons update-images", string(res), format)
}
