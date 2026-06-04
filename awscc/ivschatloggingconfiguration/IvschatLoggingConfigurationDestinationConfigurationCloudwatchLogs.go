package ivschatloggingconfiguration


type IvschatLoggingConfigurationDestinationConfigurationCloudwatchLogs struct {
	// Name of the Amazon CloudWatch Logs log group where chat activity will be logged.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ivschat_logging_configuration#log_group_name IvschatLoggingConfiguration#log_group_name}
	LogGroupName *string `field:"optional" json:"logGroupName" yaml:"logGroupName"`
}

