# Dodo Payments CLI

The official CLI for the [Dodo Payments REST API](https://docs.dodopayments.com/api-reference/introduction).

It is generated with [Stainless](https://www.stainless.com/).

<!-- x-release-please-start-version -->

## Installation

### Installing with Go

```sh
go install 'github.com/dodopayments/dodopayments-cli/cmd/dodopayments@latest'
```

<!-- x-release-please-end -->

### Running Locally

```sh
./scripts/run args...
```

## Usage

The CLI follows a resource-based command structure:

```sh
dodopayments [resource] [command] [flags]
```

```sh
dodopayments checkout-sessions create \
  --product-cart '{product_id: product_id, quantity: 0, addons: [{addon_id: addon_id, quantity: 0}], amount: 0}' \
  --allowed-payment-method-type ach \
  --billing-address '{country: AF, city: city, state: state, street: street, zipcode: zipcode}' \
  --billing-currency AED \
  --confirm \
  --customer '{customer_id: customer_id}' \
  --customization '{force_language: force_language, show_on_demand_tag: true, show_order_details: true, theme: dark}' \
  --discount-code discount_code \
  --feature-flags '{allow_currency_selection: true, allow_customer_editing_city: true, allow_customer_editing_country: true, allow_customer_editing_email: true, allow_customer_editing_name: true, allow_customer_editing_state: true, allow_customer_editing_street: true, allow_customer_editing_zipcode: true, allow_discount_code: true, allow_phone_number_collection: true, allow_tax_id: true, always_create_new_customer: true, redirect_immediately: true}' \
  --force-3ds \
  --metadata '{foo: string}' \
  --minimal-address \
  --payment-method-id payment_method_id \
  --return-url return_url \
  --short-link \
  --show-saved-payment-methods \
  --subscription-data '{on_demand: {mandate_only: true, adaptive_currency_fees_inclusive: true, product_currency: AED, product_description: product_description, product_price: 0}, trial_period_days: 0}'
```

For details about specific commands, use the `--help` flag.

## Global Flags

- `--help` - Show command line usage
- `--debug` - Enable debug logging (includes HTTP request/response details)
- `--version`, `-v` - Show the CLI version
- `--base-url` - Use a custom API backend URL
- `--format` - Change the output format (`auto`, `explore`, `json`, `jsonl`, `pretty`, `raw`, `yaml`)
- `--format-error` - Change the output format for errors (`auto`, `explore`, `json`, `jsonl`, `pretty`, `raw`, `yaml`)
- `--transform` - Transform the data output using [GJSON syntax](https://github.com/tidwall/gjson/blob/master/SYNTAX.md)
- `--transform-error` - Transform the error output using [GJSON syntax](https://github.com/tidwall/gjson/blob/master/SYNTAX.md)
