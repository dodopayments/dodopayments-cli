# Dodo Payments CLI

The official CLI for the [Dodo Payments REST API](https://docs.dodopayments.com).

It is generated with [Stainless](https://www.stainless.com/).

## Installation

### Installing with Go

```sh
go install 'github.com/stainless-sdks/dodo-payments-cli/cmd/dodo-payments@latest'
```

### Running Locally

```sh
go run cmd/dodo-payments/main.go
```

## Usage

The CLI follows a resource-based command structure:

```sh
dodo-payments [resource] [command] [flags]
```

```sh
dodo-payments checkout-sessions create \
  --product-cart.product_id product_id \
  --product-cart.quantity 0
```

For details about specific commands, use the `--help` flag.

## Global Flags

- `--debug` - Enable debug logging (includes HTTP request/response details)
- `--version`, `-v` - Show the CLI version
