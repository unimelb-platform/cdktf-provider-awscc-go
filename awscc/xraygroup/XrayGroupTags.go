package xraygroup


type XrayGroupTags struct {
	// The key name of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/xray_group#key XrayGroup#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/xray_group#value XrayGroup#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

