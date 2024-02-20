package elasticloadbalancingv2targetgroup


type Elasticloadbalancingv2TargetGroupTargetGroupAttributes struct {
	// The value of the attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticloadbalancingv2_target_group#key Elasticloadbalancingv2TargetGroup#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The name of the attribute.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/elasticloadbalancingv2_target_group#value Elasticloadbalancingv2TargetGroup#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

