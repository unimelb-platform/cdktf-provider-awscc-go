package wafv2loggingconfiguration


type Wafv2LoggingConfigurationLoggingFilterFiltersConditionsActionCondition struct {
	// Logic to apply to the filtering conditions.
	//
	// You can specify that, in order to satisfy the filter, a log must match all conditions or must match at least one condition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/wafv2_logging_configuration#action Wafv2LoggingConfiguration#action}
	Action *string `field:"required" json:"action" yaml:"action"`
}

