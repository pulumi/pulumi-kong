// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kong

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # ConsumerOauth2
//
// Resource that allows you to configure the OAuth2 plugin credentials for a consumer.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
// 	"fmt"
//
// 	"github.com/pulumi/pulumi-kong/sdk/v4/go/kong"
// 	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
// )
//
// func main() {
// 	pulumi.Run(func(ctx *pulumi.Context) error {
// 		myConsumer, err := kong.NewConsumer(ctx, "myConsumer", &kong.ConsumerArgs{
// 			CustomId: pulumi.String("123"),
// 			Username: pulumi.String("User1"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = kong.NewPlugin(ctx, "oauth2Plugin", &kong.PluginArgs{
// 			ConfigJson: pulumi.String(fmt.Sprintf("%v%v%v%v%v%v%v%v", "	{\n", "		\"global_credentials\": true,\n", "		\"enable_password_grant\": true,\n", "		\"token_expiration\": 180,\n", "		\"refresh_token_ttl\": 180,\n", "		\"provision_key\": \"testprovisionkey\"\n", "	}\n", "\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = kong.NewConsumerOauth2(ctx, "consumerOauth2", &kong.ConsumerOauth2Args{
// 			ClientId:     pulumi.String("client_id"),
// 			ClientSecret: pulumi.String("client_secret"),
// 			ConsumerId:   myConsumer.ID(),
// 			RedirectUris: pulumi.StringArray{
// 				pulumi.String("https://asdf.com/callback"),
// 				pulumi.String("https://test.cl/callback"),
// 			},
// 			Tags: pulumi.StringArray{
// 				pulumi.String("myTag"),
// 			},
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ConsumerOauth2 struct {
	pulumi.CustomResourceState

	// Unique oauth2 client id. If not set, the oauth2 plugin will generate one
	ClientId pulumi.StringPtrOutput `pulumi:"clientId"`
	// Unique oauth2 client secret. If not set, the oauth2 plugin will generate one
	ClientSecret pulumi.StringPtrOutput `pulumi:"clientSecret"`
	// The id of the consumer to be configured with oauth2.
	ConsumerId pulumi.StringOutput `pulumi:"consumerId"`
	// A boolean flag that indicates whether the clientSecret field will be stored in hashed form. If enabled on existing plugin instances, client secrets are hashed on the fly upon first usage. Default: `false`.
	HashSecret pulumi.BoolPtrOutput `pulumi:"hashSecret"`
	// The name associated with the credential.
	Name pulumi.StringOutput `pulumi:"name"`
	// An array with one or more URLs in your app where users will be sent after authorization ([RFC 6742 Section 3.1.2](https://tools.ietf.org/html/rfc6749#section-3.1.2)).
	RedirectUris pulumi.StringArrayOutput `pulumi:"redirectUris"`
	// A list of strings associated with the consumer for grouping and filtering.
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
}

// NewConsumerOauth2 registers a new resource with the given unique name, arguments, and options.
func NewConsumerOauth2(ctx *pulumi.Context,
	name string, args *ConsumerOauth2Args, opts ...pulumi.ResourceOption) (*ConsumerOauth2, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConsumerId == nil {
		return nil, errors.New("invalid value for required argument 'ConsumerId'")
	}
	if args.RedirectUris == nil {
		return nil, errors.New("invalid value for required argument 'RedirectUris'")
	}
	var resource ConsumerOauth2
	err := ctx.RegisterResource("kong:index/consumerOauth2:ConsumerOauth2", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConsumerOauth2 gets an existing ConsumerOauth2 resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConsumerOauth2(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConsumerOauth2State, opts ...pulumi.ResourceOption) (*ConsumerOauth2, error) {
	var resource ConsumerOauth2
	err := ctx.ReadResource("kong:index/consumerOauth2:ConsumerOauth2", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConsumerOauth2 resources.
type consumerOauth2State struct {
	// Unique oauth2 client id. If not set, the oauth2 plugin will generate one
	ClientId *string `pulumi:"clientId"`
	// Unique oauth2 client secret. If not set, the oauth2 plugin will generate one
	ClientSecret *string `pulumi:"clientSecret"`
	// The id of the consumer to be configured with oauth2.
	ConsumerId *string `pulumi:"consumerId"`
	// A boolean flag that indicates whether the clientSecret field will be stored in hashed form. If enabled on existing plugin instances, client secrets are hashed on the fly upon first usage. Default: `false`.
	HashSecret *bool `pulumi:"hashSecret"`
	// The name associated with the credential.
	Name *string `pulumi:"name"`
	// An array with one or more URLs in your app where users will be sent after authorization ([RFC 6742 Section 3.1.2](https://tools.ietf.org/html/rfc6749#section-3.1.2)).
	RedirectUris []string `pulumi:"redirectUris"`
	// A list of strings associated with the consumer for grouping and filtering.
	Tags []string `pulumi:"tags"`
}

type ConsumerOauth2State struct {
	// Unique oauth2 client id. If not set, the oauth2 plugin will generate one
	ClientId pulumi.StringPtrInput
	// Unique oauth2 client secret. If not set, the oauth2 plugin will generate one
	ClientSecret pulumi.StringPtrInput
	// The id of the consumer to be configured with oauth2.
	ConsumerId pulumi.StringPtrInput
	// A boolean flag that indicates whether the clientSecret field will be stored in hashed form. If enabled on existing plugin instances, client secrets are hashed on the fly upon first usage. Default: `false`.
	HashSecret pulumi.BoolPtrInput
	// The name associated with the credential.
	Name pulumi.StringPtrInput
	// An array with one or more URLs in your app where users will be sent after authorization ([RFC 6742 Section 3.1.2](https://tools.ietf.org/html/rfc6749#section-3.1.2)).
	RedirectUris pulumi.StringArrayInput
	// A list of strings associated with the consumer for grouping and filtering.
	Tags pulumi.StringArrayInput
}

func (ConsumerOauth2State) ElementType() reflect.Type {
	return reflect.TypeOf((*consumerOauth2State)(nil)).Elem()
}

type consumerOauth2Args struct {
	// Unique oauth2 client id. If not set, the oauth2 plugin will generate one
	ClientId *string `pulumi:"clientId"`
	// Unique oauth2 client secret. If not set, the oauth2 plugin will generate one
	ClientSecret *string `pulumi:"clientSecret"`
	// The id of the consumer to be configured with oauth2.
	ConsumerId string `pulumi:"consumerId"`
	// A boolean flag that indicates whether the clientSecret field will be stored in hashed form. If enabled on existing plugin instances, client secrets are hashed on the fly upon first usage. Default: `false`.
	HashSecret *bool `pulumi:"hashSecret"`
	// The name associated with the credential.
	Name *string `pulumi:"name"`
	// An array with one or more URLs in your app where users will be sent after authorization ([RFC 6742 Section 3.1.2](https://tools.ietf.org/html/rfc6749#section-3.1.2)).
	RedirectUris []string `pulumi:"redirectUris"`
	// A list of strings associated with the consumer for grouping and filtering.
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a ConsumerOauth2 resource.
type ConsumerOauth2Args struct {
	// Unique oauth2 client id. If not set, the oauth2 plugin will generate one
	ClientId pulumi.StringPtrInput
	// Unique oauth2 client secret. If not set, the oauth2 plugin will generate one
	ClientSecret pulumi.StringPtrInput
	// The id of the consumer to be configured with oauth2.
	ConsumerId pulumi.StringInput
	// A boolean flag that indicates whether the clientSecret field will be stored in hashed form. If enabled on existing plugin instances, client secrets are hashed on the fly upon first usage. Default: `false`.
	HashSecret pulumi.BoolPtrInput
	// The name associated with the credential.
	Name pulumi.StringPtrInput
	// An array with one or more URLs in your app where users will be sent after authorization ([RFC 6742 Section 3.1.2](https://tools.ietf.org/html/rfc6749#section-3.1.2)).
	RedirectUris pulumi.StringArrayInput
	// A list of strings associated with the consumer for grouping and filtering.
	Tags pulumi.StringArrayInput
}

func (ConsumerOauth2Args) ElementType() reflect.Type {
	return reflect.TypeOf((*consumerOauth2Args)(nil)).Elem()
}

type ConsumerOauth2Input interface {
	pulumi.Input

	ToConsumerOauth2Output() ConsumerOauth2Output
	ToConsumerOauth2OutputWithContext(ctx context.Context) ConsumerOauth2Output
}

func (*ConsumerOauth2) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsumerOauth2)(nil)).Elem()
}

func (i *ConsumerOauth2) ToConsumerOauth2Output() ConsumerOauth2Output {
	return i.ToConsumerOauth2OutputWithContext(context.Background())
}

func (i *ConsumerOauth2) ToConsumerOauth2OutputWithContext(ctx context.Context) ConsumerOauth2Output {
	return pulumi.ToOutputWithContext(ctx, i).(ConsumerOauth2Output)
}

// ConsumerOauth2ArrayInput is an input type that accepts ConsumerOauth2Array and ConsumerOauth2ArrayOutput values.
// You can construct a concrete instance of `ConsumerOauth2ArrayInput` via:
//
//          ConsumerOauth2Array{ ConsumerOauth2Args{...} }
type ConsumerOauth2ArrayInput interface {
	pulumi.Input

	ToConsumerOauth2ArrayOutput() ConsumerOauth2ArrayOutput
	ToConsumerOauth2ArrayOutputWithContext(context.Context) ConsumerOauth2ArrayOutput
}

type ConsumerOauth2Array []ConsumerOauth2Input

func (ConsumerOauth2Array) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConsumerOauth2)(nil)).Elem()
}

func (i ConsumerOauth2Array) ToConsumerOauth2ArrayOutput() ConsumerOauth2ArrayOutput {
	return i.ToConsumerOauth2ArrayOutputWithContext(context.Background())
}

func (i ConsumerOauth2Array) ToConsumerOauth2ArrayOutputWithContext(ctx context.Context) ConsumerOauth2ArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsumerOauth2ArrayOutput)
}

// ConsumerOauth2MapInput is an input type that accepts ConsumerOauth2Map and ConsumerOauth2MapOutput values.
// You can construct a concrete instance of `ConsumerOauth2MapInput` via:
//
//          ConsumerOauth2Map{ "key": ConsumerOauth2Args{...} }
type ConsumerOauth2MapInput interface {
	pulumi.Input

	ToConsumerOauth2MapOutput() ConsumerOauth2MapOutput
	ToConsumerOauth2MapOutputWithContext(context.Context) ConsumerOauth2MapOutput
}

type ConsumerOauth2Map map[string]ConsumerOauth2Input

func (ConsumerOauth2Map) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConsumerOauth2)(nil)).Elem()
}

func (i ConsumerOauth2Map) ToConsumerOauth2MapOutput() ConsumerOauth2MapOutput {
	return i.ToConsumerOauth2MapOutputWithContext(context.Background())
}

func (i ConsumerOauth2Map) ToConsumerOauth2MapOutputWithContext(ctx context.Context) ConsumerOauth2MapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsumerOauth2MapOutput)
}

type ConsumerOauth2Output struct{ *pulumi.OutputState }

func (ConsumerOauth2Output) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsumerOauth2)(nil)).Elem()
}

func (o ConsumerOauth2Output) ToConsumerOauth2Output() ConsumerOauth2Output {
	return o
}

func (o ConsumerOauth2Output) ToConsumerOauth2OutputWithContext(ctx context.Context) ConsumerOauth2Output {
	return o
}

type ConsumerOauth2ArrayOutput struct{ *pulumi.OutputState }

func (ConsumerOauth2ArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConsumerOauth2)(nil)).Elem()
}

func (o ConsumerOauth2ArrayOutput) ToConsumerOauth2ArrayOutput() ConsumerOauth2ArrayOutput {
	return o
}

func (o ConsumerOauth2ArrayOutput) ToConsumerOauth2ArrayOutputWithContext(ctx context.Context) ConsumerOauth2ArrayOutput {
	return o
}

func (o ConsumerOauth2ArrayOutput) Index(i pulumi.IntInput) ConsumerOauth2Output {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) *ConsumerOauth2 {
		return vs[0].([]*ConsumerOauth2)[vs[1].(int)]
	}).(ConsumerOauth2Output)
}

type ConsumerOauth2MapOutput struct{ *pulumi.OutputState }

func (ConsumerOauth2MapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConsumerOauth2)(nil)).Elem()
}

func (o ConsumerOauth2MapOutput) ToConsumerOauth2MapOutput() ConsumerOauth2MapOutput {
	return o
}

func (o ConsumerOauth2MapOutput) ToConsumerOauth2MapOutputWithContext(ctx context.Context) ConsumerOauth2MapOutput {
	return o
}

func (o ConsumerOauth2MapOutput) MapIndex(k pulumi.StringInput) ConsumerOauth2Output {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) *ConsumerOauth2 {
		return vs[0].(map[string]*ConsumerOauth2)[vs[1].(string)]
	}).(ConsumerOauth2Output)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConsumerOauth2Input)(nil)).Elem(), &ConsumerOauth2{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConsumerOauth2ArrayInput)(nil)).Elem(), ConsumerOauth2Array{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConsumerOauth2MapInput)(nil)).Elem(), ConsumerOauth2Map{})
	pulumi.RegisterOutputType(ConsumerOauth2Output{})
	pulumi.RegisterOutputType(ConsumerOauth2ArrayOutput{})
	pulumi.RegisterOutputType(ConsumerOauth2MapOutput{})
}