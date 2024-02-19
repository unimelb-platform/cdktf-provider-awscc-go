package configconfigrule


type ConfigConfigRuleSourceSourceDetails struct {
	// Source of event that can trigger the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_config_rule#event_source ConfigConfigRule#event_source}
	EventSource *string `field:"required" json:"eventSource" yaml:"eventSource"`
	// Notification type that can trigger the rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_config_rule#message_type ConfigConfigRule#message_type}
	MessageType *string `field:"required" json:"messageType" yaml:"messageType"`
	// Frequency at which the rule has to be evaluated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/config_config_rule#maximum_execution_frequency ConfigConfigRule#maximum_execution_frequency}
	MaximumExecutionFrequency *string `field:"optional" json:"maximumExecutionFrequency" yaml:"maximumExecutionFrequency"`
}

