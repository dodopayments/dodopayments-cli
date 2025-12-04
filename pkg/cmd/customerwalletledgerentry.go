// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/dodopayments/dodopayments-cli/internal/apiquery"
	"github.com/dodopayments/dodopayments-cli/internal/requestflag"
	"github.com/dodopayments/dodopayments-go"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var customersWalletsLedgerEntriesCreate = cli.Command{
	Name:  "create",
	Usage: "Perform create operation",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name: "customer-id",
		},
		&requestflag.IntFlag{
			Name: "amount",
			Config: requestflag.RequestConfig{
				BodyPath: "amount",
			},
		},
		&requestflag.StringFlag{
			Name: "currency",
			Config: requestflag.RequestConfig{
				BodyPath: "currency",
			},
		},
		&requestflag.StringFlag{
			Name:  "entry-type",
			Usage: "Type of ledger entry - credit or debit",
			Config: requestflag.RequestConfig{
				BodyPath: "entry_type",
			},
		},
		&requestflag.StringFlag{
			Name:  "idempotency-key",
			Usage: "Optional idempotency key to prevent duplicate entries",
			Config: requestflag.RequestConfig{
				BodyPath: "idempotency_key",
			},
		},
		&requestflag.StringFlag{
			Name: "reason",
			Config: requestflag.RequestConfig{
				BodyPath: "reason",
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
		&requestflag.StringFlag{
			Name: "customer-id",
		},
		&requestflag.StringFlag{
			Name: "currency",
			Config: requestflag.RequestConfig{
				QueryPath: "currency",
			},
		},
		&requestflag.IntFlag{
			Name: "page-number",
			Config: requestflag.RequestConfig{
				QueryPath: "page_number",
			},
		},
		&requestflag.IntFlag{
			Name: "page-size",
			Config: requestflag.RequestConfig{
				QueryPath: "page_size",
			},
		},
	},
	Action:          handleCustomersWalletsLedgerEntriesList,
	HideHelpCommand: true,
}

func handleCustomersWalletsLedgerEntriesCreate(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("customer-id") && len(unusedArgs) > 0 {
		cmd.Set("customer-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.CustomerWalletLedgerEntryNewParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}
	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Customers.Wallets.LedgerEntries.New(
		ctx,
		requestflag.CommandRequestValue[string](cmd, "customer-id"),
		params,
		options...,
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
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("customer-id") && len(unusedArgs) > 0 {
		cmd.Set("customer-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.CustomerWalletLedgerEntryListParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}
	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Customers.Wallets.LedgerEntries.List(
		ctx,
		requestflag.CommandRequestValue[string](cmd, "customer-id"),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("customers:wallets:ledger-entries list", json, format, transform)
}
