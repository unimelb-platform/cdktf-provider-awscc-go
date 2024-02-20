package apigatewaystage


type ApigatewayStageAccessLogSetting struct {
	// The Amazon Resource Name (ARN) of the CloudWatch Logs log group or Kinesis Data Firehose delivery stream to receive access logs.
	//
	// If you specify a Kinesis Data Firehose delivery stream, the stream name must begin with ``amazon-apigateway-``. This parameter is required to enable access logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_stage#destination_arn ApigatewayStage#destination_arn}
	DestinationArn *string `field:"optional" json:"destinationArn" yaml:"destinationArn"`
	// A single line format of the access logs of data, as specified by selected [$context variables](https://docs.aws.amazon.com/apigateway/latest/developerguide/api-gateway-mapping-template-reference.html#context-variable-reference). The format must include at least ``$context.requestId``. This parameter is required to enable access logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apigateway_stage#format ApigatewayStage#format}
	Format *string `field:"optional" json:"format" yaml:"format"`
}

