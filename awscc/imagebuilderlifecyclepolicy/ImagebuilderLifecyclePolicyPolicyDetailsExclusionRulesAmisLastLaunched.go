package imagebuilderlifecyclepolicy


type ImagebuilderLifecyclePolicyPolicyDetailsExclusionRulesAmisLastLaunched struct {
	// The value's time unit.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_lifecycle_policy#unit ImagebuilderLifecyclePolicy#unit}
	Unit *string `field:"required" json:"unit" yaml:"unit"`
	// The last launched value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_lifecycle_policy#value ImagebuilderLifecyclePolicy#value}
	Value *float64 `field:"required" json:"value" yaml:"value"`
}

