package configconfigrule


type ConfigConfigRuleEvaluationModes struct {
	// The mode of an evaluation. The valid values are Detective or Proactive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/config_config_rule#mode ConfigConfigRule#mode}
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
}

