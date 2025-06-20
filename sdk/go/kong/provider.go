// Code generated by pulumi-language-go DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package kong

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi-kong/sdk/v4/go/kong/internal"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// The provider type for the kong package. By default, resources use package-wide configuration
// settings, however an explicit `Provider` instance may be created and passed during resource
// construction to achieve fine-grained programmatic control over provider settings. See the
// [documentation](https://www.pulumi.com/docs/reference/programming-model/#providers) for more information.
type Provider struct {
	pulumi.ProviderResourceState

	// An basic auth password for kong admin
	KongAdminPassword pulumi.StringPtrOutput `pulumi:"kongAdminPassword"`
	// API key for the kong api (Enterprise Edition)
	KongAdminToken pulumi.StringPtrOutput `pulumi:"kongAdminToken"`
	// The address of the kong admin url e.g. http://localhost:8001
	KongAdminUri pulumi.StringPtrOutput `pulumi:"kongAdminUri"`
	// An basic auth user for kong admin
	KongAdminUsername pulumi.StringPtrOutput `pulumi:"kongAdminUsername"`
	// API key for the kong api (if you have locked it down)
	KongApiKey pulumi.StringPtrOutput `pulumi:"kongApiKey"`
	// Workspace context (Enterprise Edition)
	KongWorkspace pulumi.StringPtrOutput `pulumi:"kongWorkspace"`
}

// NewProvider registers a new resource with the given unique name, arguments, and options.
func NewProvider(ctx *pulumi.Context,
	name string, args *ProviderArgs, opts ...pulumi.ResourceOption) (*Provider, error) {
	if args == nil {
		args = &ProviderArgs{}
	}

	if args.StrictPluginsMatch == nil {
		if d := internal.GetEnvOrDefault(nil, internal.ParseEnvBool, "STRICT_PLUGINS_MATCH"); d != nil {
			args.StrictPluginsMatch = pulumi.BoolPtr(d.(bool))
		}
	}
	if args.TlsSkipVerify == nil {
		if d := internal.GetEnvOrDefault(false, internal.ParseEnvBool, "TLS_SKIP_VERIFY"); d != nil {
			args.TlsSkipVerify = pulumi.BoolPtr(d.(bool))
		}
	}
	opts = internal.PkgResourceDefaultOpts(opts)
	var resource Provider
	err := ctx.RegisterResource("pulumi:providers:kong", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

type providerArgs struct {
	// An basic auth password for kong admin
	KongAdminPassword *string `pulumi:"kongAdminPassword"`
	// API key for the kong api (Enterprise Edition)
	KongAdminToken *string `pulumi:"kongAdminToken"`
	// The address of the kong admin url e.g. http://localhost:8001
	KongAdminUri *string `pulumi:"kongAdminUri"`
	// An basic auth user for kong admin
	KongAdminUsername *string `pulumi:"kongAdminUsername"`
	// API key for the kong api (if you have locked it down)
	KongApiKey *string `pulumi:"kongApiKey"`
	// Workspace context (Enterprise Edition)
	KongWorkspace *string `pulumi:"kongWorkspace"`
	// Should plugins `configJson` field strictly match plugin configuration
	StrictPluginsMatch *bool `pulumi:"strictPluginsMatch"`
	// Whether to skip tls verify for https kong api endpoint using self signed or untrusted certs
	TlsSkipVerify *bool `pulumi:"tlsSkipVerify"`
}

// The set of arguments for constructing a Provider resource.
type ProviderArgs struct {
	// An basic auth password for kong admin
	KongAdminPassword pulumi.StringPtrInput
	// API key for the kong api (Enterprise Edition)
	KongAdminToken pulumi.StringPtrInput
	// The address of the kong admin url e.g. http://localhost:8001
	KongAdminUri pulumi.StringPtrInput
	// An basic auth user for kong admin
	KongAdminUsername pulumi.StringPtrInput
	// API key for the kong api (if you have locked it down)
	KongApiKey pulumi.StringPtrInput
	// Workspace context (Enterprise Edition)
	KongWorkspace pulumi.StringPtrInput
	// Should plugins `configJson` field strictly match plugin configuration
	StrictPluginsMatch pulumi.BoolPtrInput
	// Whether to skip tls verify for https kong api endpoint using self signed or untrusted certs
	TlsSkipVerify pulumi.BoolPtrInput
}

func (ProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*providerArgs)(nil)).Elem()
}

// This function returns a Terraform config object with terraform-namecased keys,to be used with the Terraform Module Provider.
func (r *Provider) TerraformConfig(ctx *pulumi.Context) (ProviderTerraformConfigResultOutput, error) {
	out, err := ctx.Call("pulumi:providers:kong/terraformConfig", nil, ProviderTerraformConfigResultOutput{}, r)
	if err != nil {
		return ProviderTerraformConfigResultOutput{}, err
	}
	return out.(ProviderTerraformConfigResultOutput), nil
}

type ProviderTerraformConfigResult struct {
	Result map[string]interface{} `pulumi:"result"`
}

type ProviderTerraformConfigResultOutput struct{ *pulumi.OutputState }

func (ProviderTerraformConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProviderTerraformConfigResult)(nil)).Elem()
}

func (o ProviderTerraformConfigResultOutput) Result() pulumi.MapOutput {
	return o.ApplyT(func(v ProviderTerraformConfigResult) map[string]interface{} { return v.Result }).(pulumi.MapOutput)
}

type ProviderInput interface {
	pulumi.Input

	ToProviderOutput() ProviderOutput
	ToProviderOutputWithContext(ctx context.Context) ProviderOutput
}

func (*Provider) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (i *Provider) ToProviderOutput() ProviderOutput {
	return i.ToProviderOutputWithContext(context.Background())
}

func (i *Provider) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProviderOutput)
}

type ProviderOutput struct{ *pulumi.OutputState }

func (ProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Provider)(nil)).Elem()
}

func (o ProviderOutput) ToProviderOutput() ProviderOutput {
	return o
}

func (o ProviderOutput) ToProviderOutputWithContext(ctx context.Context) ProviderOutput {
	return o
}

// An basic auth password for kong admin
func (o ProviderOutput) KongAdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.KongAdminPassword }).(pulumi.StringPtrOutput)
}

// API key for the kong api (Enterprise Edition)
func (o ProviderOutput) KongAdminToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.KongAdminToken }).(pulumi.StringPtrOutput)
}

// The address of the kong admin url e.g. http://localhost:8001
func (o ProviderOutput) KongAdminUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.KongAdminUri }).(pulumi.StringPtrOutput)
}

// An basic auth user for kong admin
func (o ProviderOutput) KongAdminUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.KongAdminUsername }).(pulumi.StringPtrOutput)
}

// API key for the kong api (if you have locked it down)
func (o ProviderOutput) KongApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.KongApiKey }).(pulumi.StringPtrOutput)
}

// Workspace context (Enterprise Edition)
func (o ProviderOutput) KongWorkspace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Provider) pulumi.StringPtrOutput { return v.KongWorkspace }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ProviderInput)(nil)).Elem(), &Provider{})
	pulumi.RegisterOutputType(ProviderOutput{})
	pulumi.RegisterOutputType(ProviderTerraformConfigResultOutput{})
}
