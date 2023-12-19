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
	_ "embed" // allow embedding metadata
	"fmt"
	"path"

	"github.com/kevholditch/terraform-provider-kong/kong"
	"github.com/pulumi/pulumi-kong/provider/v4/pkg/version"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge"
	"github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfbridge/tokens"
	shimv2 "github.com/pulumi/pulumi-terraform-bridge/v3/pkg/tfshim/sdk-v2"
)

// all of the token components used below.
const (
	// packages:
	mainPkg = "kong"
	// modules:
	mainMod = "index" // the y module
)

func ref[T any](t T) *T { return &t }

//go:embed cmd/pulumi-resource-kong/bridge-metadata.json
var metadata []byte

func Provider() tfbridge.ProviderInfo {
	prov := tfbridge.ProviderInfo{
		P:                 shimv2.NewProvider(kong.Provider()),
		Name:              "kong",
		Description:       "A Pulumi package for creating and managing Kong resources.",
		Keywords:          []string{"pulumi", "kong"},
		GitHubOrg:         "kevholditch",
		License:           "Apache-2.0",
		Homepage:          "https://pulumi.io",
		Repository:        "https://github.com/pulumi/pulumi-kong",
		TFProviderLicense: ref(tfbridge.MITLicenseType),
		Version:           version.Version,
		MetadataInfo:      tfbridge.NewProviderMetadata(metadata),
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
				Fields: map[string]*tfbridge.SchemaInfo{
					"certificate": {
						CSharpName: "Cert",
					},
				},
			},
			"kong_target": {
				Fields: map[string]*tfbridge.SchemaInfo{
					"target": {
						CSharpName: "TargetAddress",
					},
				},
			},
		},
		JavaScript: &tfbridge.JavaScriptInfo{
			Dependencies: map[string]string{
				"@pulumi/pulumi": "^3.0.0",
			},
			DevDependencies: map[string]string{
				"@types/node": "^10.0.0",
				"@types/mime": "^2.0.0",
			},
		},
		Python: &tfbridge.PythonInfo{
			Requires: map[string]string{
				"pulumi": ">=3.0.0,<4.0.0",
			},
			PyProject: struct{ Enabled bool }{true},
		},

		Golang: &tfbridge.GolangInfo{
			ImportBasePath: path.Join(
				fmt.Sprintf("github.com/pulumi/pulumi-%[1]s/sdk/", mainPkg),
				tfbridge.GetModuleMajorVersion(version.Version),
				"go",
				mainPkg,
			),
			GenerateResourceContainerTypes: true,
		},
		CSharp: &tfbridge.CSharpInfo{
			PackageReferences: map[string]string{
				"Pulumi": "3.*",
			},
			Namespaces: map[string]string{
				mainPkg: "Kong",
			},
		},
	}

	prov.MustComputeTokens(tokens.SingleModule("kong_", mainMod, tokens.MakeStandard(mainPkg)))

	prov.SetAutonaming(255, "-")

	prov.MustApplyAutoAliases()

	return prov
}
