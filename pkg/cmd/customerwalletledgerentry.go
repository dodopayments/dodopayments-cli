// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"
	"os"

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
		&requestflag.Flag[string]{
			Name: "customer-id",
		},
		&requestflag.Flag[int64]{
			Name:     "amount",
			BodyPath: "amount",
		},
		&requestflag.Flag[string]{
			Name:     "currency",
			BodyPath: "currency",
		},
		&requestflag.Flag[string]{
			Name:     "entry-type",
			Usage:    "Type of ledger entry - credit or debit",
			BodyPath: "entry_type",
		},
		&requestflag.Flag[string]{
			Name:     "idempotency-key",
			Usage:    "Optional idempotency key to prevent duplicate entries",
			BodyPath: "idempotency_key",
		},
		&requestflag.Flag[string]{
			Name:     "reason",
			BodyPath: "reason",
		},
	},
	Action:          handleCustomersWalletsLedgerEntriesCreate,
	HideHelpCommand: true,
}

var customersWalletsLedgerEntriesList = cli.Command{
	Name:  "list",
	Usage: "Perform list operation",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name: "customer-id",
		},
		&requestflag.Flag[string]{
			Name:      "currency",
			QueryPath: "currency",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			QueryPath: "page_number",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			QueryPath: "page_size",
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
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Customers.Wallets.LedgerEntries.New(
		ctx,
		cmd.Value("customer-id").(string),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "customers:wallets:ledger-entries create", obj, format, transform)
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
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Customers.Wallets.LedgerEntries.List(
			ctx,
			cmd.Value("customer-id").(string),
			params,
			options...,
		)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "customers:wallets:ledger-entries list", obj, format, transform)
	} else {
		iter := client.Customers.Wallets.LedgerEntries.ListAutoPaging(
			ctx,
			cmd.Value("customer-id").(string),
			params,
			options...,
		)
		return streamOutput("customers:wallets:ledger-entries list", func(w *os.File) error {
			for iter.Next() {
				item := iter.Current()
				obj := gjson.Parse(item.JSON.RawJSON())
				if err := ShowJSON(w, "customers:wallets:ledger-entries list", obj, format, transform); err != nil {
					return err
				}
			}
			return iter.Err()
		})
	}
}
