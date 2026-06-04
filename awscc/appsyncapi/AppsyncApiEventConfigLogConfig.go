package appsyncapi


type AppsyncApiEventConfigLogConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/appsync_api#cloudwatch_logs_role_arn AppsyncApi#cloudwatch_logs_role_arn}.
	CloudwatchLogsRoleArn *string `field:"optional" json:"cloudwatchLogsRoleArn" yaml:"cloudwatchLogsRoleArn"`
	// Logging level for the AppSync API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/appsync_api#log_level AppsyncApi#log_level}
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
}

