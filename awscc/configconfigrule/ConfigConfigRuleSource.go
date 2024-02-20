package configconfigrule


type ConfigConfigRuleSource struct {
	// Owner of the config rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_config_rule#owner ConfigConfigRule#owner}
	Owner *string `field:"required" json:"owner" yaml:"owner"`
	// Custom policy details when rule is custom owned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_config_rule#custom_policy_details ConfigConfigRule#custom_policy_details}
	CustomPolicyDetails *ConfigConfigRuleSourceCustomPolicyDetails `field:"optional" json:"customPolicyDetails" yaml:"customPolicyDetails"`
	// List of message types that can trigger the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_config_rule#source_details ConfigConfigRule#source_details}
	SourceDetails interface{} `field:"optional" json:"sourceDetails" yaml:"sourceDetails"`
	// Identifier for the source of events.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_config_rule#source_identifier ConfigConfigRule#source_identifier}
	SourceIdentifier *string `field:"optional" json:"sourceIdentifier" yaml:"sourceIdentifier"`
}

