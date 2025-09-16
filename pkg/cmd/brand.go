// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"

	"github.com/dodopayments/dodopayments-cli/pkg/jsonflag"
	"github.com/dodopayments/dodopayments-go"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/urfave/cli/v3"
)

var brandsCreate = cli.Command{
	Name:  "create",
	Usage: "Perform create operation",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name: "description",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "description",
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
			Name: "statement-descriptor",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "statement_descriptor",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "support-email",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "support_email",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "url",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "url",
			},
		},
	},
	Action:          handleBrandsCreate,
	HideHelpCommand: true,
}

var brandsRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Thin handler just calls `get_brand` and wraps in `Json(...)`",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "id",
		},
	},
	Action:          handleBrandsRetrieve,
	HideHelpCommand: true,
}

var brandsUpdate = cli.Command{
	Name:  "update",
	Usage: "Perform update operation",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "id",
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
		&jsonflag.JSONStringFlag{
			Name: "statement-descriptor",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "statement_descriptor",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "support-email",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "support_email",
			},
		},
	},
	Action:          handleBrandsUpdate,
	HideHelpCommand: true,
}

var brandsList = cli.Command{
	Name:            "list",
	Usage:           "Perform list operation",
	Flags:           []cli.Flag{},
	Action:          handleBrandsList,
	HideHelpCommand: true,
}

var brandsUpdateImages = cli.Command{
	Name:  "update-images",
	Usage: "Perform update-images operation",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "id",
		},
	},
	Action:          handleBrandsUpdateImages,
	HideHelpCommand: true,
}

func handleBrandsCreate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := dodopayments.BrandNewParams{}
	var res []byte
	_, err := cc.client.Brands.New(
		context.TODO(),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("brands create", string(res), format)
}

func handleBrandsRetrieve(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	var res []byte
	_, err := cc.client.Brands.Get(
		context.TODO(),
		cmd.Value("id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("brands retrieve", string(res), format)
}

func handleBrandsUpdate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := dodopayments.BrandUpdateParams{}
	var res []byte
	_, err := cc.client.Brands.Update(
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
	return ShowJSON("brands update", string(res), format)
}

func handleBrandsList(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	var res []byte
	_, err := cc.client.Brands.List(
		context.TODO(),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("brands list", string(res), format)
}

func handleBrandsUpdateImages(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	var res []byte
	_, err := cc.client.Brands.UpdateImages(
		context.TODO(),
		cmd.Value("id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("brands update-images", string(res), format)
}
