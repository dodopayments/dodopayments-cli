// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cmd

import (
	"compress/gzip"
	"context"
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"

	docs "github.com/urfave/cli-docs/v3"
	"github.com/urfave/cli/v3"
)

var (
	Command       *cli.Command
	OutputFormats = []string{"auto", "explore", "json", "pretty", "raw", "yaml"}
)

func init() {
	Command = &cli.Command{
		Name:    "dodopayments",
		Usage:   "CLI for the Dodo Payments API",
		Version: Version,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "debug",
				Usage: "Enable debug logging",
			},
			&cli.StringFlag{
				Name:        "base-url",
				DefaultText: "url",
				Usage:       "Override the base URL for API requests",
			},
			&cli.StringFlag{
				Name:  "format",
				Usage: "The format for displaying response data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "format-error",
				Usage: "The format for displaying error data (one of: " + strings.Join(OutputFormats, ", ") + ")",
				Value: "auto",
				Validator: func(format string) error {
					if !slices.Contains(OutputFormats, strings.ToLower(format)) {
						return fmt.Errorf("format must be one of: %s", strings.Join(OutputFormats, ", "))
					}
					return nil
				},
			},
			&cli.StringFlag{
				Name:  "transform",
				Usage: "The GJSON transformation for data output.",
			},
			&cli.StringFlag{
				Name:  "transform-error",
				Usage: "The GJSON transformation for errors.",
			},
			&cli.StringFlag{
				Name:  "environment",
				Usage: "Set the environment for API requests",
			},
		},
		Commands: []*cli.Command{
			{
				Name:     "checkout-sessions",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&checkoutSessionsCreate,
					&checkoutSessionsRetrieve,
				},
			},
			{
				Name:     "payments",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&paymentsCreate,
					&paymentsRetrieve,
					&paymentsList,
					&paymentsRetrieveLineItems,
				},
			},
			{
				Name:     "subscriptions",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&subscriptionsCreate,
					&subscriptionsRetrieve,
					&subscriptionsUpdate,
					&subscriptionsList,
					&subscriptionsCharge,
					&subscriptionsRetrieveUsageHistory,
					&subscriptionsUpdatePaymentMethod,
				},
			},
			{
				Name:     "licenses",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&licensesActivate,
					&licensesValidate,
				},
			},
			{
				Name:     "license-keys",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&licenseKeysRetrieve,
					&licenseKeysUpdate,
					&licenseKeysList,
				},
			},
			{
				Name:     "license-key-instances",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&licenseKeyInstancesRetrieve,
					&licenseKeyInstancesUpdate,
					&licenseKeyInstancesList,
				},
			},
			{
				Name:     "customers",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&customersCreate,
					&customersRetrieve,
					&customersUpdate,
					&customersList,
					&customersRetrievePaymentMethods,
				},
			},
			{
				Name:     "customers:customer-portal",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&customersCustomerPortalCreate,
				},
			},
			{
				Name:     "customers:wallets",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&customersWalletsList,
				},
			},
			{
				Name:     "customers:wallets:ledger-entries",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&customersWalletsLedgerEntriesCreate,
					&customersWalletsLedgerEntriesList,
				},
			},
			{
				Name:     "refunds",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&refundsCreate,
					&refundsRetrieve,
					&refundsList,
				},
			},
			{
				Name:     "disputes",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&disputesRetrieve,
					&disputesList,
				},
			},
			{
				Name:     "payouts",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&payoutsList,
				},
			},
			{
				Name:     "products",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&productsCreate,
					&productsRetrieve,
					&productsList,
					&productsUpdateFiles,
				},
			},
			{
				Name:     "products:images",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&productsImagesUpdate,
				},
			},
			{
				Name:     "misc",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&miscListSupportedCountries,
				},
			},
			{
				Name:     "discounts",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&discountsCreate,
					&discountsRetrieve,
					&discountsUpdate,
					&discountsList,
				},
			},
			{
				Name:     "addons",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&addonsCreate,
					&addonsRetrieve,
					&addonsUpdate,
					&addonsList,
					&addonsUpdateImages,
				},
			},
			{
				Name:     "brands",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&brandsCreate,
					&brandsRetrieve,
					&brandsUpdate,
					&brandsList,
					&brandsUpdateImages,
				},
			},
			{
				Name:     "webhooks",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&webhooksCreate,
					&webhooksRetrieve,
					&webhooksUpdate,
					&webhooksList,
					&webhooksRetrieveSecret,
				},
			},
			{
				Name:     "webhooks:headers",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&webhooksHeadersRetrieve,
				},
			},
			{
				Name:     "usage-events",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&usageEventsRetrieve,
					&usageEventsList,
					&usageEventsIngest,
				},
			},
			{
				Name:     "meters",
				Category: "API RESOURCE",
				Commands: []*cli.Command{
					&metersCreate,
					&metersRetrieve,
					&metersList,
				},
			},
			{
				Name:            "@manpages",
				Usage:           "Generate documentation for 'man'",
				UsageText:       "dodopayments @manpages [-o dodopayments.1] [--gzip]",
				Hidden:          true,
				Action:          generateManpages,
				HideHelpCommand: true,
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "write manpages to the given folder",
						Value:   "man",
					},
					&cli.BoolFlag{
						Name:    "gzip",
						Aliases: []string{"z"},
						Usage:   "output gzipped manpage files to .gz",
						Value:   true,
					},
					&cli.BoolFlag{
						Name:    "text",
						Aliases: []string{"z"},
						Usage:   "output uncompressed text files",
						Value:   false,
					},
				},
			},
		},
		EnableShellCompletion:      true,
		ShellCompletionCommandName: "@completion",
		HideHelpCommand:            true,
	}
}

func generateManpages(ctx context.Context, c *cli.Command) error {
	manpage, err := docs.ToManWithSection(Command, 1)
	if err != nil {
		return err
	}
	dir := c.String("output")
	err = os.MkdirAll(filepath.Join(dir, "man1"), 0755)
	if err != nil {
		// handle error
	}
	if c.Bool("text") {
		file, err := os.Create(filepath.Join(dir, "man1", "dodopayments.1"))
		if err != nil {
			return err
		}
		defer file.Close()
		if _, err := file.WriteString(manpage); err != nil {
			return err
		}
	}
	if c.Bool("gzip") {
		file, err := os.Create(filepath.Join(dir, "man1", "dodopayments.1.gz"))
		if err != nil {
			return err
		}
		defer file.Close()
		gzWriter := gzip.NewWriter(file)
		defer gzWriter.Close()
		_, err = gzWriter.Write([]byte(manpage))
		if err != nil {
			return err
		}
	}
	fmt.Printf("Wrote manpages to %s\n", dir)
	return nil
}
