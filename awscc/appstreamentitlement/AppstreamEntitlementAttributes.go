package appstreamentitlement


type AppstreamEntitlementAttributes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_entitlement#name AppstreamEntitlement#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appstream_entitlement#value AppstreamEntitlement#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

