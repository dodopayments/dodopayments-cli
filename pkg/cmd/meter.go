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

var metersCreate = cli.Command{
	Name:  "create",
	Usage: "Perform create operation",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name:  "aggregation.type",
			Usage: "Aggregation type for the meter",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "aggregation.type",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "aggregation.key",
			Usage: "Required when type is not COUNT",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "aggregation.key",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "event-name",
			Usage: "Event name to track",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "event_name",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "measurement-unit",
			Usage: "measurement unit",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "measurement_unit",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "name",
			Usage: "Name of the meter",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "name",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "description",
			Usage: "Optional description of the meter",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "description",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "filter.clauses.key",
			Usage: "Direct filter conditions - array of condition objects with key, operator, and value",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "filter.clauses.#.key",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "filter.clauses.operator",
			Usage: "Direct filter conditions - array of condition objects with key, operator, and value",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "filter.clauses.#.operator",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "filter.clauses.value",
			Usage: "Direct filter conditions - array of condition objects with key, operator, and value",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "filter.clauses.#.value",
			},
		},
		&jsonflag.JSONAnyFlag{
			Name:  "filter.+clause",
			Usage: "Direct filter conditions - array of condition objects with key, operator, and value",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "filter.clauses.-1",
				SetValue: map[string]interface{}{},
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "filter.clauses.clauses.key",
			Usage: "Nested filters - supports up to 3 levels deep",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "filter.clauses.#.clauses.#.key",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "filter.clauses.clauses.operator",
			Usage: "Nested filters - supports up to 3 levels deep",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "filter.clauses.#.clauses.#.operator",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "filter.clauses.clauses.value",
			Usage: "Nested filters - supports up to 3 levels deep",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "filter.clauses.#.clauses.#.value",
			},
		},
		&jsonflag.JSONAnyFlag{
			Name:  "filter.clauses.+clause",
			Usage: "Nested filters - supports up to 3 levels deep",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "filter.clauses.#.clauses.-1",
				SetValue: map[string]interface{}{},
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "filter.clauses.clauses.clauses.key",
			Usage: "Nested filters - supports up to 3 levels deep",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "filter.clauses.#.clauses.#.clauses.#.key",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "filter.clauses.clauses.clauses.operator",
			Usage: "Nested filters - supports up to 3 levels deep",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "filter.clauses.#.clauses.#.clauses.#.operator",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "filter.clauses.clauses.clauses.value",
			Usage: "Nested filters - supports up to 3 levels deep",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "filter.clauses.#.clauses.#.clauses.#.value",
			},
		},
		&jsonflag.JSONAnyFlag{
			Name:  "filter.clauses.clauses.+clause",
			Usage: "Nested filters - supports up to 3 levels deep",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "filter.clauses.#.clauses.#.clauses.-1",
				SetValue: map[string]interface{}{},
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "filter.clauses.clauses.clauses.clauses.key",
			Usage: "Nested filters - supports up to 3 levels deep",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "filter.clauses.#.clauses.#.clauses.#.clauses.#.key",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "filter.clauses.clauses.clauses.clauses.operator",
			Usage: "Nested filters - supports up to 3 levels deep",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "filter.clauses.#.clauses.#.clauses.#.clauses.#.operator",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "filter.clauses.clauses.clauses.clauses.value",
			Usage: "Nested filters - supports up to 3 levels deep",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "filter.clauses.#.clauses.#.clauses.#.clauses.#.value",
			},
		},
		&jsonflag.JSONAnyFlag{
			Name:  "filter.clauses.clauses.clauses.+clause",
			Usage: "Nested filters - supports up to 3 levels deep",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "filter.clauses.#.clauses.#.clauses.#.clauses.-1",
				SetValue: map[string]interface{}{},
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "filter.clauses.clauses.clauses.conjunction",
			Usage: "Nested filters - supports up to 3 levels deep",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "filter.clauses.#.clauses.#.clauses.#.conjunction",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "filter.clauses.clauses.conjunction",
			Usage: "Nested filters - supports up to 3 levels deep",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "filter.clauses.#.clauses.#.conjunction",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "filter.clauses.conjunction",
			Usage: "Nested filters - supports up to 3 levels deep",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "filter.clauses.#.conjunction",
			},
		},
		&jsonflag.JSONStringFlag{
			Name:  "filter.conjunction",
			Usage: "Logical conjunction to apply between clauses (and/or)",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "filter.conjunction",
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
		&cli.StringFlag{
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
		&jsonflag.JSONBoolFlag{
			Name:  "archived",
			Usage: "List archived meters",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Query,
				Path:     "archived",
				SetValue: true,
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "page-number",
			Usage: "Page number default is 0",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "page_number",
			},
		},
		&jsonflag.JSONIntFlag{
			Name:  "page-size",
			Usage: "Page size default is 10 max is 100",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "page_size",
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
		&cli.StringFlag{
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
		&cli.StringFlag{
			Name: "id",
		},
	},
	Action:          handleMetersUnarchive,
	HideHelpCommand: true,
}

func handleMetersCreate(_ context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.MeterNewParams{}
	var res []byte
	_, err := cc.client.Meters.New(
		context.TODO(),
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
	return ShowJSON("meters create", json, format, transform)
}

func handleMetersRetrieve(_ context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	var res []byte
	_, err := cc.client.Meters.Get(
		context.TODO(),
		cmd.Value("id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("meters retrieve", json, format, transform)
}

func handleMetersList(_ context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.MeterListParams{}
	var res []byte
	_, err := cc.client.Meters.List(
		context.TODO(),
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
	return ShowJSON("meters list", json, format, transform)
}

func handleMetersArchive(_ context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	return cc.client.Meters.Archive(
		context.TODO(),
		cmd.Value("id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
	)
}

func handleMetersUnarchive(_ context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	return cc.client.Meters.Unarchive(
		context.TODO(),
		cmd.Value("id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
	)
}
