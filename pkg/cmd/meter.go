// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"

	"github.com/dodopayments/dodopayments-go"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/stainless-sdks/dodo-payments-cli/pkg/jsonflag"
	"github.com/urfave/cli/v3"
)

var metersCreate = cli.Command{
	Name:  "create",
	Usage: "Perform create operation",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name: "aggregation.type",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "aggregation.type",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "aggregation.key",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "aggregation.key",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "event-name",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "event_name",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "measurement-unit",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "measurement_unit",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "name",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "name",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "description",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "description",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "filter.clauses.key",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "filter.clauses.#.key",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "filter.clauses.operator",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "filter.clauses.#.operator",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "filter.clauses.value",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "filter.clauses.#.value",
			},
		},
		&jsonflag.JSONAnyFlag{
			Name: "filter.+clause",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "filter.clauses.-1",
				SetValue: map[string]interface{}{},
			},
			Value: map[string]interface{}{},
		},
		&jsonflag.JSONStringFlag{
			Name: "filter.clauses.clauses.key",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "filter.clauses.#.clauses.#.key",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "filter.clauses.clauses.operator",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "filter.clauses.#.clauses.#.operator",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "filter.clauses.clauses.value",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "filter.clauses.#.clauses.#.value",
			},
		},
		&jsonflag.JSONAnyFlag{
			Name: "filter.clauses.+clause",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "filter.clauses.#.clauses.-1",
				SetValue: map[string]interface{}{},
			},
			Value: map[string]interface{}{},
		},
		&jsonflag.JSONStringFlag{
			Name: "filter.clauses.clauses.clauses.key",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "filter.clauses.#.clauses.#.clauses.#.key",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "filter.clauses.clauses.clauses.operator",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "filter.clauses.#.clauses.#.clauses.#.operator",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "filter.clauses.clauses.clauses.value",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "filter.clauses.#.clauses.#.clauses.#.value",
			},
		},
		&jsonflag.JSONAnyFlag{
			Name: "filter.clauses.clauses.+clause",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "filter.clauses.#.clauses.#.clauses.-1",
				SetValue: map[string]interface{}{},
			},
			Value: map[string]interface{}{},
		},
		&jsonflag.JSONStringFlag{
			Name: "filter.clauses.clauses.clauses.clauses.key",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "filter.clauses.#.clauses.#.clauses.#.clauses.#.key",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "filter.clauses.clauses.clauses.clauses.operator",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "filter.clauses.#.clauses.#.clauses.#.clauses.#.operator",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "filter.clauses.clauses.clauses.clauses.value",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "filter.clauses.#.clauses.#.clauses.#.clauses.#.value",
			},
		},
		&jsonflag.JSONAnyFlag{
			Name: "filter.clauses.clauses.clauses.+clause",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "filter.clauses.#.clauses.#.clauses.#.clauses.-1",
				SetValue: map[string]interface{}{},
			},
			Value: map[string]interface{}{},
		},
		&jsonflag.JSONStringFlag{
			Name: "filter.clauses.clauses.clauses.conjunction",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "filter.clauses.#.clauses.#.clauses.#.conjunction",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "filter.clauses.clauses.conjunction",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "filter.clauses.#.clauses.#.conjunction",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "filter.clauses.conjunction",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "filter.clauses.#.conjunction",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "filter.conjunction",
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
			Name: "archived",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Query,
				Path:     "archived",
				SetValue: true,
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

func handleMetersCreate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := dodopayments.MeterNewParams{}
	res := []byte{}
	_, err := cc.client.Meters.New(
		context.TODO(),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("meters create", string(res), format)
}

func handleMetersRetrieve(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	res := []byte{}
	_, err := cc.client.Meters.Get(
		context.TODO(),
		cmd.Value("id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("meters retrieve", string(res), format)
}

func handleMetersList(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := dodopayments.MeterListParams{}
	res := []byte{}
	_, err := cc.client.Meters.List(
		context.TODO(),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("meters list", string(res), format)
}

func handleMetersArchive(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	return cc.client.Meters.Archive(
		context.TODO(),
		cmd.Value("id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
	)
}

func handleMetersUnarchive(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	return cc.client.Meters.Unarchive(
		context.TODO(),
		cmd.Value("id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
	)
}
