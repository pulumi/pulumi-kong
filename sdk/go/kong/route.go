// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kong

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type Route struct {
	pulumi.CustomResourceState

	Destinations  RouteDestinationArrayOutput `pulumi:"destinations"`
	Hosts         pulumi.StringArrayOutput    `pulumi:"hosts"`
	Methods       pulumi.StringArrayOutput    `pulumi:"methods"`
	Name          pulumi.StringOutput         `pulumi:"name"`
	Paths         pulumi.StringArrayOutput    `pulumi:"paths"`
	PreserveHost  pulumi.BoolPtrOutput        `pulumi:"preserveHost"`
	Protocols     pulumi.StringArrayOutput    `pulumi:"protocols"`
	RegexPriority pulumi.IntPtrOutput         `pulumi:"regexPriority"`
	ServiceId     pulumi.StringOutput         `pulumi:"serviceId"`
	Snis          pulumi.StringArrayOutput    `pulumi:"snis"`
	Sources       RouteSourceArrayOutput      `pulumi:"sources"`
	StripPath     pulumi.BoolPtrOutput        `pulumi:"stripPath"`
}

// NewRoute registers a new resource with the given unique name, arguments, and options.
func NewRoute(ctx *pulumi.Context,
	name string, args *RouteArgs, opts ...pulumi.ResourceOption) (*Route, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Protocols == nil {
		return nil, errors.New("invalid value for required argument 'Protocols'")
	}
	if args.ServiceId == nil {
		return nil, errors.New("invalid value for required argument 'ServiceId'")
	}
	var resource Route
	err := ctx.RegisterResource("kong:index/route:Route", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// GetRoute gets an existing Route resource's state with the given name, ID, and optional
// state properties that are used to uniquely qualify the lookup (nil if not required).
func GetRoute(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RouteState, opts ...pulumi.ResourceOption) (*Route, error) {
	var resource Route
	err := ctx.ReadResource("kong:index/route:Route", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}

// Input properties used for looking up and filtering Route resources.
type routeState struct {
	Destinations  []RouteDestination `pulumi:"destinations"`
	Hosts         []string           `pulumi:"hosts"`
	Methods       []string           `pulumi:"methods"`
	Name          *string            `pulumi:"name"`
	Paths         []string           `pulumi:"paths"`
	PreserveHost  *bool              `pulumi:"preserveHost"`
	Protocols     []string           `pulumi:"protocols"`
	RegexPriority *int               `pulumi:"regexPriority"`
	ServiceId     *string            `pulumi:"serviceId"`
	Snis          []string           `pulumi:"snis"`
	Sources       []RouteSource      `pulumi:"sources"`
	StripPath     *bool              `pulumi:"stripPath"`
}

type RouteState struct {
	Destinations  RouteDestinationArrayInput
	Hosts         pulumi.StringArrayInput
	Methods       pulumi.StringArrayInput
	Name          pulumi.StringPtrInput
	Paths         pulumi.StringArrayInput
	PreserveHost  pulumi.BoolPtrInput
	Protocols     pulumi.StringArrayInput
	RegexPriority pulumi.IntPtrInput
	ServiceId     pulumi.StringPtrInput
	Snis          pulumi.StringArrayInput
	Sources       RouteSourceArrayInput
	StripPath     pulumi.BoolPtrInput
}

func (RouteState) ElementType() reflect.Type {
	return reflect.TypeOf((*routeState)(nil)).Elem()
}

type routeArgs struct {
	Destinations  []RouteDestination `pulumi:"destinations"`
	Hosts         []string           `pulumi:"hosts"`
	Methods       []string           `pulumi:"methods"`
	Name          *string            `pulumi:"name"`
	Paths         []string           `pulumi:"paths"`
	PreserveHost  *bool              `pulumi:"preserveHost"`
	Protocols     []string           `pulumi:"protocols"`
	RegexPriority *int               `pulumi:"regexPriority"`
	ServiceId     string             `pulumi:"serviceId"`
	Snis          []string           `pulumi:"snis"`
	Sources       []RouteSource      `pulumi:"sources"`
	StripPath     *bool              `pulumi:"stripPath"`
}

// The set of arguments for constructing a Route resource.
type RouteArgs struct {
	Destinations  RouteDestinationArrayInput
	Hosts         pulumi.StringArrayInput
	Methods       pulumi.StringArrayInput
	Name          pulumi.StringPtrInput
	Paths         pulumi.StringArrayInput
	PreserveHost  pulumi.BoolPtrInput
	Protocols     pulumi.StringArrayInput
	RegexPriority pulumi.IntPtrInput
	ServiceId     pulumi.StringInput
	Snis          pulumi.StringArrayInput
	Sources       RouteSourceArrayInput
	StripPath     pulumi.BoolPtrInput
}

func (RouteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*routeArgs)(nil)).Elem()
}

type RouteInput interface {
	pulumi.Input

	ToRouteOutput() RouteOutput
	ToRouteOutputWithContext(ctx context.Context) RouteOutput
}

func (*Route) ElementType() reflect.Type {
	return reflect.TypeOf((*Route)(nil))
}

func (i *Route) ToRouteOutput() RouteOutput {
	return i.ToRouteOutputWithContext(context.Background())
}

func (i *Route) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteOutput)
}

func (i *Route) ToRoutePtrOutput() RoutePtrOutput {
	return i.ToRoutePtrOutputWithContext(context.Background())
}

func (i *Route) ToRoutePtrOutputWithContext(ctx context.Context) RoutePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutePtrOutput)
}

type RoutePtrInput interface {
	pulumi.Input

	ToRoutePtrOutput() RoutePtrOutput
	ToRoutePtrOutputWithContext(ctx context.Context) RoutePtrOutput
}

type routePtrType RouteArgs

func (*routePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Route)(nil))
}

func (i *routePtrType) ToRoutePtrOutput() RoutePtrOutput {
	return i.ToRoutePtrOutputWithContext(context.Background())
}

func (i *routePtrType) ToRoutePtrOutputWithContext(ctx context.Context) RoutePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RoutePtrOutput)
}

// RouteArrayInput is an input type that accepts RouteArray and RouteArrayOutput values.
// You can construct a concrete instance of `RouteArrayInput` via:
//
//          RouteArray{ RouteArgs{...} }
type RouteArrayInput interface {
	pulumi.Input

	ToRouteArrayOutput() RouteArrayOutput
	ToRouteArrayOutputWithContext(context.Context) RouteArrayOutput
}

type RouteArray []RouteInput

func (RouteArray) ElementType() reflect.Type {
	return reflect.TypeOf(([]*Route)(nil))
}

func (i RouteArray) ToRouteArrayOutput() RouteArrayOutput {
	return i.ToRouteArrayOutputWithContext(context.Background())
}

func (i RouteArray) ToRouteArrayOutputWithContext(ctx context.Context) RouteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteArrayOutput)
}

// RouteMapInput is an input type that accepts RouteMap and RouteMapOutput values.
// You can construct a concrete instance of `RouteMapInput` via:
//
//          RouteMap{ "key": RouteArgs{...} }
type RouteMapInput interface {
	pulumi.Input

	ToRouteMapOutput() RouteMapOutput
	ToRouteMapOutputWithContext(context.Context) RouteMapOutput
}

type RouteMap map[string]RouteInput

func (RouteMap) ElementType() reflect.Type {
	return reflect.TypeOf((map[string]*Route)(nil))
}

func (i RouteMap) ToRouteMapOutput() RouteMapOutput {
	return i.ToRouteMapOutputWithContext(context.Background())
}

func (i RouteMap) ToRouteMapOutputWithContext(ctx context.Context) RouteMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RouteMapOutput)
}

type RouteOutput struct {
	*pulumi.OutputState
}

func (RouteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Route)(nil))
}

func (o RouteOutput) ToRouteOutput() RouteOutput {
	return o
}

func (o RouteOutput) ToRouteOutputWithContext(ctx context.Context) RouteOutput {
	return o
}

func (o RouteOutput) ToRoutePtrOutput() RoutePtrOutput {
	return o.ToRoutePtrOutputWithContext(context.Background())
}

func (o RouteOutput) ToRoutePtrOutputWithContext(ctx context.Context) RoutePtrOutput {
	return o.ApplyT(func(v Route) *Route {
		return &v
	}).(RoutePtrOutput)
}

type RoutePtrOutput struct {
	*pulumi.OutputState
}

func (RoutePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Route)(nil))
}

func (o RoutePtrOutput) ToRoutePtrOutput() RoutePtrOutput {
	return o
}

func (o RoutePtrOutput) ToRoutePtrOutputWithContext(ctx context.Context) RoutePtrOutput {
	return o
}

type RouteArrayOutput struct{ *pulumi.OutputState }

func (RouteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Route)(nil))
}

func (o RouteArrayOutput) ToRouteArrayOutput() RouteArrayOutput {
	return o
}

func (o RouteArrayOutput) ToRouteArrayOutputWithContext(ctx context.Context) RouteArrayOutput {
	return o
}

func (o RouteArrayOutput) Index(i pulumi.IntInput) RouteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Route {
		return vs[0].([]Route)[vs[1].(int)]
	}).(RouteOutput)
}

type RouteMapOutput struct{ *pulumi.OutputState }

func (RouteMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]Route)(nil))
}

func (o RouteMapOutput) ToRouteMapOutput() RouteMapOutput {
	return o
}

func (o RouteMapOutput) ToRouteMapOutputWithContext(ctx context.Context) RouteMapOutput {
	return o
}

func (o RouteMapOutput) MapIndex(k pulumi.StringInput) RouteOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) Route {
		return vs[0].(map[string]Route)[vs[1].(string)]
	}).(RouteOutput)
}

func init() {
	pulumi.RegisterOutputType(RouteOutput{})
	pulumi.RegisterOutputType(RoutePtrOutput{})
	pulumi.RegisterOutputType(RouteArrayOutput{})
	pulumi.RegisterOutputType(RouteMapOutput{})
}
