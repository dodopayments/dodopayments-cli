# Contributing to Dodo Payments CLI

First off, thank you for considering contributing to the Dodo Payments CLI! It's people like you that make this tool great for everyone.

We welcome contributions of all kinds - bug reports, feature requests, documentation improvements, and code contributions.

## Table of Contents

- [Code of Conduct](#code-of-conduct)
- [Getting Started](#getting-started)
- [Development Setup](#development-setup)
- [Project Structure](#project-structure)
- [How to Contribute](#how-to-contribute)
- [Coding Guidelines](#coding-guidelines)
- [Pull Request Process](#pull-request-process)
- [Adding New Commands](#adding-new-commands)
- [Community](#community)

## Code of Conduct

This project and everyone participating in it is governed by our [Code of Conduct](./CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code. Please report unacceptable behavior to [support@dodopayments.com](mailto:support@dodopayments.com).

## Getting Started

### Prerequisites

- [Bun](https://bun.sh/) >= 1.0 (recommended) or Node.js >= 18
- A Dodo Payments account (for testing)

### Development Setup

1. **Fork the repository** on GitHub

2. **Clone your fork** locally:
   ```bash
   git clone https://github.com/YOUR_USERNAME/dodopayments-cli.git
   cd dodopayments-cli
   ```

3. **Install dependencies**:
   ```bash
   bun install
   ```

4. **Create a branch** for your changes:
   ```bash
   git checkout -b feature/your-feature-name
   ```

5. **Run the CLI locally**:
   ```bash
   bun run index.ts <command>
   ```

6. **Build the project**:
   ```bash
   bun run build
   ```

## Project Structure

```
dodopayments-cli/
├── index.ts              # Main CLI entrypoint
├── utils/                # Utility functions
├── dodo-webhooks/        # Webhook testing functionality
├── .github/
│   ├── ISSUE_TEMPLATE/   # Issue templates
│   ├── PULL_REQUEST_TEMPLATE.md
│   └── workflows/        # CI/CD workflows
├── CONTRIBUTING.md       # This file
├── CODE_OF_CONDUCT.md
├── LICENSE
└── README.md
```

## How to Contribute

### Reporting Bugs

Before creating a bug report, please check existing issues to avoid duplicates. When creating a bug report, include:

- A clear, descriptive title
- Steps to reproduce the issue
- Expected behavior vs actual behavior
- Your environment (OS, Bun/Node version, CLI version)
- Any relevant error messages or logs

### Suggesting Features

Feature requests are welcome! Please include:

- A clear description of the feature
- The problem it solves or use case it addresses
- Any alternative solutions you've considered

### Contributing Code

1. Check the [issues](https://github.com/dodopayments/dodopayments-cli/issues) for open tasks
2. Comment on an issue to let others know you're working on it
3. Follow the [coding guidelines](#coding-guidelines)
4. Submit a [pull request](#pull-request-process)

## Coding Guidelines

### General Principles

- **Use TypeScript**: All code should be written in TypeScript
- **Use Bun APIs**: Prefer Bun's built-in APIs when available
- **Keep it simple**: Write clear, readable code
- **Add comments**: Document complex logic and non-obvious decisions

### Code Style

- Use 2 spaces for indentation
- Use single quotes for strings
- Add semicolons at end of statements
- Use `const` by default, `let` when reassignment is needed

### CLI Design Principles

- **Interactive by default**: Use prompts for required inputs when not provided as arguments
- **Helpful error messages**: Provide clear, actionable error messages
- **Consistent output**: Use consistent formatting across all commands

## Pull Request Process

1. **Update documentation** if you're adding or changing functionality

2. **Follow the PR template** - Fill out all sections completely

3. **Link related issues** - Reference any issues your PR addresses

4. **Keep PRs focused** - One feature or fix per PR

5. **Write clear commit messages** - Describe what and why

6. **Ensure the build passes** - Run `bun run build` before submitting

7. **Request review** - Tag maintainers for review

## Adding New Commands

To add a new command to the CLI:

1. **Add the command handler** in `index.ts`:
   ```typescript
   case 'your-command':
     await handleYourCommand();
     break;
   ```

2. **Implement the handler function**:
   - Use `inquirer` for interactive prompts
   - Use the `dodopayments` SDK for API calls
   - Handle errors gracefully with helpful messages

3. **Update the README** with documentation for the new command

4. **Test thoroughly** in both test and live modes

### Example Command Structure

```typescript
async function handleYourCommand() {
  // 1. Parse any arguments
  const args = process.argv.slice(3);
  
  // 2. Prompt for missing inputs
  const answers = await inquirer.prompt([
    {
      type: 'input',
      name: 'someInput',
      message: 'Enter value:',
    }
  ]);
  
  // 3. Make API call
  const result = await dodo.someResource.someMethod({
    // ...
  });
  
  // 4. Display results
  console.log(result);
}
```

## Community

- **Discord**: [Join our community](https://discord.gg/bYqAp4ayYh) for help and discussion
- **GitHub Issues**: For bug reports and feature requests
- **Twitter**: Follow [@dodopayments](https://twitter.com/dodopayments) for updates

## Recognition

We appreciate all contributors! Your contributions help make the Dodo Payments CLI better for everyone.

---

Thank you for contributing to Dodo Payments CLI!
