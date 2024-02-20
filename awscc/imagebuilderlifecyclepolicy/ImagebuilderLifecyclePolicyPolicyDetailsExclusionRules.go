package imagebuilderlifecyclepolicy


type ImagebuilderLifecyclePolicyPolicyDetailsExclusionRules struct {
	// The AMI exclusion rules for the policy detail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_lifecycle_policy#amis ImagebuilderLifecyclePolicy#amis}
	Amis *ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmis `field:"optional" json:"amis" yaml:"amis"`
	// The Image Builder tags to filter on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_lifecycle_policy#tag_map ImagebuilderLifecyclePolicy#tag_map}
	TagMap *map[string]*string `field:"optional" json:"tagMap" yaml:"tagMap"`
}

