package applicationinsightsapplication


type ApplicationinsightsApplicationComponentMonitoringSettingsDefaultOverwriteComponentConfigurationConfigurationDetailsLogs struct {
	// The log type decides the log patterns against which Application Insights analyzes the log.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#log_type ApplicationinsightsApplication#log_type}
	LogType *string `field:"required" json:"logType" yaml:"logType"`
	// The type of encoding of the logs to be monitored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#encoding ApplicationinsightsApplication#encoding}
	Encoding *string `field:"optional" json:"encoding" yaml:"encoding"`
	// The CloudWatch log group name to be associated to the monitored log.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#log_group_name ApplicationinsightsApplication#log_group_name}
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
	// The path of the logs to be monitored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#log_path ApplicationinsightsApplication#log_path}
	LogPath *string `field:"optional" json:"logPath" yaml:"logPath"`
	// The name of the log pattern set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/applicationinsights_application#pattern_set ApplicationinsightsApplication#pattern_set}
	PatternSet *string `field:"optional" json:"patternSet" yaml:"patternSet"`
}

