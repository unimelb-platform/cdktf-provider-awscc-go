package apigatewaydeployment


type ApigatewayDeploymentStageDescriptionAccessLogSetting struct {
	// The Amazon Resource Name (ARN) of the CloudWatch Logs log group or Kinesis Data Firehose delivery stream to receive access logs.
	//
	// If you specify a Kinesis Data Firehose delivery stream, the stream name must begin with ``amazon-apigateway-``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_deployment#destination_arn ApigatewayDeployment#destination_arn}
	DestinationArn *string `field:"optional" json:"destinationArn" yaml:"destinationArn"`
	// A single line format of the access logs of data, as specified by selected $context variables.
	//
	// The format must include at least ``$context.requestId``.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_deployment#format ApigatewayDeployment#format}
	Format *string `field:"optional" json:"format" yaml:"format"`
}

