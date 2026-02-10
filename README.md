# Dodo Payments CLI

<p align="center">
  <strong>A powerful Command Line Interface for Dodo Payments</strong>
</p>

<p align="center">
  <a href="https://www.npmjs.com/package/dodopayments-cli">
    <img src="https://img.shields.io/npm/v/dodopayments-cli?color=cb3837&label=npm&logo=npm" alt="npm version" />
  </a>
  <a href="https://discord.gg/bYqAp4ayYh">
    <img src="https://img.shields.io/discord/1305511580854779984?label=Discord&logo=discord&color=5865F2" alt="Discord" />
  </a>
  <a href="LICENSE">
    <img src="https://img.shields.io/badge/license-MIT-blue.svg" alt="License: MIT" />
  </a>
</p>

<p align="center">
  <a href="#installation">Installation</a> •
  <a href="#authentication">Authentication</a> •
  <a href="#usage">Usage</a> •
  <a href="#webhook-testing">Webhooks</a> •
  <a href="#contributing">Contributing</a>
</p>

---

Manage your [Dodo Payments](https://dodopayments.com/) resources and test webhooks directly from the terminal. Built for developers who prefer the command line.

## Installation

We provide various ways to install the CLI:

> **Note:** If you have Node or Bun installed, it's highly recommended to use that installation method.

### Using npm (Recommended)

```bash
npm install -g dodopayments-cli
```

### Using Bun

```bash
bun install -g dodopayments-cli
```

### Manual Installation

1. Download the latest release from [GitHub Releases](https://github.com/dodopayments/dodopayments-cli/releases) that matches your system.

2. Extract the downloaded file to a directory of your choice.

3. Rename the binary file to `dodo`:
   - **Linux/macOS:** `mv ./dodopayments-cli-* ./dodo`
   - **Windows:** `ren .\dodopayments-cli-* .\dodo`

4. Move the binary to a directory in your system's PATH:
   - **Linux/macOS:** `sudo mv ./dodo /usr/local/bin/`
   - **Windows:** `move .\dodo C:\Windows\System32\dodo` (requires admin mode)

## Authentication

Before using the CLI, you must authenticate:

```bash
dodo login
```

This command will:
1. Open your browser to the Dodo Payments API Keys page
2. Prompt you to enter your API Key
3. Ask you to select the environment (**Test Mode** or **Live Mode**)
4. Store your credentials locally to `~/.dodopayments/api-key`

## Usage

The general syntax is:

```bash
dodo <category> <sub-command>
```

### Products

Manage your products catalog.

| Command | Description |
|---------|-------------|
| `dodo products list` | List all products |
| `dodo products create` | Open dashboard to create a product |
| `dodo products info` | View details for a specific product |

### Payments

View payment transactions.

| Command | Description |
|---------|-------------|
| `dodo payments list` | List all payments |
| `dodo payments info` | Get information about a specific payment |

### Customers

Manage your customer base.

| Command | Description |
|---------|-------------|
| `dodo customers list` | List all customers |
| `dodo customers create` | Create a new customer profile |
| `dodo customers update` | Update an existing customer's details |

### Discounts

Manage coupons and discounts.

| Command | Description |
|---------|-------------|
| `dodo discounts list` | List all discounts |
| `dodo discounts create` | Create a new percentage-based discount |
| `dodo discounts delete` | Remove a discount by ID |

### Licenses

Manage software licenses.

| Command | Description |
|---------|-------------|
| `dodo licences list` | List all licenses |

### Addons

Manage your add-ons.

| Command | Description |
|---------|-------------|
| `dodo addons list` | List all addons |
| `dodo addons create` | Open dashboard to create an addon |
| `dodo addons info` | View details for a specific addon |

### Refunds

Manage your refunds.

| Command | Description |
|---------|-------------|
| `dodo refund list` | List all refunds |
| `dodo refund info` | View details for a specific refund |

### Webhooks

Manage and test webhooks directly from the CLI.

| Command | Description |
|---------|-------------|
| `dodo wh listen` | Listen for webhooks from Dodo Payments in real time and forward them to your local dev server |
| `dodo wh trigger` | Trigger a test webhook event interactively |

> **Note:** The webhook triggering doesn't support signing requests yet. Please disable webhook signature verification while triggering. A simple way to do this is using `unsafe_unwrap()` instead of `unwrap()` in the webhook endpoint **during testing only**.

> **Note:** The webhook listening tool will only work with a test mode API key. If you use a live mode API key, it won't work.

This interactive tool guides you through:
1. Setting a destination **endpoint URL**
3. Selecting a specific **Event** to trigger

### Supported Webhook Events

| Category | Events |
|----------|--------|
| **Subscription** | `active`, `updated`, `on_hold`, `renewed`, `plan_changed`, `cancelled`, `failed`, `expired` |
| **Payment** | `success`, `failed`, `processing`, `cancelled` |
| **Refund** | `success`, `failed` |
| **Dispute** | `opened`, `expired`, `accepted`, `cancelled`, `challenged`, `won`, `lost` |
| **License** | `created` |

## Contributing

We welcome contributions from the community! Whether you're fixing bugs, adding new features, or improving documentation, your help is appreciated.

Please read our [Contributing Guide](./CONTRIBUTING.md) to get started. We also have a [Code of Conduct](./CODE_OF_CONDUCT.md) that we expect all contributors to follow.

### Ways to Contribute

- Report bugs and suggest features
- Improve documentation
- Add new CLI commands
- Write tests
- Review pull requests

## Support

- [Discord Community](https://discord.gg/bYqAp4ayYh) - Get help and discuss with the community
- [GitHub Issues](https://github.com/dodopayments/dodopayments-cli/issues) - Report bugs or request features
- [Documentation](https://docs.dodopayments.com) - Learn more about Dodo Payments

## License

This project is licensed under the MIT License - see the [LICENSE](./LICENSE) file for details.

---

<p align="center">
  Made with ❤️ by <a href="https://dodopayments.com">Dodo Payments</a>
</p>
