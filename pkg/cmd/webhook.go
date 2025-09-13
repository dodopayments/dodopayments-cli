// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"

	"github.com/dodopayments/dodopayments-cli/pkg/jsonflag"
	"github.com/dodopayments/dodopayments-go"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/urfave/cli/v3"
)

var webhooksCreate = cli.Command{
	Name:  "create",
	Usage: "Create a new webhook",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name: "url",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "url",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "description",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "description",
			},
		},
		&jsonflag.JSONBoolFlag{
			Name: "disabled",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "disabled",
				SetValue: true,
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "filter-types",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "filter_types.#",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "+filter_type",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "filter_types.-1",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "idempotency-key",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "idempotency_key",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "rate-limit",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "rate_limit",
			},
		},
	},
	Action:          handleWebhooksCreate,
	HideHelpCommand: true,
}

var webhooksRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Get a webhook by id",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "webhook-id",
		},
	},
	Action:          handleWebhooksRetrieve,
	HideHelpCommand: true,
}

var webhooksUpdate = cli.Command{
	Name:  "update",
	Usage: "Patch a webhook by id",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "webhook-id",
		},
		&jsonflag.JSONStringFlag{
			Name: "description",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "description",
			},
		},
		&jsonflag.JSONBoolFlag{
			Name: "disabled",
			Config: jsonflag.JSONConfig{
				Kind:     jsonflag.Body,
				Path:     "disabled",
				SetValue: true,
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "filter-types",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "filter_types.#",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "+filter_type",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "filter_types.-1",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "rate-limit",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "rate_limit",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "url",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "url",
			},
		},
	},
	Action:          handleWebhooksUpdate,
	HideHelpCommand: true,
}

var webhooksList = cli.Command{
	Name:  "list",
	Usage: "List all webhooks",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name: "iterator",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "iterator",
			},
		},
		&jsonflag.JSONIntFlag{
			Name: "limit",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "limit",
			},
		},
	},
	Action:          handleWebhooksList,
	HideHelpCommand: true,
}

var webhooksDelete = cli.Command{
	Name:  "delete",
	Usage: "Delete a webhook by id",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "webhook-id",
		},
	},
	Action:          handleWebhooksDelete,
	HideHelpCommand: true,
}

var webhooksRetrieveSecret = cli.Command{
	Name:  "retrieve-secret",
	Usage: "Get webhook secret by id",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "webhook-id",
		},
	},
	Action:          handleWebhooksRetrieveSecret,
	HideHelpCommand: true,
}

func handleWebhooksCreate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := dodopayments.WebhookNewParams{}
	res := []byte{}
	_, err := cc.client.Webhooks.New(
		context.TODO(),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("webhooks create", string(res), format)
}

func handleWebhooksRetrieve(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	res := []byte{}
	_, err := cc.client.Webhooks.Get(
		context.TODO(),
		cmd.Value("webhook-id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("webhooks retrieve", string(res), format)
}

func handleWebhooksUpdate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := dodopayments.WebhookUpdateParams{}
	res := []byte{}
	_, err := cc.client.Webhooks.Update(
		context.TODO(),
		cmd.Value("webhook-id").(string),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("webhooks update", string(res), format)
}

func handleWebhooksList(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := dodopayments.WebhookListParams{}
	res := []byte{}
	_, err := cc.client.Webhooks.List(
		context.TODO(),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("webhooks list", string(res), format)
}

func handleWebhooksDelete(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	return cc.client.Webhooks.Delete(
		context.TODO(),
		cmd.Value("webhook-id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
	)
}

func handleWebhooksRetrieveSecret(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	res := []byte{}
	_, err := cc.client.Webhooks.GetSecret(
		context.TODO(),
		cmd.Value("webhook-id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("webhooks retrieve-secret", string(res), format)
}
