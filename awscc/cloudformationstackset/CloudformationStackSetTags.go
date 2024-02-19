package cloudformationstackset


type CloudformationStackSetTags struct {
	// A string used to identify this tag. You can specify a maximum of 127 characters for a tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudformation_stack_set#key CloudformationStackSet#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// A string containing the value for this tag.
	//
	// You can specify a maximum of 256 characters for a tag value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudformation_stack_set#value CloudformationStackSet#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

