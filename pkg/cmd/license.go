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

var licensesActivate = cli.Command{
	Name:  "activate",
	Usage: "Perform activate operation",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name: "license-key",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "license_key",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "name",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "name",
			},
		},
	},
	Action:          handleLicensesActivate,
	HideHelpCommand: true,
}

var licensesDeactivate = cli.Command{
	Name:  "deactivate",
	Usage: "Perform deactivate operation",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name: "license-key",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "license_key",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "license-key-instance-id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "license_key_instance_id",
			},
		},
	},
	Action:          handleLicensesDeactivate,
	HideHelpCommand: true,
}

var licensesValidate = cli.Command{
	Name:  "validate",
	Usage: "Perform validate operation",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name: "license-key",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "license_key",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "license-key-instance-id",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "license_key_instance_id",
			},
		},
	},
	Action:          handleLicensesValidate,
	HideHelpCommand: true,
}

func handleLicensesActivate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.LicenseActivateParams{}
	var res []byte
	_, err := cc.client.Licenses.Activate(
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
	return ShowJSON("licenses activate", json, format, transform)
}

func handleLicensesDeactivate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.LicenseDeactivateParams{}
	return cc.client.Licenses.Deactivate(
		context.TODO(),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
	)
}

func handleLicensesValidate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.LicenseValidateParams{}
	var res []byte
	_, err := cc.client.Licenses.Validate(
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
	return ShowJSON("licenses validate", json, format, transform)
}
