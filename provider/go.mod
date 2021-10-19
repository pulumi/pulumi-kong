module github.com/pulumi/pulumi-kong/provider/v4

go 1.16

require (
	github.com/kevholditch/terraform-provider-kong v1.9.2-0.20211008132823-82d072eec576
	github.com/pulumi/pulumi-terraform-bridge/v3 v3.0.0
	github.com/pulumi/pulumi/sdk/v3 v3.0.0
)

replace (
	github.com/hashicorp/terraform-plugin-sdk/v2 => github.com/pulumi/terraform-plugin-sdk/v2 v2.0.0-20210402103405-f5979773e8ba
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
)
