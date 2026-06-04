package xraysamplingrule


type XraySamplingRuleTags struct {
	// The key name of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/xray_sampling_rule#key XraySamplingRule#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/xray_sampling_rule#value XraySamplingRule#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

