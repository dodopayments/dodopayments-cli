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

var disputesRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Perform retrieve operation",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
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
		&requestflag.Flag[requestflag.DateTimeValue]{
			Name:      "created-at-gte",
			Usage:     "Get events after this created time",
			QueryPath: "created_at_gte",
		},
		&requestflag.Flag[requestflag.DateTimeValue]{
			Name:      "created-at-lte",
			Usage:     "Get events created before this time",
			QueryPath: "created_at_lte",
		},
		&requestflag.Flag[string]{
			Name:      "customer-id",
			Usage:     "Filter by customer_id",
			QueryPath: "customer_id",
		},
		&requestflag.Flag[string]{
			Name:      "dispute-stage",
			Usage:     "Filter by dispute stage",
			QueryPath: "dispute_stage",
		},
		&requestflag.Flag[string]{
			Name:      "dispute-status",
			Usage:     "Filter by dispute status",
			QueryPath: "dispute_status",
		},
		&requestflag.Flag[int64]{
			Name:      "page-number",
			Usage:     "Page number default is 0",
			QueryPath: "page_number",
		},
		&requestflag.Flag[int64]{
			Name:      "page-size",
			Usage:     "Page size default is 10 max is 100",
			QueryPath: "page_size",
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
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Disputes.Get(ctx, cmd.Value("dispute-id").(string), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "disputes retrieve", obj, format, transform)
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
		_, err = client.Disputes.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "disputes list", obj, format, transform)
	} else {
		iter := client.Disputes.ListAutoPaging(ctx, params, options...)
		return streamOutput("disputes list", func(w *os.File) error {
			for iter.Next() {
				item := iter.Current()
				obj := gjson.Parse(item.RawJSON())
				if err := ShowJSON(w, "disputes list", obj, format, transform); err != nil {
					return err
				}
			}
			return iter.Err()
		})
	}
}
