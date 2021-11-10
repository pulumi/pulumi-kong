[![Actions Status](https://github.com/pulumi/pulumi-kong/workflows/master/badge.svg)](https://github.com/pulumi/pulumi-kong/actions)
[![Slack](http://www.pulumi.com/images/docs/badges/slack.svg)](https://slack.pulumi.com)
[![NPM version](https://badge.fury.io/js/%40pulumi%2Fkong.svg)](https://www.npmjs.com/package/@pulumi/kong)
[![Python version](https://badge.fury.io/py/pulumi-kong.svg)](https://pypi.org/project/pulumi-kong)
[![NuGet version](https://badge.fury.io/nu/pulumi.kong.svg)](https://badge.fury.io/nu/pulumi.kong)
[![PkgGoDev](https://pkg.go.dev/badge/github.com/pulumi/pulumi-kong/sdk/v3/go)](https://pkg.go.dev/github.com/pulumi/pulumi-kong/sdk/v3/go)
[![License](https://img.shields.io/npm/l/%40pulumi%2Fpulumi.svg)](https://github.com/pulumi/pulumi-kong/blob/master/LICENSE)

# Kong Resource Provider

The Kong resource provider for Pulumi lets you manage Kong resources in your cloud programs. To use
this package, please [install the Pulumi CLI first](https://www.mailgun.com//).

## Installing

This package is available in many languages in the standard packaging formats.

### Node.js (Java/TypeScript)

To use from JavaScript or TypeScript in Node.js, install using either `npm`:

    $ npm install @pulumi/kong

or `yarn`:

    $ yarn add @pulumi/kong

### Python

To use from Python, install using `pip`:

    $ pip install pulumi_kong

### Go

To use from Go, use `go get` to grab the latest version of the library

    $ go get github.com/pulumi/pulumi-kong/sdk/v4

### .NET

To use from .NET, install using `dotnet add package`:

    $ dotnet add package Pulumi.Kong

## Configuration

The following configuration points are available:

- `kong:kongAdminUri` - The url of the kong admin api. May be set via the `KONG_ADMIN_ADDR` environment variable. Defaults to `http://localhost:8001`.
- `kong:kongAdminUsername` - Username for the kong admin api. May be set via the `KONG_ADMIN_USERNAME` environment variable.
- `kong:kongAdminPassword` - Password for the kong admin api. May be set via the `KONG_ADMIN_PASSWORD` environment variable.
- `kong:tlsSkipVerify` - Whether to skip tls certificate verification for the kong api when using https. May be set via the `TLS_SKIP_VERIFY` environment variable. Defaults to `false`.
- `kong:kongApiKey` - API key used to secure the kong admin API. May be set via the `KONG_API_KEY` environment variable.
- `kong:kongAdminToken` - API key used to secure the kong admin API in the Enterprise Edition. May be set via the `KONG_ADMIN_TOKEN` environment variable.
- `kong:strictPluginsMatch` - Should plugins `config_json` field strictly match plugin configuration. May be set via the `STRICT_PLUGINS_MATCH` environment variable. Defaults to `false`.

## Reference

For further information, please visit [the Kong provider docs](https://www.pulumi.com/docs/intro/cloud-providers/kong) or for detailed reference documentation, please visit [the API docs](https://www.pulumi.com/docs/reference/pkg/kong).
