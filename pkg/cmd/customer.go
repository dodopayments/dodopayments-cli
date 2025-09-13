// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"

	"github.com/dodopayments/dodopayments-cli/pkg/jsonflag"
	"github.com/dodopayments/dodopayments-go"
	"github.com/dodopayments/dodopayments-go/option"
	"github.com/urfave/cli/v3"
)

var customersCreate = cli.Command{
	Name:  "create",
	Usage: "Perform create operation",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name: "email",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "email",
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
			Name: "phone-number",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "phone_number",
			},
		},
	},
	Action:          handleCustomersCreate,
	HideHelpCommand: true,
}

var customersRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Perform retrieve operation",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "customer-id",
		},
	},
	Action:          handleCustomersRetrieve,
	HideHelpCommand: true,
}

var customersUpdate = cli.Command{
	Name:  "update",
	Usage: "Perform update operation",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name: "customer-id",
		},
		&jsonflag.JSONStringFlag{
			Name: "name",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "name",
			},
		},
		&jsonflag.JSONStringFlag{
			Name: "phone-number",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Body,
				Path: "phone_number",
			},
		},
	},
	Action:          handleCustomersUpdate,
	HideHelpCommand: true,
}

var customersList = cli.Command{
	Name:  "list",
	Usage: "Perform list operation",
	Flags: []cli.Flag{
		&jsonflag.JSONStringFlag{
			Name: "email",
			Config: jsonflag.JSONConfig{
				Kind: jsonflag.Query,
				Path: "email",
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
	Action:          handleCustomersList,
	HideHelpCommand: true,
}

func handleCustomersCreate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := dodopayments.CustomerNewParams{}
	res := []byte{}
	_, err := cc.client.Customers.New(
		context.TODO(),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("customers create", string(res), format)
}

func handleCustomersRetrieve(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	res := []byte{}
	_, err := cc.client.Customers.Get(
		context.TODO(),
		cmd.Value("customer-id").(string),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("customers retrieve", string(res), format)
}

func handleCustomersUpdate(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := dodopayments.CustomerUpdateParams{}
	res := []byte{}
	_, err := cc.client.Customers.Update(
		context.TODO(),
		cmd.Value("customer-id").(string),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("customers update", string(res), format)
}

func handleCustomersList(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	params := dodopayments.CustomerListParams{}
	res := []byte{}
	_, err := cc.client.Customers.List(
		context.TODO(),
		params,
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("customers list", string(res), format)
}
