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

var webhooksCreate = cli.Command{
	Name:  "create",
	Usage: "Create a new webhook",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name:  "url",
			Usage: "Url of the webhook",
			Config: requestflag.RequestConfig{
				BodyPath: "url",
			},
		},
		&requestflag.StringFlag{
			Name: "description",
			Config: requestflag.RequestConfig{
				BodyPath: "description",
			},
		},
		&requestflag.BoolFlag{
			Name:  "disabled",
			Usage: "Create the webhook in a disabled state.\n\nDefault is false",
			Config: requestflag.RequestConfig{
				BodyPath: "disabled",
			},
		},
		&requestflag.StringSliceFlag{
			Name:  "filter-type",
			Usage: "Filter events to the webhook.\n\nWebhook event will only be sent for events in the list.",
			Config: requestflag.RequestConfig{
				BodyPath: "filter_types",
			},
		},
		&requestflag.YAMLFlag{
			Name:  "headers",
			Usage: "Custom headers to be passed",
			Config: requestflag.RequestConfig{
				BodyPath: "headers",
			},
		},
		&requestflag.StringFlag{
			Name:  "idempotency-key",
			Usage: "The request's idempotency key",
			Config: requestflag.RequestConfig{
				BodyPath: "idempotency_key",
			},
		},
		&requestflag.YAMLFlag{
			Name:  "metadata",
			Usage: "Metadata to be passed to the webhook\nDefaut is {}",
			Config: requestflag.RequestConfig{
				BodyPath: "metadata",
			},
		},
		&requestflag.IntFlag{
			Name: "rate-limit",
			Config: requestflag.RequestConfig{
				BodyPath: "rate_limit",
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
		&requestflag.StringFlag{
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
		&requestflag.StringFlag{
			Name: "webhook-id",
		},
		&requestflag.StringFlag{
			Name:  "description",
			Usage: "Description of the webhook",
			Config: requestflag.RequestConfig{
				BodyPath: "description",
			},
		},
		&requestflag.BoolFlag{
			Name:  "disabled",
			Usage: "To Disable the endpoint, set it to true.",
			Config: requestflag.RequestConfig{
				BodyPath: "disabled",
			},
		},
		&requestflag.StringSliceFlag{
			Name:  "filter-type",
			Usage: "Filter events to the endpoint.\n\nWebhook event will only be sent for events in the list.",
			Config: requestflag.RequestConfig{
				BodyPath: "filter_types",
			},
		},
		&requestflag.YAMLFlag{
			Name:  "metadata",
			Usage: "Metadata",
			Config: requestflag.RequestConfig{
				BodyPath: "metadata",
			},
		},
		&requestflag.IntFlag{
			Name:  "rate-limit",
			Usage: "Rate limit",
			Config: requestflag.RequestConfig{
				BodyPath: "rate_limit",
			},
		},
		&requestflag.StringFlag{
			Name:  "url",
			Usage: "Url endpoint",
			Config: requestflag.RequestConfig{
				BodyPath: "url",
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
		&requestflag.StringFlag{
			Name:  "iterator",
			Usage: "The iterator returned from a prior invocation",
			Config: requestflag.RequestConfig{
				QueryPath: "iterator",
			},
		},
		&requestflag.IntFlag{
			Name:  "limit",
			Usage: "Limit the number of returned items",
			Config: requestflag.RequestConfig{
				QueryPath: "limit",
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
		&requestflag.StringFlag{
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
		&requestflag.StringFlag{
			Name: "webhook-id",
		},
	},
	Action:          handleWebhooksRetrieveSecret,
	HideHelpCommand: true,
}

func handleWebhooksCreate(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.WebhookNewParams{}

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
	_, err = client.Webhooks.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "webhooks create", obj, format, transform)
}

func handleWebhooksRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Webhooks.Get(ctx, requestflag.CommandRequestValue[string](cmd, "webhook-id"), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "webhooks retrieve", obj, format, transform)
}

func handleWebhooksUpdate(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("webhook-id") && len(unusedArgs) > 0 {
		cmd.Set("webhook-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.WebhookUpdateParams{}

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
	_, err = client.Webhooks.Update(
		ctx,
		requestflag.CommandRequestValue[string](cmd, "webhook-id"),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "webhooks update", obj, format, transform)
}

func handleWebhooksList(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.WebhookListParams{}

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
		_, err = client.Webhooks.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "webhooks list", obj, format, transform)
	} else {
		iter := client.Webhooks.ListAutoPaging(ctx, params, options...)
		return streamOutput("webhooks list", func(w *os.File) error {
			for iter.Next() {
				item := iter.Current()
				obj := gjson.Parse(item.JSON.RawJSON())
				if err := ShowJSON(w, "webhooks list", obj, format, transform); err != nil {
					return err
				}
			}
			return iter.Err()
		})
	}
}

func handleWebhooksDelete(ctx context.Context, cmd *cli.Command) error {
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

	return client.Webhooks.Delete(ctx, requestflag.CommandRequestValue[string](cmd, "webhook-id"), options...)
}

func handleWebhooksRetrieveSecret(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.Webhooks.GetSecret(ctx, requestflag.CommandRequestValue[string](cmd, "webhook-id"), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "webhooks retrieve-secret", obj, format, transform)
}
