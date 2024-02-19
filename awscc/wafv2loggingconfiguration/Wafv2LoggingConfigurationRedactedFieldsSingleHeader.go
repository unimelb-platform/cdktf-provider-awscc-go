package wafv2loggingconfiguration


type Wafv2LoggingConfigurationRedactedFieldsSingleHeader struct {
	// The name of the query header to inspect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/wafv2_logging_configuration#name Wafv2LoggingConfiguration#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

