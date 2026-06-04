package vpclatticerule


type VpclatticeRuleMatch struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/vpclattice_rule#http_match VpclatticeRule#http_match}.
	HttpMatch *VpclatticeRuleMatchHttpMatch `field:"required" json:"httpMatch" yaml:"httpMatch"`
}

