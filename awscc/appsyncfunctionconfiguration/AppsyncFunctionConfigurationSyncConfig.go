package appsyncfunctionconfiguration


type AppsyncFunctionConfigurationSyncConfig struct {
	// The Conflict Detection strategy to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appsync_function_configuration#conflict_detection AppsyncFunctionConfiguration#conflict_detection}
	ConflictDetection *string `field:"required" json:"conflictDetection" yaml:"conflictDetection"`
	// The Conflict Resolution strategy to perform in the event of a conflict.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appsync_function_configuration#conflict_handler AppsyncFunctionConfiguration#conflict_handler}
	ConflictHandler *string `field:"optional" json:"conflictHandler" yaml:"conflictHandler"`
	// The LambdaConflictHandlerConfig when configuring LAMBDA as the Conflict Handler.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appsync_function_configuration#lambda_conflict_handler_config AppsyncFunctionConfiguration#lambda_conflict_handler_config}
	LambdaConflictHandlerConfig *AppsyncFunctionConfigurationSyncConfigLambdaConflictHandlerConfig `field:"optional" json:"lambdaConflictHandlerConfig" yaml:"lambdaConflictHandlerConfig"`
}

