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
			Name:  "activations-limit",
			Usage: "The updated activation limit for the license key.\nUse `null` to remove the limit, or omit this field to leave it unchanged.",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "activations_limit",
			},
		},
		&jsonflag.JSONBoolFlag{
			Name:  "disabled",
			Usage: "Indicates whether the license key should be disabled.\nA value of `true` disables the key, while `false` enables it. Omit this field to leave it unchanged.",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "disabled",
				SetValue: true,
			},
		},
		&jsonflag.JSONDatetimeFlag{
			Name:  "expires-at",
			Usage: "The updated expiration timestamp for the license key in UTC.\nUse `null` to remove the expiration date, or omit this field to leave it unchanged.",
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
			Name:  "customer-id",
			Usage: "Filter by customer ID",
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
			Name:  "product-id",
			Usage: "Filter by product ID",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "product_id",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "status",
			Usage: "Filter by license key status",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "status",
			},
		},
	},
	Action:          handleLicenseKeysList,
	HideHelpCommand: true,
}

func handleLicenseKeysRetrieve(_ context.Context, cmd *cli.Command) error {
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
	_, err := cc.client.LicenseKeys.Get(
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
	return ShowJSON("license-keys retrieve", json, format, transform)
}

func handleLicenseKeysUpdate(_ context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.LicenseKeyUpdateParams{}
	var res []byte
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

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("license-keys update", json, format, transform)
}

func handleLicenseKeysList(_ context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.LicenseKeyListParams{}
	var res []byte
	_, err := cc.client.LicenseKeys.List(
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
	return ShowJSON("license-keys list", json, format, transform)
}
