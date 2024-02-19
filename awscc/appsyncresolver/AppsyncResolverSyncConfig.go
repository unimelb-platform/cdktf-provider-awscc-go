package appsyncresolver


type AppsyncResolverSyncConfig struct {
	// The Conflict Detection strategy to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appsync_resolver#conflict_detection AppsyncResolver#conflict_detection}
	ConflictDetection *string `field:"required" json:"conflictDetection" yaml:"conflictDetection"`
	// The Conflict Resolution strategy to perform in the event of a conflict.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appsync_resolver#conflict_handler AppsyncResolver#conflict_handler}
	ConflictHandler *string `field:"optional" json:"conflictHandler" yaml:"conflictHandler"`
	// The LambdaConflictHandlerConfig when configuring LAMBDA as the Conflict Handler.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appsync_resolver#lambda_conflict_handler_config AppsyncResolver#lambda_conflict_handler_config}
	LambdaConflictHandlerConfig *AppsyncResolverSyncConfigLambdaConflictHandlerConfig `field:"optional" json:"lambdaConflictHandlerConfig" yaml:"lambdaConflictHandlerConfig"`
}

