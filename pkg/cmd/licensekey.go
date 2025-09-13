// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"

	"github.com/dodopayments/dodopayments-cli/pkg/jsonflag"
	"github.com/dodopayments/dodopayments-go"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/urfave/cli/v3"
)

var licenseKeysRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Perform retrieve operation",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "id",
		},
	},
	Action:          handleLicenseKeysRetrieve,
	HideHelpCommand: true,
}

var licenseKeysUpdate = cli.Command{
	Name:  "update",
	Usage: "Perform update operation",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "id",
		},
		&jsonflag.JSONIntFlag{
			Name: "activations-limit",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "activations_limit",
			},
		},
		&jsonflag.JSONBoolFlag{
			Name: "disabled",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "disabled",
				SetValue: true,
			},
		},
		&jsonflag.JSONDatetimeFlag{
			Name: "expires-at",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "expires_at",
			},
		},
	},
	Action:          handleLicenseKeysUpdate,
	HideHelpCommand: true,
}

var licenseKeysList = cli.Command{
	Name:  "list",
	Usage: "Perform list operation",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name: "customer-id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "customer_id",
			},
		},
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
		&jsonflag.JSONStringFlag{
			Name: "product-id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "product_id",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "status",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "status",
			},
		},
	},
	Action:          handleLicenseKeysList,
	HideHelpCommand: true,
}

func handleLicenseKeysRetrieve(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	res := []byte{}
	_, err := cc.client.LicenseKeys.Get(
		context.TODO(),
		cmd.Value("id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("license-keys retrieve", string(res), format)
}

func handleLicenseKeysUpdate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := dodopayments.LicenseKeyUpdateParams{}
	res := []byte{}
	_, err := cc.client.LicenseKeys.Update(
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
	return ShowJSON("license-keys update", string(res), format)
}

func handleLicenseKeysList(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := dodopayments.LicenseKeyListParams{}
	res := []byte{}
	_, err := cc.client.LicenseKeys.List(
		context.TODO(),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("license-keys list", string(res), format)
}
