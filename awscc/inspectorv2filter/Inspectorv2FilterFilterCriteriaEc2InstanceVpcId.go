package inspectorv2filter


type Inspectorv2FilterFilterCriteriaEc2InstanceVpcId struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#comparison Inspectorv2Filter#comparison}.
	Comparison *string `field:"required" json:"comparison" yaml:"comparison"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/inspectorv2_filter#value Inspectorv2Filter#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

