package wafv2loggingconfiguration


type Wafv2LoggingConfigurationLoggingFilter struct {
	// Default handling for logs that don't match any of the specified filtering conditions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/wafv2_logging_configuration#default_behavior Wafv2LoggingConfiguration#default_behavior}
	DefaultBehavior *string `field:"required" json:"defaultBehavior" yaml:"defaultBehavior"`
	// The filters that you want to apply to the logs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/wafv2_logging_configuration#filters Wafv2LoggingConfiguration#filters}
	Filters interface{} `field:"required" json:"filters" yaml:"filters"`
}

