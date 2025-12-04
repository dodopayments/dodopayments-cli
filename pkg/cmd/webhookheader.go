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

var webhooksHeadersRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Get a webhook by id",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name: "webhook-id",
		},
	},
	Action:          handleWebhooksHeadersRetrieve,
	HideHelpCommand: true,
}

var webhooksHeadersUpdate = cli.Command{
	Name:  "update",
	Usage: "Patch a webhook by id",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name: "webhook-id",
		},
		&requestflag.YAMLFlag{
			Name:  "headers",
			Usage: "Object of header-value pair to update or add",
			Config: requestflag.RequestConfig{
				BodyPath: "headers",
			},
		},
	},
	Action:          handleWebhooksHeadersUpdate,
	HideHelpCommand: true,
}

func handleWebhooksHeadersRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("webhook-id") && len(unusedArgs) > 0 {
		cmd.Set("webhook-id", unusedArgs[0])
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
	_, err = client.Webhooks.Headers.Get(
		ctx,
		requestflag.CommandRequestValue[string](cmd, "webhook-id"),
		options...,
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("webhooks:headers retrieve", json, format, transform)
}

func handleWebhooksHeadersUpdate(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("webhook-id") && len(unusedArgs) > 0 {
		cmd.Set("webhook-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.WebhookHeaderUpdateParams{}

	options, err := flagOptions(
		cmd,
		apiquery.NestedQueryFormatBrackets,
		apiquery.ArrayQueryFormatComma,
		ApplicationJSON,
	)
	if err != nil {
		return err
	}
	return client.Webhooks.Headers.Update(
		ctx,
		requestflag.CommandRequestValue[string](cmd, "webhook-id"),
		params,
		options...,
	)
}
