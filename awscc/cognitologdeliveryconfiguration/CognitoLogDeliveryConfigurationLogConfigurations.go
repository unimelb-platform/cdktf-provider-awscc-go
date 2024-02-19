package cognitologdeliveryconfiguration


type CognitoLogDeliveryConfigurationLogConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cognito_log_delivery_configuration#cloudwatch_logs_configuration CognitoLogDeliveryConfiguration#cloudwatch_logs_configuration}.
	CloudwatchLogsConfiguration *CognitoLogDeliveryConfigurationLogConfigurationsCloudwatchLogsConfiguration `field:"optional" json:"cloudwatchLogsConfiguration" yaml:"cloudwatchLogsConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cognito_log_delivery_configuration#event_source CognitoLogDeliveryConfiguration#event_source}.
	EventSource *string `field:"optional" json:"eventSource" yaml:"eventSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cognito_log_delivery_configuration#log_level CognitoLogDeliveryConfiguration#log_level}.
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
}

