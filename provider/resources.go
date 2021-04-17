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
	"fmt"
	"path/filepath"
	"unicode"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/kevholditch/terraform-provider-kong/kong"
	"github.com/pulumi/pulumi-kong/provider/v3/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	shimv1 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v1"
	"github.com/pulumi/pulumi/sdk/v3/go/common/tokens"
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
	p := shimv1.NewProvider(kong.Provider().(*schema.Provider))
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
			"tls_skip_verify": {
				Default: &tfbridge.DefaultInfo{
					EnvVars: []string{"TLS_SKIP_VERIFY"},
					Value:   false,
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
				"@pulumi/pulumi": "^3.0.0-alpha.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0",
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=3.0.0a1,<4.0.0",
			},
		},
		Golang: &tfbridge.GolangInfo{
			ImportBasePath: filepath.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi":                       "3.*-*",
				"System.Collections.Immutable": "1.6.0",
			},
			Namespaces: map[string]string{
				mainPkg: "Kong",
			},
		},
	}

	prov.SetAutonaming(255, "-")

	return prov
}
