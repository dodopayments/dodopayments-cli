// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"

	"github.com/dodopayments/dodopayments-go"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/urfave/cli/v3"
)

var webhooksHeadersRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Get a webhook by id",
	Flags: []cli.Flag{
		&cli.StringFlag{
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
		&cli.StringFlag{
			Name: "webhook-id",
		},
	},
	Action:          handleWebhooksHeadersUpdate,
	HideHelpCommand: true,
}

func handleWebhooksHeadersRetrieve(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	res := []byte{}
	_, err := cc.client.Webhooks.Headers.Get(
		context.TODO(),
		cmd.Value("webhook-id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("webhooks:headers retrieve", string(res), format)
}

func handleWebhooksHeadersUpdate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := dodopayments.WebhookHeaderUpdateParams{}
	return cc.client.Webhooks.Headers.Update(
		context.TODO(),
		cmd.Value("webhook-id").(string),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
	)
}
