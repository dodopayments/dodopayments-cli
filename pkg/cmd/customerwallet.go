// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"

	"github.com/dodopayments/dodopayments-go/option"
	"github.com/urfave/cli/v3"
)

var customersWalletsList = cli.Command{
	Name:  "list",
	Usage: "Perform list operation",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "customer-id",
		},
	},
	Action:          handleCustomersWalletsList,
	HideHelpCommand: true,
}

func handleCustomersWalletsList(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	var res []byte
	_, err := cc.client.Customers.Wallets.List(
		context.TODO(),
		cmd.Value("customer-id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("customers:wallets list", string(res), format)
}
