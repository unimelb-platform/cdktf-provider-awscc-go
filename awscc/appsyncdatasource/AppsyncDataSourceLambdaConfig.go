package appsyncdatasource


type AppsyncDataSourceLambdaConfig struct {
	// The ARN for the Lambda function.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/appsync_data_source#lambda_function_arn AppsyncDataSource#lambda_function_arn}
	LambdaFunctionArn *string `field:"optional" json:"lambdaFunctionArn" yaml:"lambdaFunctionArn"`
}

