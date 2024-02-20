package appflowconnector


type AppflowConnectorConnectorProvisioningConfigLambda struct {
	// Lambda ARN of the connector being registered.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_connector#lambda_arn AppflowConnector#lambda_arn}
	LambdaArn *string `field:"required" json:"lambdaArn" yaml:"lambdaArn"`
}

