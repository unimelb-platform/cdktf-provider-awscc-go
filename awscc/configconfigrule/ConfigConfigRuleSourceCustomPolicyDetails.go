package configconfigrule


type ConfigConfigRuleSourceCustomPolicyDetails struct {
	// Logging toggle for custom policy rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_config_rule#enable_debug_log_delivery ConfigConfigRule#enable_debug_log_delivery}
	EnableDebugLogDelivery interface{} `field:"optional" json:"enableDebugLogDelivery" yaml:"enableDebugLogDelivery"`
	// Runtime system for custom policy rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_config_rule#policy_runtime ConfigConfigRule#policy_runtime}
	PolicyRuntime *string `field:"optional" json:"policyRuntime" yaml:"policyRuntime"`
	// Policy definition containing logic for custom policy rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_config_rule#policy_text ConfigConfigRule#policy_text}
	PolicyText *string `field:"optional" json:"policyText" yaml:"policyText"`
}

