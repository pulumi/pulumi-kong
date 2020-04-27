module github.com/pulumi/pulumi-kong/provider/v2

go 1.13

require (
	github.com/hashicorp/terraform-plugin-sdk v1.7.0
	github.com/kevholditch/terraform-provider-kong v1.9.2-0.20200124095244-a53d2fc45429
	github.com/pulumi/pulumi-terraform-bridge/v2 v2.1.0
	github.com/pulumi/pulumi/sdk/v2 v2.0.1-0.20200424001829-090f390d7b1a
)

replace (
	github.com/Azure/go-autorest => github.com/Azure/go-autorest v12.4.3+incompatible
	github.com/hashicorp/vault => github.com/hashicorp/vault v1.2.0
	github.com/kevholditch/terraform-provider-kong => github.com/pulumi/terraform-provider-kong v1.9.2-0.20200421202609-106afef2f0dc
)
