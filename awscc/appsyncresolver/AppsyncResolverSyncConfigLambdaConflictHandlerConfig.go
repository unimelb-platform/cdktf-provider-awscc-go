package appsyncresolver


type AppsyncResolverSyncConfigLambdaConflictHandlerConfig struct {
	// The Amazon Resource Name (ARN) for the Lambda function to use as the Conflict Handler.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appsync_resolver#lambda_conflict_handler_arn AppsyncResolver#lambda_conflict_handler_arn}
	LambdaConflictHandlerArn *string `field:"optional" json:"lambdaConflictHandlerArn" yaml:"lambdaConflictHandlerArn"`
}

