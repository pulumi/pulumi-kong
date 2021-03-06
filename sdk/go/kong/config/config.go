// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package config

import (
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi/config"
)

// An basic auth password for kong admin
func GetKongAdminPassword(ctx *pulumi.Context) string {
	return config.Get(ctx, "kong:kongAdminPassword")
}

// API key for the kong api (Enterprise Edition)
func GetKongAdminToken(ctx *pulumi.Context) string {
	return config.Get(ctx, "kong:kongAdminToken")
}

// The address of the kong admin url e.g. http://localhost:8001
func GetKongAdminUri(ctx *pulumi.Context) string {
	return config.Get(ctx, "kong:kongAdminUri")
}

// An basic auth user for kong admin
func GetKongAdminUsername(ctx *pulumi.Context) string {
	return config.Get(ctx, "kong:kongAdminUsername")
}

// API key for the kong api (if you have locked it down)
func GetKongApiKey(ctx *pulumi.Context) string {
	return config.Get(ctx, "kong:kongApiKey")
}

// Should plugins `config_json` field strictly match plugin configuration
func GetStrictPluginsMatch(ctx *pulumi.Context) bool {
	v, err := config.TryBool(ctx, "kong:strictPluginsMatch")
	if err == nil {
		return v
	}
	return getEnvOrDefault(false, parseEnvBool, "STRICT_PLUGINS_MATCH").(bool)
}

// Whether to skip tls verify for https kong api endpoint using self signed or untrusted certs
func GetTlsSkipVerify(ctx *pulumi.Context) bool {
	v, err := config.TryBool(ctx, "kong:tlsSkipVerify")
	if err == nil {
		return v
	}
	return getEnvOrDefault(false, parseEnvBool, "TLS_SKIP_VERIFY").(bool)
}
