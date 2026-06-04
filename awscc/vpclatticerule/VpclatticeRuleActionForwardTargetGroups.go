package vpclatticerule


type VpclatticeRuleActionForwardTargetGroups struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/vpclattice_rule#target_group_identifier VpclatticeRule#target_group_identifier}.
	TargetGroupIdentifier *string `field:"optional" json:"targetGroupIdentifier" yaml:"targetGroupIdentifier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/vpclattice_rule#weight VpclatticeRule#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

