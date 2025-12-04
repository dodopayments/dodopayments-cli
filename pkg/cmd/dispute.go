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

var disputesRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Perform retrieve operation",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name: "dispute-id",
		},
	},
	Action:          handleDisputesRetrieve,
	HideHelpCommand: true,
}

var disputesList = cli.Command{
	Name:  "list",
	Usage: "Perform list operation",
	Flags: []cli.Flag{
		&requestflag.DateTimeFlag{
			Name:  "created-at-gte",
			Usage: "Get events after this created time",
			Config: requestflag.RequestConfig{
				QueryPath: "created_at_gte",
			},
		},
		&requestflag.DateTimeFlag{
			Name:  "created-at-lte",
			Usage: "Get events created before this time",
			Config: requestflag.RequestConfig{
				QueryPath: "created_at_lte",
			},
		},
		&requestflag.StringFlag{
			Name:  "customer-id",
			Usage: "Filter by customer_id",
			Config: requestflag.RequestConfig{
				QueryPath: "customer_id",
			},
		},
		&requestflag.StringFlag{
			Name:  "dispute-stage",
			Usage: "Filter by dispute stage",
			Config: requestflag.RequestConfig{
				QueryPath: "dispute_stage",
			},
		},
		&requestflag.StringFlag{
			Name:  "dispute-status",
			Usage: "Filter by dispute status",
			Config: requestflag.RequestConfig{
				QueryPath: "dispute_status",
			},
		},
		&requestflag.IntFlag{
			Name:  "page-number",
			Usage: "Page number default is 0",
			Config: requestflag.RequestConfig{
				QueryPath: "page_number",
			},
		},
		&requestflag.IntFlag{
			Name:  "page-size",
			Usage: "Page size default is 10 max is 100",
			Config: requestflag.RequestConfig{
				QueryPath: "page_size",
			},
		},
	},
	Action:          handleDisputesList,
	HideHelpCommand: true,
}

func handleDisputesRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("dispute-id") && len(unusedArgs) > 0 {
		cmd.Set("dispute-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
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
	_, err = client.Disputes.Get(
		ctx,
		requestflag.CommandRequestValue[string](cmd, "dispute-id"),
		options...,
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("disputes retrieve", json, format, transform)
}

func handleDisputesList(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.DisputeListParams{}

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
	_, err = client.Disputes.List(
		ctx,
		params,
		options...,
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("disputes list", json, format, transform)
}
