module github.com/pulumi/pulumi-kong/provider/v3

go 1.16

require (
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/kevholditch/terraform-provider-kong v1.9.2-0.20200124095244-a53d2fc45429
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.0.0-beta.1
	github.com/pulumi/pulumi/sdk/v3 v3.0.0-beta.2
)

replace (
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
	github.com/kevholditch/terraform-provider-kong => github.com/pulumi/terraform-provider-kong v1.9.2-0.20200421202609-106afef2f0dc
)
