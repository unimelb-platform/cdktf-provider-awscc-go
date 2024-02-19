package lambdafunction


type LambdaFunctionLoggingConfig struct {
	// Application log granularity level, can only be used when LogFormat is set to JSON.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#application_log_level LambdaFunction#application_log_level}
	ApplicationLogLevel *string `field:"optional" json:"applicationLogLevel" yaml:"applicationLogLevel"`
	// Log delivery format for the lambda function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#log_format LambdaFunction#log_format}
	LogFormat *string `field:"optional" json:"logFormat" yaml:"logFormat"`
	// The log group name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#log_group LambdaFunction#log_group}
	LogGroup *string `field:"optional" json:"logGroup" yaml:"logGroup"`
	// System log granularity level, can only be used when LogFormat is set to JSON.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lambda_function#system_log_level LambdaFunction#system_log_level}
	SystemLogLevel *string `field:"optional" json:"systemLogLevel" yaml:"systemLogLevel"`
}

