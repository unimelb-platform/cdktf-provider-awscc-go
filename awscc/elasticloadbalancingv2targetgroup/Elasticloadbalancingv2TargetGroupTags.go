package elasticloadbalancingv2targetgroup


type Elasticloadbalancingv2TargetGroupTags struct {
	// The value for the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticloadbalancingv2_target_group#key Elasticloadbalancingv2TargetGroup#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The key name of the tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticloadbalancingv2_target_group#value Elasticloadbalancingv2TargetGroup#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

