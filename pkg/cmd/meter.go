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

var metersCreate = requestflag.WithInnerFlags(cli.Command{
	Name:  "create",
	Usage: "Perform create operation",
	Flags: []cli.Flag{
		&requestflag.Flag[map[string]any]{
			Name:     "aggregation",
			Required: true,
			BodyPath: "aggregation",
		},
		&requestflag.Flag[string]{
			Name:     "event-name",
			Usage:    "Event name to track",
			Required: true,
			BodyPath: "event_name",
		},
		&requestflag.Flag[string]{
			Name:     "measurement-unit",
			Usage:    "measurement unit",
			Required: true,
			BodyPath: "measurement_unit",
		},
		&requestflag.Flag[string]{
			Name:     "name",
			Usage:    "Name of the meter",
			Required: true,
			BodyPath: "name",
		},
		&requestflag.Flag[string]{
			Name:     "description",
			Usage:    "Optional description of the meter",
			BodyPath: "description",
		},
		&requestflag.Flag[map[string]any]{
			Name:     "filter",
			Usage:    "A filter structure that combines multiple conditions with logical conjunctions (AND/OR).\n\nSupports up to 3 levels of nesting to create complex filter expressions.\nEach filter has a conjunction (and/or) and clauses that can be either direct conditions or nested filters.",
			BodyPath: "filter",
		},
	},
	Action:          handleMetersCreate,
	HideHelpCommand: true,
}, map[string][]requestflag.HasOuterFlag{
	"aggregation": {
		&requestflag.InnerFlag[string]{
			Name:       "aggregation.type",
			Usage:      "Aggregation type for the meter",
			InnerField: "type",
		},
		&requestflag.InnerFlag[string]{
			Name:       "aggregation.key",
			Usage:      "Required when type is not COUNT",
			InnerField: "key",
		},
	},
	"filter": {
		&requestflag.InnerFlag[[]map[string]any]{
			Name:       "filter.clauses",
			Usage:      "Filter clauses - can be direct conditions or nested filters (up to 3 levels deep)",
			InnerField: "clauses",
		},
		&requestflag.InnerFlag[string]{
			Name:       "filter.conjunction",
			Usage:      "Logical conjunction to apply between clauses (and/or)",
			InnerField: "conjunction",
		},
	},
})

var metersRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Perform retrieve operation",
	Flags: []cli.Flag{
		&requestflag.Flag[string]{
			Name:     "id",
			Required: true,
		},
	},
	Action:          handleMetersRetrieve,
	HideHelpCommand: true,
}

var metersList = cli.Command{
	Name:  "list",
	Usage: "Perform list operation",
	Flags: []cli.Flag{
		&requestflag.Flag[bool]{
			Name:      "archived",
			Usage:     "List archived meters",
			QueryPath: "archived",
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
	Action:          handleMetersList,
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
		false,
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
		EmptyBody,
		false,
	)
	if err != nil {
		return err
	}

	var res []byte
	options = append(options, option.WithResponseBodyInto(&res))
	_, err = client.Meters.Get(ctx, cmd.Value("id").(string), options...)
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
		_, err = client.Meters.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "meters list", obj, format, transform)
	} else {
		iter := client.Meters.ListAutoPaging(ctx, params, options...)
		return ShowJSONIterator(os.Stdout, "meters list", iter, format, transform)
	}
}
