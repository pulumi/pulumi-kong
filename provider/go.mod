module github.com/pulumi/pulumi-kong/provider/v4

go 1.16

require (
	github.com/kevholditch/terraform-provider-kong v1.9.2-0.20220328204855-9e50bd93437f
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.21.0
	github.com/pulumi/pulumi/sdk/v3 v3.30.0
)

replace (
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20210629210550-59d24255d71f
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)
