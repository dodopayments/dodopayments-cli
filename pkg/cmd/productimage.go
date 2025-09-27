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

var productsImagesUpdate = cli.Command{
	Name:  "update",
	Usage: "Perform update operation",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "id",
		},
		&jsonflag.JSONBoolFlag{
			Name: "force-update",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Query,
				Path:     "force_update",
				SetValue: true,
			},
		},
	},
	Action:          handleProductsImagesUpdate,
	HideHelpCommand: true,
}

func handleProductsImagesUpdate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := dodopayments.ProductImageUpdateParams{}
	var res []byte
	_, err := cc.client.Products.Images.Update(
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
	return ShowJSON("products:images update", json, format, transform)
}
