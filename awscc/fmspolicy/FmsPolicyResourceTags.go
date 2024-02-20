package fmspolicy


type FmsPolicyResourceTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/fms_policy#key FmsPolicy#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/fms_policy#value FmsPolicy#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

