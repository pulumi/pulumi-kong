// Copyright 2016-2018, Pulumi Corporation.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package kong

import (
	"unicode"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/kevholditch/terraform-provider-kong/kong"
	"github.com/pulumi/pulumi-terraform-bridge/v2/pkg/tfbridge"
	"github.com/pulumi/pulumi/sdk/v2/go/common/tokens"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "kong"
	// modules:
	mainMod = "index" // the y module
)

// makeMember manufactures a type token for the package and the given module and type.
func makeMember(mod string, mem string) tokens.ModuleMember {
	return tokens.ModuleMember(mainPkg + ":" + mod + ":" + mem)
}

// makeType manufactures a type token for the package and the given module and type.
func makeType(mod string, typ string) tokens.Type {
	return tokens.Type(makeMember(mod, typ))
}

// makeResource manufactures a standard resource token given a module and resource name.  It
// automatically uses the main package and names the file by simply lower casing the resource's
// first character.
func makeResource(mod string, res string) tokens.Type {
	fn := string(unicode.ToLower(rune(res[0]))) + res[1:]
	return makeType(mod+"/"+fn, res)
}

func refProviderLicense(license tfbridge.TFProviderLicense) *tfbridge.TFProviderLicense {
	return &license
}

func Provider() tfbridge.ProviderInfo {
	p := kong.Provider().(*schema.Provider)
	prov := tfbridge.ProviderInfo{
		P:                 p,
		Name:              "kong",
		Description:       "A Pulumi package for creating and managing Kong resources.",
		Keywords:          []string{"pulumi", "kong"},
		GitHubOrg:         "kevholditch",
		License:           "Apache-2.0",
		Homepage:          "https://pulumi.io",
		Repository:        "https://github.com/pulumi/pulumi-kong",
		TFProviderLicense: refProviderLicense(tfbridge.MITLicenseType),
		Config: map[string]*tfbridge.SchemaInfo{
			"kong_admin_uri": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"KONG_ADMIN_ADDR"},
					Value:   "http://localhost:8001",
				},
			},
			"kong_admin_username": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"KONG_ADMIN_USERNAME"},
				},
			},
			"kong_admin_password": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"KONG_ADMIN_PASSWORD"},
				},
			},
			"tls_skip_verify": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"TLS_SKIP_VERIFY"},
					Value:   false,
				},
			},
			"kong_api_key": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"KONG_API_KEY"},
				},
			},
			"kong_admin_token": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"KONG_ADMIN_TOKEN"},
				},
			},
			"strict_plugins_match": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"STRICT_PLUGINS_MATCH"},
				},
			},
		},
		Resources: map[string]*tfbridge.ResourceInfo{
			"kong_certificate": {
				Tok: makeResource(mainMod, "Certificate"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"certificate": {
						CSharpName: "Cert",
					},
				},
			},
			"kong_consumer":               {Tok: makeResource(mainMod, "Consumer")},
			"kong_consumer_plugin_config": {Tok: makeResource(mainMod, "ConsumerPluginConfig")},
			"kong_plugin":                 {Tok: makeResource(mainMod, "Plugin")},
			"kong_sni":                    {Tok: makeResource(mainMod, "Sni")},
			"kong_upstream":               {Tok: makeResource(mainMod, "Upstream")},
			"kong_target": {
				Tok: makeResource(mainMod, "Target"),
				Fields: map[string]*tfbridge.SchemaInfo{
					"target": {
						CSharpName: "TargetAddress",
					},
				},
			},
			"kong_service": {Tok: makeResource(mainMod, "Service")},
			"kong_route":   {Tok: makeResource(mainMod, "Route")},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^2.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^8.0.25",
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=2.9.0,<3.0.0",
			},
			UsesIOClasses: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi":                       "2.*",
				"System.Collections.Immutable": "1.6.0",
			},
			Namespaces: map[string]string{
				mainPkg: "Kong",
			},
		},
	}

	// For all resources with name properties, we will add an auto-name property.  Make sure to skip those that
	// already have a name mapping entry, since those may have custom overrides set above (e.g., for length).
	const nameProperty = "name"
	for resname, res := range prov.Resources {
		if schema := p.ResourcesMap[resname]; schema != nil {
			// Only apply auto-name to input properties (Optional || Required) named `name`
			if tfs, has := schema.Schema[nameProperty]; has && (tfs.Optional || tfs.Required) {
				if _, hasfield := res.Fields[nameProperty]; !hasfield {
					if res.Fields == nil {
						res.Fields = make(map[string]*tfbridge.SchemaInfo)
					}
					res.Fields[nameProperty] = tfbridge.AutoName(nameProperty, 255)
				}
			}
		}
	}

	return prov
}
