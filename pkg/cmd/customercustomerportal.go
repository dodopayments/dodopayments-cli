// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"

	"github.com/dodopayments/dodopayments-cli/pkg/jsonflag"
	"github.com/dodopayments/dodopayments-go"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/urfave/cli/v3"
)

var customersCustomerPortalCreate = cli.Command{
	Name:  "create",
	Usage: "Perform create operation",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "customer-id",
		},
		&jsonflag.JSONBoolFlag{
			Name: "send-email",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Query,
				Path:     "send_email",
				SetValue: true,
			},
		},
	},
	Action:          handleCustomersCustomerPortalCreate,
	HideHelpCommand: true,
}

func handleCustomersCustomerPortalCreate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := dodopayments.CustomerCustomerPortalNewParams{}
	var res []byte
	_, err := cc.client.Customers.CustomerPortal.New(
		context.TODO(),
		cmd.Value("customer-id").(string),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("customers:customer-portal create", string(res), format)
}
