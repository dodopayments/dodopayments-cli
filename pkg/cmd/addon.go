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
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.AddonNewParams{}
	var res []byte
	_, err := cc.client.Addons.New(
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
	return ShowJSON("addons create", json, format, transform)
}

func handleAddonsRetrieve(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	var res []byte
	_, err := cc.client.Addons.Get(
		context.TODO(),
		cmd.Value("id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("addons retrieve", json, format, transform)
}

func handleAddonsUpdate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.AddonUpdateParams{}
	var res []byte
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

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("addons update", json, format, transform)
}

func handleAddonsList(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.AddonListParams{}
	var res []byte
	_, err := cc.client.Addons.List(
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
	return ShowJSON("addons list", json, format, transform)
}

func handleAddonsUpdateImages(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	var res []byte
	_, err := cc.client.Addons.UpdateImages(
		context.TODO(),
		cmd.Value("id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("addons update-images", json, format, transform)
}
