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

var addonsCreate = cli.Command{
	Name:  "create",
	Usage: "Perform create operation",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name: "currency",
			Config: requestflag.RequestConfig{
				BodyPath: "currency",
			},
		},
		&requestflag.StringFlag{
			Name:  "name",
			Usage: "Name of the Addon",
			Config: requestflag.RequestConfig{
				BodyPath: "name",
			},
		},
		&requestflag.IntFlag{
			Name:  "price",
			Usage: "Amount of the addon",
			Config: requestflag.RequestConfig{
				BodyPath: "price",
			},
		},
		&requestflag.StringFlag{
			Name:  "tax-category",
			Usage: "Represents the different categories of taxation applicable to various products and services.",
			Config: requestflag.RequestConfig{
				BodyPath: "tax_category",
			},
		},
		&requestflag.StringFlag{
			Name:  "description",
			Usage: "Optional description of the Addon",
			Config: requestflag.RequestConfig{
				BodyPath: "description",
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
		&requestflag.StringFlag{
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
		&requestflag.StringFlag{
			Name: "id",
		},
		&requestflag.StringFlag{
			Name: "currency",
			Config: requestflag.RequestConfig{
				BodyPath: "currency",
			},
		},
		&requestflag.StringFlag{
			Name:  "description",
			Usage: "Description of the Addon, optional and must be at most 1000 characters.",
			Config: requestflag.RequestConfig{
				BodyPath: "description",
			},
		},
		&requestflag.StringFlag{
			Name:  "image-id",
			Usage: "Addon image id after its uploaded to S3",
			Config: requestflag.RequestConfig{
				BodyPath: "image_id",
			},
		},
		&requestflag.StringFlag{
			Name:  "name",
			Usage: "Name of the Addon, optional and must be at most 100 characters.",
			Config: requestflag.RequestConfig{
				BodyPath: "name",
			},
		},
		&requestflag.IntFlag{
			Name:  "price",
			Usage: "Amount of the addon",
			Config: requestflag.RequestConfig{
				BodyPath: "price",
			},
		},
		&requestflag.StringFlag{
			Name:  "tax-category",
			Usage: "Represents the different categories of taxation applicable to various products and services.",
			Config: requestflag.RequestConfig{
				BodyPath: "tax_category",
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
		&requestflag.IntFlag{
			Name:  "page-number",
			Usage: "Page number default is 0",
			Config: requestflag.RequestConfig{
				QueryPath: "page_number",
			},
		},
		&requestflag.IntFlag{
			Name:  "page-size",
			Usage: "Page size default is 10 max is 100",
			Config: requestflag.RequestConfig{
				QueryPath: "page_size",
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
		&requestflag.StringFlag{
			Name: "id",
		},
	},
	Action:          handleAddonsUpdateImages,
	HideHelpCommand: true,
}

func handleAddonsCreate(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.AddonNewParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Addons.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "addons create", obj, format, transform)
}

func handleAddonsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
		ApplicationJSON,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Addons.Get(ctx, requestflag.CommandRequestValue[string](cmd, "id"), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "addons retrieve", obj, format, transform)
}

func handleAddonsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.AddonUpdateParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Addons.Update(
		ctx,
		requestflag.CommandRequestValue[string](cmd, "id"),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "addons update", obj, format, transform)
}

func handleAddonsList(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.AddonListParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Addons.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "addons list", obj, format, transform)
	} else {
		iter := client.Addons.ListAutoPaging(ctx, params, options...)
		return streamOutput("addons list", func(w *os.File) error {
			for iter.Next() {
				item := iter.Current()
				obj := gjson.Parse(item.JSON.RawJSON())
				if err := ShowJSON(w, "addons list", obj, format, transform); err != nil {
					return err
				}
			}
			return iter.Err()
		})
	}
}

func handleAddonsUpdateImages(ctx context.Context, cmd *cli.Command) error {
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
		ApplicationJSON,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Addons.UpdateImages(ctx, requestflag.CommandRequestValue[string](cmd, "id"), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "addons update-images", obj, format, transform)
}
