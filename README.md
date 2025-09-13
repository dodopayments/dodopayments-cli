# Dodo Payments CLI

The official CLI for the [Dodo Payments REST API](https://docs.dodopayments.com).

It is generated with [Stainless](https://www.stainless.com/).

## Installation

### Installing with Go

<!-- x-release-please-start-version -->

```sh
go install 'github.com/dodopayments/dodopayments-cli/cmd/dodopayments@latest'
```

### Running Locally

<!-- x-release-please-start-version -->

```sh
go run cmd/dodopayments/main.go
```

<!-- x-release-please-end -->

## Usage

The CLI follows a resource-based command structure:

```sh
dodopayments [resource] [command] [flags]
```

```sh
dodopayments checkout-sessions create \
  --product-cart.product_id product_id \
  --product-cart.quantity 0
```

For details about specific commands, use the `--help` flag.

## Global Flags

- `--debug` - Enable debug logging (includes HTTP request/response details)
- `--version`, `-v` - Show the CLI version
