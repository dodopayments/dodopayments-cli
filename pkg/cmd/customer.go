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

var customersCreate = cli.Command{
	Name:  "create",
	Usage: "Perform create operation",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name: "email",
			Config: requestflag.RequestConfig{
				BodyPath: "email",
			},
		},
		&requestflag.StringFlag{
			Name: "name",
			Config: requestflag.RequestConfig{
				BodyPath: "name",
			},
		},
		&requestflag.YAMLFlag{
			Name:  "metadata",
			Usage: "Additional metadata for the customer",
			Config: requestflag.RequestConfig{
				BodyPath: "metadata",
			},
		},
		&requestflag.StringFlag{
			Name: "phone-number",
			Config: requestflag.RequestConfig{
				BodyPath: "phone_number",
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
		&requestflag.StringFlag{
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
		&requestflag.StringFlag{
			Name: "customer-id",
		},
		&requestflag.YAMLFlag{
			Name:  "metadata",
			Usage: "Additional metadata for the customer",
			Config: requestflag.RequestConfig{
				BodyPath: "metadata",
			},
		},
		&requestflag.StringFlag{
			Name: "name",
			Config: requestflag.RequestConfig{
				BodyPath: "name",
			},
		},
		&requestflag.StringFlag{
			Name: "phone-number",
			Config: requestflag.RequestConfig{
				BodyPath: "phone_number",
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
		&requestflag.StringFlag{
			Name:  "email",
			Usage: "Filter by customer email",
			Config: requestflag.RequestConfig{
				QueryPath: "email",
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
	Action:          handleCustomersList,
	HideHelpCommand: true,
}

var customersRetrievePaymentMethods = cli.Command{
	Name:  "retrieve-payment-methods",
	Usage: "Perform retrieve-payment-methods operation",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name: "customer-id",
		},
	},
	Action:          handleCustomersRetrievePaymentMethods,
	HideHelpCommand: true,
}

func handleCustomersCreate(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.CustomerNewParams{}

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
	_, err = client.Customers.New(ctx, params, options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "customers create", obj, format, transform)
}

func handleCustomersRetrieve(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("customer-id") && len(unusedArgs) > 0 {
		cmd.Set("customer-id", unusedArgs[0])
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
	_, err = client.Customers.Get(ctx, requestflag.CommandRequestValue[string](cmd, "customer-id"), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "customers retrieve", obj, format, transform)
}

func handleCustomersUpdate(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("customer-id") && len(unusedArgs) > 0 {
		cmd.Set("customer-id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.CustomerUpdateParams{}

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
	_, err = client.Customers.Update(
		ctx,
		requestflag.CommandRequestValue[string](cmd, "customer-id"),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "customers update", obj, format, transform)
}

func handleCustomersList(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.CustomerListParams{}

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
		_, err = client.Customers.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "customers list", obj, format, transform)
	} else {
		iter := client.Customers.ListAutoPaging(ctx, params, options...)
		return streamOutput("customers list", func(w *os.File) error {
			for iter.Next() {
				item := iter.Current()
				obj := gjson.Parse(item.JSON.RawJSON())
				if err := ShowJSON(w, "customers list", obj, format, transform); err != nil {
					return err
				}
			}
			return iter.Err()
		})
	}
}

func handleCustomersRetrievePaymentMethods(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("customer-id") && len(unusedArgs) > 0 {
		cmd.Set("customer-id", unusedArgs[0])
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
	_, err = client.Customers.GetPaymentMethods(ctx, requestflag.CommandRequestValue[string](cmd, "customer-id"), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "customers retrieve-payment-methods", obj, format, transform)
}
