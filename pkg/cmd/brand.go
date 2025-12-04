// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/dodopayments/dodopayments-cli/internal/apiquery"
	"github.com/dodopayments/dodopayments-cli/internal/requestflag"
	"github.com/dodopayments/dodopayments-go"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var brandsCreate = cli.Command{
	Name:  "create",
	Usage: "Perform create operation",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name: "description",
			Config: requestflag.RequestConfig{
				BodyPath: "description",
			},
		},
		&requestflag.StringFlag{
			Name: "name",
			Config: requestflag.RequestConfig{
				BodyPath: "name",
			},
		},
		&requestflag.StringFlag{
			Name: "statement-descriptor",
			Config: requestflag.RequestConfig{
				BodyPath: "statement_descriptor",
			},
		},
		&requestflag.StringFlag{
			Name: "support-email",
			Config: requestflag.RequestConfig{
				BodyPath: "support_email",
			},
		},
		&requestflag.StringFlag{
			Name: "url",
			Config: requestflag.RequestConfig{
				BodyPath: "url",
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
		&requestflag.StringFlag{
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
		&requestflag.StringFlag{
			Name: "id",
		},
		&requestflag.StringFlag{
			Name:  "image-id",
			Usage: "The UUID you got back from the presigned‐upload call",
			Config: requestflag.RequestConfig{
				BodyPath: "image_id",
			},
		},
		&requestflag.StringFlag{
			Name: "name",
			Config: requestflag.RequestConfig{
				BodyPath: "name",
			},
		},
		&requestflag.StringFlag{
			Name: "statement-descriptor",
			Config: requestflag.RequestConfig{
				BodyPath: "statement_descriptor",
			},
		},
		&requestflag.StringFlag{
			Name: "support-email",
			Config: requestflag.RequestConfig{
				BodyPath: "support_email",
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
		&requestflag.StringFlag{
			Name: "id",
		},
	},
	Action:          handleBrandsUpdateImages,
	HideHelpCommand: true,
}

func handleBrandsCreate(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.BrandNewParams{}

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
	_, err = client.Brands.New(
		ctx,
		params,
		options...,
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("brands create", json, format, transform)
}

func handleBrandsRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Brands.Get(
		ctx,
		requestflag.CommandRequestValue[string](cmd, "id"),
		options...,
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("brands retrieve", json, format, transform)
}

func handleBrandsUpdate(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.BrandUpdateParams{}

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
	_, err = client.Brands.Update(
		ctx,
		requestflag.CommandRequestValue[string](cmd, "id"),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("brands update", json, format, transform)
}

func handleBrandsList(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
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
	_, err = client.Brands.List(ctx, options...)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("brands list", json, format, transform)
}

func handleBrandsUpdateImages(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Brands.UpdateImages(
		ctx,
		requestflag.CommandRequestValue[string](cmd, "id"),
		options...,
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("brands update-images", json, format, transform)
}
