// *** WARNING: this file was generated by the Pulumi Terraform Bridge (tfgen) Tool. ***
// *** Do not edit by hand unless you're certain you know what you are doing! ***

package kong

import (
	"fmt"

	"github.com/blang/semver"
	"github.com/pulumi/pulumi/sdk/v2/go/pulumi"
)

type module struct {
	version semver.Version
}

func (m *module) Version() semver.Version {
	return m.version
}

func (m *module) Construct(ctx *pulumi.Context, name, typ, urn string) (r pulumi.Resource, err error) {
	switch typ {
	case "kong:index/certificate:Certificate":
		r, err = NewCertificate(ctx, name, nil, pulumi.URN_(urn))
	case "kong:index/consumer:Consumer":
		r, err = NewConsumer(ctx, name, nil, pulumi.URN_(urn))
	case "kong:index/consumerPluginConfig:ConsumerPluginConfig":
		r, err = NewConsumerPluginConfig(ctx, name, nil, pulumi.URN_(urn))
	case "kong:index/plugin:Plugin":
		r, err = NewPlugin(ctx, name, nil, pulumi.URN_(urn))
	case "kong:index/route:Route":
		r, err = NewRoute(ctx, name, nil, pulumi.URN_(urn))
	case "kong:index/service:Service":
		r, err = NewService(ctx, name, nil, pulumi.URN_(urn))
	case "kong:index/sni:Sni":
		r, err = NewSni(ctx, name, nil, pulumi.URN_(urn))
	case "kong:index/target:Target":
		r, err = NewTarget(ctx, name, nil, pulumi.URN_(urn))
	case "kong:index/upstream:Upstream":
		r, err = NewUpstream(ctx, name, nil, pulumi.URN_(urn))
	default:
		return nil, fmt.Errorf("unknown resource type: %s", typ)
	}

	return
}

type pkg struct {
	version semver.Version
}

func (p *pkg) Version() semver.Version {
	return p.version
}

func (p *pkg) ConstructProvider(ctx *pulumi.Context, name, typ, urn string) (pulumi.ProviderResource, error) {
	if typ != "pulumi:providers:kong" {
		return nil, fmt.Errorf("unknown provider type: %s", typ)
	}

	return NewProvider(ctx, name, nil, pulumi.URN_(urn))
}

func init() {
	version, err := PkgVersion()
	if err != nil {
		fmt.Println("failed to determine package version. defaulting to v1: %v", err)
	}
	pulumi.RegisterResourceModule(
		"kong",
		"index/certificate",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"kong",
		"index/consumer",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"kong",
		"index/consumerPluginConfig",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"kong",
		"index/plugin",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"kong",
		"index/route",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"kong",
		"index/service",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"kong",
		"index/sni",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"kong",
		"index/target",
		&module{version},
	)
	pulumi.RegisterResourceModule(
		"kong",
		"index/upstream",
		&module{version},
	)
	pulumi.RegisterResourcePackage(
		"kong",
		&pkg{version},
	)
}