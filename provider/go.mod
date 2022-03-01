module github.com/pulumi/pulumi-kong/provider/v4

go 1.16

require (
	github.com/kevholditch/terraform-provider-kong v1.9.2-0.20220208144503-ca93cf0d1a10
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.19.1
	github.com/pulumi/pulumi/sdk/v3 v3.25.0
)

replace (
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20210629210550-59d24255d71f
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)
