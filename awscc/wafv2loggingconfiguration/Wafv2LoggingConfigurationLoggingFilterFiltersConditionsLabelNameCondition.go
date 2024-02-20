package wafv2loggingconfiguration


type Wafv2LoggingConfigurationLoggingFilterFiltersConditionsLabelNameCondition struct {
	// The label name that a log record must contain in order to meet the condition.
	//
	// This must be a fully qualified label name. Fully qualified labels have a prefix, optional namespaces, and label name. The prefix identifies the rule group or web ACL context of the rule that added the label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/wafv2_logging_configuration#label_name Wafv2LoggingConfiguration#label_name}
	LabelName *string `field:"required" json:"labelName" yaml:"labelName"`
}

