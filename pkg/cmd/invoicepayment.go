// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"

	"github.com/dodopayments/dodopayments-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var invoicesPaymentsRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Perform retrieve operation",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "payment-id",
		},
	},
	Action:          handleInvoicesPaymentsRetrieve,
	HideHelpCommand: true,
}

var invoicesPaymentsRetrieveRefund = cli.Command{
	Name:  "retrieve-refund",
	Usage: "Perform retrieve-refund operation",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "refund-id",
		},
	},
	Action:          handleInvoicesPaymentsRetrieveRefund,
	HideHelpCommand: true,
}

func handleInvoicesPaymentsRetrieve(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	var res []byte
	_, err := cc.client.Invoices.Payments.Get(
		context.TODO(),
		cmd.Value("payment-id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("invoices:payments retrieve", json, format, transform)
}

func handleInvoicesPaymentsRetrieveRefund(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	var res []byte
	_, err := cc.client.Invoices.Payments.GetRefund(
		context.TODO(),
		cmd.Value("refund-id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("invoices:payments retrieve-refund", json, format, transform)
}
