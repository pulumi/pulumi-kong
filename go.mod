module github.com/pulumi/pulumi-kong

go 1.13

require (
	github.com/hashicorp/terraform v0.12.7
	github.com/kevholditch/gokong v6.0.0+incompatible // indirect
	github.com/kevholditch/terraform-provider-kong v1.9.2-0.20200124095244-a53d2fc45429
	github.com/pkg/errors v0.8.1
	github.com/pulumi/pulumi v1.6.0
	github.com/pulumi/pulumi-terraform v0.18.4-0.20191202134852-87cfb4dc8ae1
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)
