# Security Policy

## Supported Versions

We actively support the following versions with security updates:

| Version | Supported          |
| ------- | ------------------ |
| 0.1.x   | :white_check_mark: |

## Reporting a Vulnerability

We take the security of Dodo Payments CLI seriously. If you believe you have found a security vulnerability, please report it to us responsibly.

### How to Report

**Please do NOT report security vulnerabilities through public GitHub issues.**

Instead, please report them via one of the following methods:

1. **GitHub Security Advisories** (Preferred): [Create a security advisory](https://github.com/dodopayments/dodopayments-cli/security/advisories/new)

2. **Email**: Send an email to [security@dodopayments.com](mailto:security@dodopayments.com)

### What to Include

Please include the following information in your report:

- Type of vulnerability (e.g., credential exposure, code injection, etc.)
- Full paths of source file(s) related to the vulnerability
- The location of the affected source code (tag/branch/commit or direct URL)
- Any special configuration required to reproduce the issue
- Step-by-step instructions to reproduce the issue
- Proof-of-concept or exploit code (if possible)
- Impact of the issue, including how an attacker might exploit it

### What to Expect

- **Acknowledgment**: We will acknowledge receipt of your vulnerability report within 48 hours.

- **Communication**: We will keep you informed of our progress throughout the process.

- **Resolution Timeline**: We aim to investigate and address security issues within 90 days, depending on complexity.

- **Disclosure**: We will coordinate with you on the timing of public disclosure.

- **Credit**: We will credit you for the discovery in our release notes and security advisory (unless you prefer to remain anonymous).

## Security Best Practices for Users

When using Dodo Payments CLI, please follow these security best practices:

### API Keys

- **Never share your API keys** with anyone
- **Use test mode keys** for development and testing
- **Rotate keys** if you suspect they may have been compromised
- The CLI stores your API key locally at `~/.dodopayments/api-key` - ensure this file has appropriate permissions

### Running the CLI

- **Download only from official sources**: npm, GitHub releases, or bun
- **Verify the package** before installation when possible
- **Keep the CLI updated** to the latest version

### Webhook Testing

- **Only test webhooks** against endpoints you control
- **Use test mode** for webhook testing when possible
- **Re-enable signature verification** after testing

## Local Storage

The CLI stores the following data locally:

| Data | Location | Purpose |
|------|----------|---------|
| API Key | `~/.dodopayments/api-key` | Authentication |

Ensure your home directory has appropriate access controls to protect this sensitive data.

## Vulnerability Disclosure Policy

We follow a coordinated vulnerability disclosure process:

1. Reporter submits vulnerability privately
2. We acknowledge and begin investigation
3. We develop and test a fix
4. We release the fix and publish a security advisory
5. We credit the reporter (if desired)

We ask that you:

- Give us reasonable time to address the issue before public disclosure
- Make a good faith effort to avoid privacy violations, destruction of data, and interruption of services
- Do not access or modify data that does not belong to you

Thank you for helping keep Dodo Payments CLI and our users safe!
