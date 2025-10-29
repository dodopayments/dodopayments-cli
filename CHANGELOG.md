# Changelog

## 1.56.3 (2025-10-29)

Full Changelog: [v1.56.0...v1.56.3](https://github.com/dodopayments/dodopayments-cli/compare/v1.56.0...v1.56.3)

### Chores

* **internal:** codegen related update ([07d4b67](https://github.com/dodopayments/dodopayments-cli/commit/07d4b6775c58db08bfe69457d44862af4827ffa2))

## 1.56.0 (2025-10-27)

Full Changelog: [v1.55.8...v1.56.0](https://github.com/dodopayments/dodopayments-cli/compare/v1.55.8...v1.56.0)

### Features

* **api:** updated to openapi spec v1.56.0 ([750617d](https://github.com/dodopayments/dodopayments-cli/commit/750617d6a176bea28432e533a9379f526aa11d4e))

## 1.55.8 (2025-10-25)

Full Changelog: [v1.55.7...v1.55.8](https://github.com/dodopayments/dodopayments-cli/compare/v1.55.7...v1.55.8)

### Bug Fixes

* fix builds for non-public Go repos ([a768b8e](https://github.com/dodopayments/dodopayments-cli/commit/a768b8e638910f2b4e5bd04af87cfe5242b58f1b))
* remove some bootstrapping logic ([efeec41](https://github.com/dodopayments/dodopayments-cli/commit/efeec413ca127340a6eeb6144f7899b0e0b9c9f5))

## 1.55.7 (2025-10-18)

Full Changelog: [v1.55.0...v1.55.7](https://github.com/dodopayments/dodopayments-cli/compare/v1.55.0...v1.55.7)

### Features

* **api:** updates for openapi spec v1.55.7 ([de77b0c](https://github.com/dodopayments/dodopayments-cli/commit/de77b0c715ac2e89db997e0380c1c26c593a3980))


### Chores

* **internal:** codegen related update ([b69bd93](https://github.com/dodopayments/dodopayments-cli/commit/b69bd93cb95d10629af72ba2784761fb6e1e97d4))

## 1.55.0 (2025-10-16)

Full Changelog: [v1.54.0...v1.55.0](https://github.com/dodopayments/dodopayments-cli/compare/v1.54.0...v1.55.0)

### Features

* **api:** updated openapi spec to v1.55.0 ([f08040b](https://github.com/dodopayments/dodopayments-cli/commit/f08040b4c1872f0ce7992ac334c6ed6bd0822a88))
* arguments now have defaults and descriptions ([d25f0b3](https://github.com/dodopayments/dodopayments-cli/commit/d25f0b3023e4a4ba06ee652d7b3dae3f9be0cccd))


### Bug Fixes

* pass through context parameter correctly ([3c02e4d](https://github.com/dodopayments/dodopayments-cli/commit/3c02e4d684f6bb9f35211be1e610bf2f8c95d26a))


### Chores

* **internal:** codegen related update ([0b971d2](https://github.com/dodopayments/dodopayments-cli/commit/0b971d2ae0220cb121fb3702b46826a2d173de55))
* **internal:** codegen related update ([df5f77a](https://github.com/dodopayments/dodopayments-cli/commit/df5f77a9e5516ed535fb57ec053781e3dcc5242f))

## 1.54.0 (2025-10-07)

Full Changelog: [v1.53.7...v1.54.0](https://github.com/dodopayments/dodopayments-cli/compare/v1.53.7...v1.54.0)

### Features

* added `--output-filter` flag and `--error-format` flag to support better visualization options ([f77dd2a](https://github.com/dodopayments/dodopayments-cli/commit/f77dd2a54ad060dcd70b3494cf9e8fb8eee28e2a))
* better support for positional arguments ([63750f1](https://github.com/dodopayments/dodopayments-cli/commit/63750f1b1030bd83feda023df2a23437af2b6c0e))


### Bug Fixes

* downgrade urfave/cli-docs dependency ([d439477](https://github.com/dodopayments/dodopayments-cli/commit/d43947799d18935a86f421fe12439b44443c5db1))


### Chores

* bump Go version ([ad01d72](https://github.com/dodopayments/dodopayments-cli/commit/ad01d7286fe09bde2e01d05755f20c438ee24c1f))
* **internal:** codegen related update ([7dcc932](https://github.com/dodopayments/dodopayments-cli/commit/7dcc93274b63beaa5cd37be044d969b120961d56))
* **internal:** codegen related update ([5859fa2](https://github.com/dodopayments/dodopayments-cli/commit/5859fa256db4aab6de6a19a9798be1102298ea5a))

## 1.53.7 (2025-09-25)

Full Changelog: [v1.53.6...v1.53.7](https://github.com/dodopayments/dodopayments-cli/compare/v1.53.6...v1.53.7)

### Chores

* **internal:** codegen related update ([efa1517](https://github.com/dodopayments/dodopayments-cli/commit/efa1517ccab51949bf2b664414182a1e67e0db35))

## 1.53.6 (2025-09-20)

Full Changelog: [v1.53.5...v1.53.6](https://github.com/dodopayments/dodopayments-cli/compare/v1.53.5...v1.53.6)

### Chores

* do not install brew dependencies in ./scripts/bootstrap by default ([eba3565](https://github.com/dodopayments/dodopayments-cli/commit/eba3565dc10ac51c0f1969e3cf79728cd05f6c5c))

## 1.53.5 (2025-09-16)

Full Changelog: [v1.53.4...v1.53.5](https://github.com/dodopayments/dodopayments-cli/compare/v1.53.4...v1.53.5)

### Bug Fixes

* fix for issue with nil responses ([1fd332d](https://github.com/dodopayments/dodopayments-cli/commit/1fd332d21e1eedad9dc5cccafd536afa2a78dc18))


### Chores

* code cleanup for `interface{}` ([1e0c3bd](https://github.com/dodopayments/dodopayments-cli/commit/1e0c3bd5d62cbad22d05cf0a803191a53ac6ef8b))

## 1.53.4 (2025-09-15)

Full Changelog: [v1.53.3...v1.53.4](https://github.com/dodopayments/dodopayments-cli/compare/v1.53.3...v1.53.4)

### Features

* **api:** added ghcr deployment for mcp ([539873f](https://github.com/dodopayments/dodopayments-cli/commit/539873fb75f09a962008030dacf80a6d797b35f9))

## 1.53.3 (2025-09-13)

Full Changelog: [v1.53.2...v1.53.3](https://github.com/dodopayments/dodopayments-cli/compare/v1.53.2...v1.53.3)

### Features

* **api:** added typescript sdk for migration and updated org info ([33e8156](https://github.com/dodopayments/dodopayments-cli/commit/33e8156c308e2a7c4da1dc25c3e84e342da6e140))
* **api:** manual updates ([5ed95d6](https://github.com/dodopayments/dodopayments-cli/commit/5ed95d6e4dc7e6ba8be0fda765107ee7520c269f))

## 1.53.2 (2025-09-13)

Full Changelog: [v0.1.0...v1.53.2](https://github.com/dodopayments/dodopayments-cli/compare/v0.1.0...v1.53.2)

### Features

* **api:** updated openapi spec to v1.53.2 with customer credits. ([c2af955](https://github.com/dodopayments/dodopayments-cli/commit/c2af955c40e46d280703d1bd6f0f1b4e31916e8c))

## 0.1.0 (2025-09-13)

Full Changelog: [v0.0.1...v0.1.0](https://github.com/dodopayments/dodopayments-cli/compare/v0.0.1...v0.1.0)

### Features

* **api:** manual updates ([c38a05c](https://github.com/dodopayments/dodopayments-cli/commit/c38a05ca21712da1d9cfc9fb1903cbbecc191535))


### Chores

* configure new SDK language ([23560c7](https://github.com/dodopayments/dodopayments-cli/commit/23560c714804729994025c02108484ba8287e10b))
* update SDK settings ([7b6afa3](https://github.com/dodopayments/dodopayments-cli/commit/7b6afa382acdc591534e11627fde07782dcdbf5b))
