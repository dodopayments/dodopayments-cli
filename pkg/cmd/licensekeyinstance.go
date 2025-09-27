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

var licenseKeyInstancesRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Perform retrieve operation",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "id",
		},
	},
	Action:          handleLicenseKeyInstancesRetrieve,
	HideHelpCommand: true,
}

var licenseKeyInstancesUpdate = cli.Command{
	Name:  "update",
	Usage: "Perform update operation",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "id",
		},
		&jsonflag.JSONStringFlag{
			Name: "name",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "name",
			},
		},
	},
	Action:          handleLicenseKeyInstancesUpdate,
	HideHelpCommand: true,
}

var licenseKeyInstancesList = cli.Command{
	Name:  "list",
	Usage: "Perform list operation",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name: "license-key-id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "license_key_id",
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
	},
	Action:          handleLicenseKeyInstancesList,
	HideHelpCommand: true,
}

func handleLicenseKeyInstancesRetrieve(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	var res []byte
	_, err := cc.client.LicenseKeyInstances.Get(
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
	return ShowJSON("license-key-instances retrieve", json, format, transform)
}

func handleLicenseKeyInstancesUpdate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := dodopayments.LicenseKeyInstanceUpdateParams{}
	var res []byte
	_, err := cc.client.LicenseKeyInstances.Update(
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
	return ShowJSON("license-key-instances update", json, format, transform)
}

func handleLicenseKeyInstancesList(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := dodopayments.LicenseKeyInstanceListParams{}
	var res []byte
	_, err := cc.client.LicenseKeyInstances.List(
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
	return ShowJSON("license-key-instances list", json, format, transform)
}
