// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"
	"fmt"

	"github.com/dodopayments/dodopayments-go/option"
	"github.com/tidwall/gjson"
	"github.com/urfave/cli/v3"
)

var miscListSupportedCountries = cli.Command{
	Name:            "list-supported-countries",
	Usage:           "Perform list-supported-countries operation",
	Flags:           []cli.Flag{},
	Action:          handleMiscListSupportedCountries,
	HideHelpCommand: true,
}

func handleMiscListSupportedCountries(_ context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	unusedArgs := cmd.Args().Slice()
	if len(unusedArgs) > 0 {
		return fmt.Errorf("Unexpected extra arguments: %v", unusedArgs)
	}
	var res []byte
	_, err := cc.client.Misc.ListSupportedCountries(
		context.TODO(),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	json := gjson.Parse(string(res))
	format := cmd.Root().String("format")
	transform := cmd.Root().String("transform")
	return ShowJSON("misc list-supported-countries", json, format, transform)
}
