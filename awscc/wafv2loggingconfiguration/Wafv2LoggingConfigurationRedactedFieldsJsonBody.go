package wafv2loggingconfiguration


type Wafv2LoggingConfigurationRedactedFieldsJsonBody struct {
	// The patterns to look for in the JSON body.
	//
	// AWS WAF inspects the results of these pattern matches against the rule inspection criteria.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/wafv2_logging_configuration#match_pattern Wafv2LoggingConfiguration#match_pattern}
	MatchPattern *Wafv2LoggingConfigurationRedactedFieldsJsonBodyMatchPattern `field:"required" json:"matchPattern" yaml:"matchPattern"`
	// The parts of the JSON to match against using the MatchPattern.
	//
	// If you specify All, AWS WAF matches against keys and values.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/wafv2_logging_configuration#match_scope Wafv2LoggingConfiguration#match_scope}
	MatchScope *string `field:"required" json:"matchScope" yaml:"matchScope"`
	// What AWS WAF should do if it fails to completely parse the JSON body.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/wafv2_logging_configuration#invalid_fallback_behavior Wafv2LoggingConfiguration#invalid_fallback_behavior}
	InvalidFallbackBehavior *string `field:"optional" json:"invalidFallbackBehavior" yaml:"invalidFallbackBehavior"`
}

