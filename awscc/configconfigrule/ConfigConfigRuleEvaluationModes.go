package configconfigrule


type ConfigConfigRuleEvaluationModes struct {
	// Mode of evaluation of AWS Config rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_config_rule#mode ConfigConfigRule#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

