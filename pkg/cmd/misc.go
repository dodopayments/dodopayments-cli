// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"context"

	"github.com/dodopayments/dodopayments-go/option"
	"github.com/urfave/cli/v3"
)

var miscListSupportedCountries = cli.Command{
	Name:            "list-supported-countries",
	Usage:           "Perform list-supported-countries operation",
	Flags:           []cli.Flag{},
	Action:          handleMiscListSupportedCountries,
	HideHelpCommand: true,
}

func handleMiscListSupportedCountries(ctx context.Context, cmd *cli.Command) error {
	cc := getAPICommandContext(cmd)
	res := []byte{}
	_, err := cc.client.Misc.ListSupportedCountries(
		context.TODO(),
		option.WithMiddleware(cc.AsMiddleware()),
		option.WithResponseBodyInto(&res),
	)
	if err != nil {
		return err
	}

	format := cmd.Root().String("format")
	return ShowJSON("misc list-supported-countries", string(res), format)
}
