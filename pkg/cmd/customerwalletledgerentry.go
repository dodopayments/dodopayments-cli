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

var customersWalletsLedgerEntriesCreate = cli.Command{
	Name:  "create",
	Usage: "Perform create operation",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "customer-id",
		},
		&jsonflag.JSONIntFlag{
			Name: "amount",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "amount",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "currency",
			Usage: "Currency of the wallet to adjust",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "currency",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "entry-type",
			Usage: "Type of ledger entry - credit or debit",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "entry_type",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "idempotency-key",
			Usage: "Optional idempotency key to prevent duplicate entries",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "idempotency_key",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "reason",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "reason",
			},
		},
	},
	Action:          handleCustomersWalletsLedgerEntriesCreate,
	HideHelpCommand: true,
}

var customersWalletsLedgerEntriesList = cli.Command{
	Name:  "list",
	Usage: "Perform list operation",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "customer-id",
		},
		&jsonflag.JSONStringFlag{
			Name:  "currency",
			Usage: "Optional currency filter",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "currency",
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
	Action:          handleCustomersWalletsLedgerEntriesList,
	HideHelpCommand: true,
}

func handleCustomersWalletsLedgerEntriesCreate(_ context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("customer-id") && len(unusedArgs) > 0 {
		cmd.Set("customer-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.CustomerWalletLedgerEntryNewParams{}
	var res []byte
	_, err := cc.client.Customers.Wallets.LedgerEntries.New(
		context.TODO(),
		cmd.Value("customer-id").(string),
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
	return ShowJSON("customers:wallets:ledger-entries create", json, format, transform)
}

func handleCustomersWalletsLedgerEntriesList(_ context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("customer-id") && len(unusedArgs) > 0 {
		cmd.Set("customer-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.CustomerWalletLedgerEntryListParams{}
	var res []byte
	_, err := cc.client.Customers.Wallets.LedgerEntries.List(
		context.TODO(),
		cmd.Value("customer-id").(string),
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
	return ShowJSON("customers:wallets:ledger-entries list", json, format, transform)
}
