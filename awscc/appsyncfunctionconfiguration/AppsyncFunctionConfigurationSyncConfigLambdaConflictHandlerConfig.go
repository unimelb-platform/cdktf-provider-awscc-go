package appsyncfunctionconfiguration


type AppsyncFunctionConfigurationSyncConfigLambdaConflictHandlerConfig struct {
	// The Amazon Resource Name (ARN) for the Lambda function to use as the Conflict Handler.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appsync_function_configuration#lambda_conflict_handler_arn AppsyncFunctionConfiguration#lambda_conflict_handler_arn}
	LambdaConflictHandlerArn *string `field:"optional" json:"lambdaConflictHandlerArn" yaml:"lambdaConflictHandlerArn"`
}

