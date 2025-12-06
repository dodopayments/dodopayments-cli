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

var licenseKeyInstancesRetrieve = cli.Command{
	Name:  "retrieve",
	Usage: "Perform retrieve operation",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name: "id",
		},
	},
	Action:          handleLicenseKeyInstancesRetrieve,
	HideHelpCommand: true,
}

var licenseKeyInstancesUpdate = cli.Command{
	Name:  "update",
	Usage: "Perform update operation",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name: "id",
		},
		&requestflag.StringFlag{
			Name: "name",
			Config: requestflag.RequestConfig{
				BodyPath: "name",
			},
		},
	},
	Action:          handleLicenseKeyInstancesUpdate,
	HideHelpCommand: true,
}

var licenseKeyInstancesList = cli.Command{
	Name:  "list",
	Usage: "Perform list operation",
	Flags: []cli.Flag{
		&requestflag.StringFlag{
			Name:  "license-key-id",
			Usage: "Filter by license key ID",
			Config: requestflag.RequestConfig{
				QueryPath: "license_key_id",
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
	Action:          handleLicenseKeyInstancesList,
	HideHelpCommand: true,
}

func handleLicenseKeyInstancesRetrieve(ctx context.Context, cmd *cli.Command) error {
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
	_, err = client.LicenseKeyInstances.Get(ctx, requestflag.CommandRequestValue[string](cmd, "id"), options...)
	if err != nil {
		return err
	}

	obj := gjson.ParseBytes(res)
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON(os.Stdout, "license-key-instances retrieve", obj, format, transform)
}

func handleLicenseKeyInstancesUpdate(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()
	if !cmd.IsSet("id") && len(unusedArgs) > 0 {
		cmd.Set("id", unusedArgs[0])
		unusedArgs = unusedArgs[1:]
	}
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.LicenseKeyInstanceUpdateParams{}

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
	_, err = client.LicenseKeyInstances.Update(
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
	return ShowJSON(os.Stdout, "license-key-instances update", obj, format, transform)
}

func handleLicenseKeyInstancesList(ctx context.Context, cmd *cli.Command) error {
	client := dodopayments.NewClient(getDefaultRequestOptions(cmd)...)
	unusedArgs := cmd.Args().Slice()

	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	params := dodopayments.LicenseKeyInstanceListParams{}

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
		_, err = client.LicenseKeyInstances.List(ctx, params, options...)
		if err != nil {
			return err
		}
		obj := gjson.ParseBytes(res)
		return ShowJSON(os.Stdout, "license-key-instances list", obj, format, transform)
	} else {
		iter := client.LicenseKeyInstances.ListAutoPaging(ctx, params, options...)
		return streamOutput("license-key-instances list", func(w *os.File) error {
			for iter.Next() {
				item := iter.Current()
				obj := gjson.Parse(item.JSON.RawJSON())
				if err := ShowJSON(w, "license-key-instances list", obj, format, transform); err != nil {
					return err
				}
			}
			return iter.Err()
		})
	}
}
