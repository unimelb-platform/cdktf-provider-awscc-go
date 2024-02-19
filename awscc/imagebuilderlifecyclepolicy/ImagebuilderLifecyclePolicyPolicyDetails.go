package imagebuilderlifecyclepolicy


type ImagebuilderLifecyclePolicyPolicyDetails struct {
	// The action of the policy detail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_lifecycle_policy#action ImagebuilderLifecyclePolicy#action}
	Action *ImagebuilderLifecyclePolicyPolicyDetailsAction `field:"required" json:"action" yaml:"action"`
	// The filters to apply of the policy detail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_lifecycle_policy#filter ImagebuilderLifecyclePolicy#filter}
	Filter *ImagebuilderLifecyclePolicyPolicyDetailsFilter `field:"required" json:"filter" yaml:"filter"`
	// The exclusion rules to apply of the policy detail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_lifecycle_policy#exclusion_rules ImagebuilderLifecyclePolicy#exclusion_rules}
	ExclusionRules *ImagebuilderLifecyclePolicyPolicyDetailsExclusionRules `field:"optional" json:"exclusionRules" yaml:"exclusionRules"`
}

