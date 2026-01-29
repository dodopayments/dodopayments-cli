# Dodo Payments CLI

A powerful Command Line Interface for [Dodo Payments](https://dodopayments.com/), allowing developers to manage their Dodo Payments resources and test webhooks directly from the terminal.

### Setup
We provide various ways to install the CLI:  
Note: If you have Node or Bun installed, it's highly recommended to use that installation method.

Use NPM:
```bash
npm install -g dodopayments-cli
```

Or

Use Bun:
```bash
bun install -g dodopayments-cli
```

Or

Install manually:
- Download the latest release from [GitHub](https://github.com/dodopayments/dodopayments-cli/releases) that matches your system.

- Extract the downloaded file to a directory of your choice and open a terminal in that directory.

- Rename the binary file to `dodo`:  
Linux: `mv ./dodo-cli-* ./dodo`  
Windows: `ren .\dodo-cli-* .\dodo`

- Move the binary file to a directory in your system's PATH environment variable.
On Linux, you can use `sudo mv ./dodo /usr/local/bin/`.
On Windows, you can use `move .\dodo C:\Windows\System32\dodo` (This may require admin mode).

## Authentication

Before using the CLI, you must authenticate.

```bash
dodo login
```

This command will:
1. Open your browser to the Dodo Payments API Keys page.
2. Prompt you to enter your API Key.
3. Ask you to select the environment (**Test Mode** or **Live Mode**).
4. Will store your credentials (API Key) locally to `~/.dodopayments/api-key`.

## Usage

The general syntax is:

```bash
dodo <category> <sub-command>
```

### Products

Manage your products catalog.

- **List Products**:
  ```bash
  dodo products list
  ```
- **Create Product**:
  Opens the Dodo Payments dashboard to create a new product.
  ```bash
  dodo products create
  ```
- **Get Product Info**:
  View details for a specific product.
  ```bash
  dodo products info
  ```

### Payments

View payment transactions.

- **List Payments**:
  ```bash
  dodo payments list
  ```
- **Payment Info**:
  Get information about a specific payment.
  ```bash
  dodo payments info
  ```

### Customers

Manage your customer base.

- **List Customers**:
  ```bash
  dodo customers list
  ```
- **Create Customer**:
  Create a new customer profile.
  ```bash
  dodo customers create
  ```
- **Update Customer**:
  Update an existing customer's details.
  ```bash
  dodo customers update
  ```

### Discounts

Manage coupons and discounts.

- **List Discounts**:
  ```bash
  dodo discounts list
  ```
- **Create Discount**:
  Create a new percentage-based discount.
  ```bash
  dodo discounts create
  ```
- **Delete Discount**:
  Remove a discount by ID.
  ```bash
  dodo discounts delete
  ```

### Licenses

Manage software licenses.

- **List Licenses**:
  ```bash
  dodo licences list
  ```

### Webhook Testing

The CLI includes a robust tool for testing webhooks by simulating events.

Note: The webhook testing tool doesn't support signing requests yet. Please disable webhook signature verification while testing. A simple way of doing this would be using `unsafe_unwrap()` instead of just `unwrap` in the webhook endpoint DURING TESTING ONLY.

```bash
dodo wh
```

This interactive tool guides you through:
1. Setting a destination **endpoint URL**.
2. Configuring **Business ID**, **Product ID**, and **Metadata** (Interactive prompts).
3. Selecting a specific **Event** to trigger.

**Supported Webhook Events:**
- **Subscription**: `active`, `updated`, `on_hold`, `renewed`, `plan_changed`, `cancelled`, `failed`, `expired`
- **Payment**: `success`, `failed`, `processing`, `cancelled`
- **Refund**: `success`, `failed`
- **Dispute**: `opened`, `expired`, `accepted`, `cancelled`, `challenged`, `won`, `lost`
- **Licence**: `created`