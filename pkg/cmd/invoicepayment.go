// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

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
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("payment-id") && len(unusedArgs) > 0 {
		cmd.Set("payment-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	var res []byte
	_, err := cc.client.Invoices.Payments.Get(
		ctx,
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
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("refund-id") && len(unusedArgs) > 0 {
		cmd.Set("refund-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	var res []byte
	_, err := cc.client.Invoices.Payments.GetRefund(
		ctx,
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
