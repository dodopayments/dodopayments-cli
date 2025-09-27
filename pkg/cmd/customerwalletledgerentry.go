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
			Name: "currency",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "currency",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "entry-type",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "entry_type",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "idempotency-key",
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
			Name: "currency",
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

func handleCustomersWalletsLedgerEntriesCreate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
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

func handleCustomersWalletsLedgerEntriesList(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
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
