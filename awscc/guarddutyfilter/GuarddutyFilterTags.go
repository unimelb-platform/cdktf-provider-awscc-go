package guarddutyfilter


type GuarddutyFilterTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/guardduty_filter#key GuarddutyFilter#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/guardduty_filter#value GuarddutyFilter#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

