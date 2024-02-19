package wafv2loggingconfiguration


type Wafv2LoggingConfigurationLoggingFilterFilters struct {
	// How to handle logs that satisfy the filter's conditions and requirement.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/wafv2_logging_configuration#behavior Wafv2LoggingConfiguration#behavior}
	Behavior *string `field:"required" json:"behavior" yaml:"behavior"`
	// Match conditions for the filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/wafv2_logging_configuration#conditions Wafv2LoggingConfiguration#conditions}
	Conditions interface{} `field:"required" json:"conditions" yaml:"conditions"`
	// Logic to apply to the filtering conditions.
	//
	// You can specify that, in order to satisfy the filter, a log must match all conditions or must match at least one condition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/wafv2_logging_configuration#requirement Wafv2LoggingConfiguration#requirement}
	Requirement *string `field:"required" json:"requirement" yaml:"requirement"`
}

