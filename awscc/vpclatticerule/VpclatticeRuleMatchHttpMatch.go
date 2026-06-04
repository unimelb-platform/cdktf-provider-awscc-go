package vpclatticerule


type VpclatticeRuleMatchHttpMatch struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/vpclattice_rule#header_matches VpclatticeRule#header_matches}.
	HeaderMatches interface{} `field:"optional" json:"headerMatches" yaml:"headerMatches"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/vpclattice_rule#method VpclatticeRule#method}.
	Method *string `field:"optional" json:"method" yaml:"method"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/vpclattice_rule#path_match VpclatticeRule#path_match}.
	PathMatch *VpclatticeRuleMatchHttpMatchPathMatch `field:"optional" json:"pathMatch" yaml:"pathMatch"`
}

