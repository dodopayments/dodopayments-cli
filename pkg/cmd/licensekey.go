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

var licenseKeysRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Perform retrieve operation",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name: "id",
		},
	},
	Action:          handleLicenseKeysRetrieve,
	HideHelpCommand: true,
}

var licenseKeysUpdate = cli.Command{
	Name:  "update",
	Usage: "Perform update operation",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name: "id",
		},
		&requestflag.IntFlag{
			Name:  "activations-limit",
			Usage: "The updated activation limit for the license key.\nUse `null` to remove the limit, or omit this field to leave it unchanged.",
			Config: requestflag.RequestConfig{
				BodyPath: "activations_limit",
			},
		},
		&requestflag.BoolFlag{
			Name:  "disabled",
			Usage: "Indicates whether the license key should be disabled.\nA value of `true` disables the key, while `false` enables it. Omit this field to leave it unchanged.",
			Config: requestflag.RequestConfig{
				BodyPath: "disabled",
			},
		},
		&requestflag.DateTimeFlag{
			Name:  "expires-at",
			Usage: "The updated expiration timestamp for the license key in UTC.\nUse `null` to remove the expiration date, or omit this field to leave it unchanged.",
			Config: requestflag.RequestConfig{
				BodyPath: "expires_at",
			},
		},
	},
	Action:          handleLicenseKeysUpdate,
	HideHelpCommand: true,
}

var licenseKeysList = cli.Command{
	Name:  "list",
	Usage: "Perform list operation",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name:  "customer-id",
			Usage: "Filter by customer ID",
			Config: requestflag.RequestConfig{
				QueryPath: "customer_id",
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
		&requestflag.StringFlag{
			Name:  "product-id",
			Usage: "Filter by product ID",
			Config: requestflag.RequestConfig{
				QueryPath: "product_id",
			},
		},
		&requestflag.StringFlag{
			Name:  "status",
			Usage: "Filter by license key status",
			Config: requestflag.RequestConfig{
				QueryPath: "status",
			},
		},
	},
	Action:          handleLicenseKeysList,
	HideHelpCommand: true,
}

func handleLicenseKeysRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.LicenseKeys.Get(ctx, requestflag.CommandRequestValue[string](cmd, "id"), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "license-keys retrieve", obj, format, transform)
}

func handleLicenseKeysUpdate(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.LicenseKeyUpdateParams{}

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
	_, err = client.LicenseKeys.Update(
		ctx,
		requestflag.CommandRequestValue[string](cmd, "id"),
		params,
		options...,
	)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "license-keys update", obj, format, transform)
}

func handleLicenseKeysList(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.LicenseKeyListParams{}

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
		_, err = client.LicenseKeys.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "license-keys list", obj, format, transform)
	} else {
		iter := client.LicenseKeys.ListAutoPaging(ctx, params, options...)
		return streamOutput("license-keys list", func(w *os.File) error {
			for iter.Next() {
				item := iter.Current()
				obj := gjson.Parse(item.JSON.RawJSON())
				if err := ShowJSON(w, "license-keys list", obj, format, transform); err != nil {
					return err
				}
			}
			return iter.Err()
		})
	}
}
