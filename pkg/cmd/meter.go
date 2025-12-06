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

var metersCreate = cli.Command{
	Name:  "create",
	Usage: "Perform create operation",
	Flags: []cli.Flag{
		&requestflag.YAMLFlag{
			Name: "aggregation",
			Config: requestflag.RequestConfig{
				BodyPath: "aggregation",
			},
		},
		&requestflag.StringFlag{
			Name:  "event-name",
			Usage: "Event name to track",
			Config: requestflag.RequestConfig{
				BodyPath: "event_name",
			},
		},
		&requestflag.StringFlag{
			Name:  "measurement-unit",
			Usage: "measurement unit",
			Config: requestflag.RequestConfig{
				BodyPath: "measurement_unit",
			},
		},
		&requestflag.StringFlag{
			Name:  "name",
			Usage: "Name of the meter",
			Config: requestflag.RequestConfig{
				BodyPath: "name",
			},
		},
		&requestflag.StringFlag{
			Name:  "description",
			Usage: "Optional description of the meter",
			Config: requestflag.RequestConfig{
				BodyPath: "description",
			},
		},
		&requestflag.YAMLFlag{
			Name:  "filter",
			Usage: "A filter structure that combines multiple conditions with logical conjunctions (AND/OR).\n\nSupports up to 3 levels of nesting to create complex filter expressions.\nEach filter has a conjunction (and/or) and clauses that can be either direct conditions or nested filters.",
			Config: requestflag.RequestConfig{
				BodyPath: "filter",
			},
		},
	},
	Action:          handleMetersCreate,
	HideHelpCommand: true,
}

var metersRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Perform retrieve operation",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name: "id",
		},
	},
	Action:          handleMetersRetrieve,
	HideHelpCommand: true,
}

var metersList = cli.Command{
	Name:  "list",
	Usage: "Perform list operation",
	Flags: []cli.Flag{
		&requestflag.BoolFlag{
			Name:  "archived",
			Usage: "List archived meters",
			Config: requestflag.RequestConfig{
				QueryPath: "archived",
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
	Action:          handleMetersList,
	HideHelpCommand: true,
}

var metersArchive = cli.Command{
	Name:  "archive",
	Usage: "Perform archive operation",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name: "id",
		},
	},
	Action:          handleMetersArchive,
	HideHelpCommand: true,
}

var metersUnarchive = cli.Command{
	Name:  "unarchive",
	Usage: "Perform unarchive operation",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name: "id",
		},
	},
	Action:          handleMetersUnarchive,
	HideHelpCommand: true,
}

func handleMetersCreate(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.MeterNewParams{}

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
	_, err = client.Meters.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "meters create", obj, format, transform)
}

func handleMetersRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
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
	_, err = client.Meters.Get(ctx, requestflag.CommandRequestValue[string](cmd, "id"), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "meters retrieve", obj, format, transform)
}

func handleMetersList(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.MeterListParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	if format == "raw" {
		var res []byte
		options = append(options, option.WithResponseBodyInto(&res))
		_, err = client.Meters.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "meters list", obj, format, transform)
	} else {
		iter := client.Meters.ListAutoPaging(ctx, params, options...)
		return streamOutput("meters list", func(w *os.File) error {
			for iter.Next() {
				item := iter.Current()
				obj := gjson.Parse(item.JSON.RawJSON())
				if err := ShowJSON(w, "meters list", obj, format, transform); err != nil {
					return err
				}
			}
			return iter.Err()
		})
	}
}

func handleMetersArchive(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
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

	return client.Meters.Archive(ctx, requestflag.CommandRequestValue[string](cmd, "id"), options...)
}

func handleMetersUnarchive(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
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

	return client.Meters.Unarchive(ctx, requestflag.CommandRequestValue[string](cmd, "id"), options...)
}
