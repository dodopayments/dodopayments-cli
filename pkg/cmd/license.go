// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"

	"github.com/dodopayments/dodopayments-cli/pkg/jsonflag"
	"github.com/dodopayments/dodopayments-go"
	"github.com/dodopayments/dodopayments-go/option"
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
	params := dodopayments.LicenseActivateParams{}
	res := []byte{}
	_, err := cc.client.Licenses.Activate(
		context.TODO(),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("licenses activate", string(res), format)
}

func handleLicensesDeactivate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := dodopayments.LicenseDeactivateParams{}
	return cc.client.Licenses.Deactivate(
		context.TODO(),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
	)
}

func handleLicensesValidate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := dodopayments.LicenseValidateParams{}
	res := []byte{}
	_, err := cc.client.Licenses.Validate(
		context.TODO(),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("licenses validate", string(res), format)
}
