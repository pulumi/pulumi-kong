// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kong

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// ## # ConsumerJwtAuth
//
// Consumer jwt auth is a resource that allows you to configure the jwt auth plugin for a consumer.
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
// 		_, err = kong.NewPlugin(ctx, "jwtPlugin", &kong.PluginArgs{
// 			ConfigJson: pulumi.String(fmt.Sprintf("%v%v%v%v", "	{\n", "		\"claims_to_verify\": [\"exp\"]\n", "	}\n", "\n")),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		_, err = kong.NewConsumerJwtAuth(ctx, "consumerJwtConfig", &kong.ConsumerJwtAuthArgs{
// 			Algorithm:    pulumi.String("HS256"),
// 			ConsumerId:   myConsumer.ID(),
// 			Key:          pulumi.String("my_key"),
// 			RsaPublicKey: pulumi.String("foo"),
// 			Secret:       pulumi.String("my_secret"),
// 		})
// 		if err != nil {
// 			return err
// 		}
// 		return nil
// 	})
// }
// ```
type ConsumerJwtAuth struct {
	pulumi.CustomResourceState

	// The algorithm used to verify the token’s signature. Can be HS256, HS384, HS512, RS256, or ES256, Default is `HS256`
	Algorithm pulumi.StringPtrOutput `pulumi:"algorithm"`
	// the id of the consumer to be configured with jwt auth
	ConsumerId pulumi.StringOutput `pulumi:"consumerId"`
	// A unique string identifying the credential. If left out, it will be auto-generated.
	Key pulumi.StringPtrOutput `pulumi:"key"`
	// If algorithm is `RS256` or `ES256`, the public key (in PEM format) to use to verify the token’s signature
	RsaPublicKey pulumi.StringOutput `pulumi:"rsaPublicKey"`
	// If algorithm is `HS256` or `ES256`, the secret used to sign JWTs for this credential. If left out, will be auto-generated
	Secret pulumi.StringPtrOutput `pulumi:"secret"`
	// A list of strings associated with the consumer JWT auth for grouping and filtering
	Tags pulumi.StringArrayOutput `pulumi:"tags"`
}

// NewConsumerJwtAuth registers a new resource with the given unique name, arguments, and options.
func NewConsumerJwtAuth(ctx *pulumi.Context,
	name string, args *ConsumerJwtAuthArgs, opts ...pulumi.ResourceOption) (*ConsumerJwtAuth, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ConsumerId == nil {
		return nil, errors.New("invalid value for required argument 'ConsumerId'")
	}
	if args.RsaPublicKey == nil {
		return nil, errors.New("invalid value for required argument 'RsaPublicKey'")
	}
	var resource ConsumerJwtAuth
	err := ctx.RegisterResource("kong:index/consumerJwtAuth:ConsumerJwtAuth", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetConsumerJwtAuth gets an existing ConsumerJwtAuth resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetConsumerJwtAuth(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ConsumerJwtAuthState, opts ...pulumi.ResourceOption) (*ConsumerJwtAuth, error) {
	var resource ConsumerJwtAuth
	err := ctx.ReadResource("kong:index/consumerJwtAuth:ConsumerJwtAuth", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering ConsumerJwtAuth resources.
type consumerJwtAuthState struct {
	// The algorithm used to verify the token’s signature. Can be HS256, HS384, HS512, RS256, or ES256, Default is `HS256`
	Algorithm *string `pulumi:"algorithm"`
	// the id of the consumer to be configured with jwt auth
	ConsumerId *string `pulumi:"consumerId"`
	// A unique string identifying the credential. If left out, it will be auto-generated.
	Key *string `pulumi:"key"`
	// If algorithm is `RS256` or `ES256`, the public key (in PEM format) to use to verify the token’s signature
	RsaPublicKey *string `pulumi:"rsaPublicKey"`
	// If algorithm is `HS256` or `ES256`, the secret used to sign JWTs for this credential. If left out, will be auto-generated
	Secret *string `pulumi:"secret"`
	// A list of strings associated with the consumer JWT auth for grouping and filtering
	Tags []string `pulumi:"tags"`
}

type ConsumerJwtAuthState struct {
	// The algorithm used to verify the token’s signature. Can be HS256, HS384, HS512, RS256, or ES256, Default is `HS256`
	Algorithm pulumi.StringPtrInput
	// the id of the consumer to be configured with jwt auth
	ConsumerId pulumi.StringPtrInput
	// A unique string identifying the credential. If left out, it will be auto-generated.
	Key pulumi.StringPtrInput
	// If algorithm is `RS256` or `ES256`, the public key (in PEM format) to use to verify the token’s signature
	RsaPublicKey pulumi.StringPtrInput
	// If algorithm is `HS256` or `ES256`, the secret used to sign JWTs for this credential. If left out, will be auto-generated
	Secret pulumi.StringPtrInput
	// A list of strings associated with the consumer JWT auth for grouping and filtering
	Tags pulumi.StringArrayInput
}

func (ConsumerJwtAuthState) ElementType() reflect.Type {
	return reflect.TypeOf((*consumerJwtAuthState)(nil)).Elem()
}

type consumerJwtAuthArgs struct {
	// The algorithm used to verify the token’s signature. Can be HS256, HS384, HS512, RS256, or ES256, Default is `HS256`
	Algorithm *string `pulumi:"algorithm"`
	// the id of the consumer to be configured with jwt auth
	ConsumerId string `pulumi:"consumerId"`
	// A unique string identifying the credential. If left out, it will be auto-generated.
	Key *string `pulumi:"key"`
	// If algorithm is `RS256` or `ES256`, the public key (in PEM format) to use to verify the token’s signature
	RsaPublicKey string `pulumi:"rsaPublicKey"`
	// If algorithm is `HS256` or `ES256`, the secret used to sign JWTs for this credential. If left out, will be auto-generated
	Secret *string `pulumi:"secret"`
	// A list of strings associated with the consumer JWT auth for grouping and filtering
	Tags []string `pulumi:"tags"`
}

// The set of arguments for constructing a ConsumerJwtAuth resource.
type ConsumerJwtAuthArgs struct {
	// The algorithm used to verify the token’s signature. Can be HS256, HS384, HS512, RS256, or ES256, Default is `HS256`
	Algorithm pulumi.StringPtrInput
	// the id of the consumer to be configured with jwt auth
	ConsumerId pulumi.StringInput
	// A unique string identifying the credential. If left out, it will be auto-generated.
	Key pulumi.StringPtrInput
	// If algorithm is `RS256` or `ES256`, the public key (in PEM format) to use to verify the token’s signature
	RsaPublicKey pulumi.StringInput
	// If algorithm is `HS256` or `ES256`, the secret used to sign JWTs for this credential. If left out, will be auto-generated
	Secret pulumi.StringPtrInput
	// A list of strings associated with the consumer JWT auth for grouping and filtering
	Tags pulumi.StringArrayInput
}

func (ConsumerJwtAuthArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*consumerJwtAuthArgs)(nil)).Elem()
}

type ConsumerJwtAuthInput interface {
	pulumi.Input

	ToConsumerJwtAuthOutput() ConsumerJwtAuthOutput
	ToConsumerJwtAuthOutputWithContext(ctx context.Context) ConsumerJwtAuthOutput
}

func (*ConsumerJwtAuth) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsumerJwtAuth)(nil))
}

func (i *ConsumerJwtAuth) ToConsumerJwtAuthOutput() ConsumerJwtAuthOutput {
	return i.ToConsumerJwtAuthOutputWithContext(context.Background())
}

func (i *ConsumerJwtAuth) ToConsumerJwtAuthOutputWithContext(ctx context.Context) ConsumerJwtAuthOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsumerJwtAuthOutput)
}

func (i *ConsumerJwtAuth) ToConsumerJwtAuthPtrOutput() ConsumerJwtAuthPtrOutput {
	return i.ToConsumerJwtAuthPtrOutputWithContext(context.Background())
}

func (i *ConsumerJwtAuth) ToConsumerJwtAuthPtrOutputWithContext(ctx context.Context) ConsumerJwtAuthPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsumerJwtAuthPtrOutput)
}

type ConsumerJwtAuthPtrInput interface {
	pulumi.Input

	ToConsumerJwtAuthPtrOutput() ConsumerJwtAuthPtrOutput
	ToConsumerJwtAuthPtrOutputWithContext(ctx context.Context) ConsumerJwtAuthPtrOutput
}

type consumerJwtAuthPtrType ConsumerJwtAuthArgs

func (*consumerJwtAuthPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsumerJwtAuth)(nil))
}

func (i *consumerJwtAuthPtrType) ToConsumerJwtAuthPtrOutput() ConsumerJwtAuthPtrOutput {
	return i.ToConsumerJwtAuthPtrOutputWithContext(context.Background())
}

func (i *consumerJwtAuthPtrType) ToConsumerJwtAuthPtrOutputWithContext(ctx context.Context) ConsumerJwtAuthPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsumerJwtAuthPtrOutput)
}

// ConsumerJwtAuthArrayInput is an input type that accepts ConsumerJwtAuthArray and ConsumerJwtAuthArrayOutput values.
// You can construct a concrete instance of `ConsumerJwtAuthArrayInput` via:
//
//          ConsumerJwtAuthArray{ ConsumerJwtAuthArgs{...} }
type ConsumerJwtAuthArrayInput interface {
	pulumi.Input

	ToConsumerJwtAuthArrayOutput() ConsumerJwtAuthArrayOutput
	ToConsumerJwtAuthArrayOutputWithContext(context.Context) ConsumerJwtAuthArrayOutput
}

type ConsumerJwtAuthArray []ConsumerJwtAuthInput

func (ConsumerJwtAuthArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]*ConsumerJwtAuth)(nil)).Elem()
}

func (i ConsumerJwtAuthArray) ToConsumerJwtAuthArrayOutput() ConsumerJwtAuthArrayOutput {
	return i.ToConsumerJwtAuthArrayOutputWithContext(context.Background())
}

func (i ConsumerJwtAuthArray) ToConsumerJwtAuthArrayOutputWithContext(ctx context.Context) ConsumerJwtAuthArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsumerJwtAuthArrayOutput)
}

// ConsumerJwtAuthMapInput is an input type that accepts ConsumerJwtAuthMap and ConsumerJwtAuthMapOutput values.
// You can construct a concrete instance of `ConsumerJwtAuthMapInput` via:
//
//          ConsumerJwtAuthMap{ "key": ConsumerJwtAuthArgs{...} }
type ConsumerJwtAuthMapInput interface {
	pulumi.Input

	ToConsumerJwtAuthMapOutput() ConsumerJwtAuthMapOutput
	ToConsumerJwtAuthMapOutputWithContext(context.Context) ConsumerJwtAuthMapOutput
}

type ConsumerJwtAuthMap map[string]ConsumerJwtAuthInput

func (ConsumerJwtAuthMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]*ConsumerJwtAuth)(nil)).Elem()
}

func (i ConsumerJwtAuthMap) ToConsumerJwtAuthMapOutput() ConsumerJwtAuthMapOutput {
	return i.ToConsumerJwtAuthMapOutputWithContext(context.Background())
}

func (i ConsumerJwtAuthMap) ToConsumerJwtAuthMapOutputWithContext(ctx context.Context) ConsumerJwtAuthMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsumerJwtAuthMapOutput)
}

type ConsumerJwtAuthOutput struct{ *pulumi.OutputState }

func (ConsumerJwtAuthOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsumerJwtAuth)(nil))
}

func (o ConsumerJwtAuthOutput) ToConsumerJwtAuthOutput() ConsumerJwtAuthOutput {
	return o
}

func (o ConsumerJwtAuthOutput) ToConsumerJwtAuthOutputWithContext(ctx context.Context) ConsumerJwtAuthOutput {
	return o
}

func (o ConsumerJwtAuthOutput) ToConsumerJwtAuthPtrOutput() ConsumerJwtAuthPtrOutput {
	return o.ToConsumerJwtAuthPtrOutputWithContext(context.Background())
}

func (o ConsumerJwtAuthOutput) ToConsumerJwtAuthPtrOutputWithContext(ctx context.Context) ConsumerJwtAuthPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConsumerJwtAuth) *ConsumerJwtAuth {
		return &v
	}).(ConsumerJwtAuthPtrOutput)
}

type ConsumerJwtAuthPtrOutput struct{ *pulumi.OutputState }

func (ConsumerJwtAuthPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsumerJwtAuth)(nil))
}

func (o ConsumerJwtAuthPtrOutput) ToConsumerJwtAuthPtrOutput() ConsumerJwtAuthPtrOutput {
	return o
}

func (o ConsumerJwtAuthPtrOutput) ToConsumerJwtAuthPtrOutputWithContext(ctx context.Context) ConsumerJwtAuthPtrOutput {
	return o
}

func (o ConsumerJwtAuthPtrOutput) Elem() ConsumerJwtAuthOutput {
	return o.ApplyT(func(v *ConsumerJwtAuth) ConsumerJwtAuth {
		if v != nil {
			return *v
		}
		var ret ConsumerJwtAuth
		return ret
	}).(ConsumerJwtAuthOutput)
}

type ConsumerJwtAuthArrayOutput struct{ *pulumi.OutputState }

func (ConsumerJwtAuthArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConsumerJwtAuth)(nil))
}

func (o ConsumerJwtAuthArrayOutput) ToConsumerJwtAuthArrayOutput() ConsumerJwtAuthArrayOutput {
	return o
}

func (o ConsumerJwtAuthArrayOutput) ToConsumerJwtAuthArrayOutputWithContext(ctx context.Context) ConsumerJwtAuthArrayOutput {
	return o
}

func (o ConsumerJwtAuthArrayOutput) Index(i pulumi.IntInput) ConsumerJwtAuthOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConsumerJwtAuth {
		return vs[0].([]ConsumerJwtAuth)[vs[1].(int)]
	}).(ConsumerJwtAuthOutput)
}

type ConsumerJwtAuthMapOutput struct{ *pulumi.OutputState }

func (ConsumerJwtAuthMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ConsumerJwtAuth)(nil))
}

func (o ConsumerJwtAuthMapOutput) ToConsumerJwtAuthMapOutput() ConsumerJwtAuthMapOutput {
	return o
}

func (o ConsumerJwtAuthMapOutput) ToConsumerJwtAuthMapOutputWithContext(ctx context.Context) ConsumerJwtAuthMapOutput {
	return o
}

func (o ConsumerJwtAuthMapOutput) MapIndex(k pulumi.StringInput) ConsumerJwtAuthOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ConsumerJwtAuth {
		return vs[0].(map[string]ConsumerJwtAuth)[vs[1].(string)]
	}).(ConsumerJwtAuthOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ConsumerJwtAuthInput)(nil)).Elem(), &ConsumerJwtAuth{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConsumerJwtAuthPtrInput)(nil)).Elem(), &ConsumerJwtAuth{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConsumerJwtAuthArrayInput)(nil)).Elem(), ConsumerJwtAuthArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConsumerJwtAuthMapInput)(nil)).Elem(), ConsumerJwtAuthMap{})
	pulumi.RegisterOutputType(ConsumerJwtAuthOutput{})
	pulumi.RegisterOutputType(ConsumerJwtAuthPtrOutput{})
	pulumi.RegisterOutputType(ConsumerJwtAuthArrayOutput{})
	pulumi.RegisterOutputType(ConsumerJwtAuthMapOutput{})
}