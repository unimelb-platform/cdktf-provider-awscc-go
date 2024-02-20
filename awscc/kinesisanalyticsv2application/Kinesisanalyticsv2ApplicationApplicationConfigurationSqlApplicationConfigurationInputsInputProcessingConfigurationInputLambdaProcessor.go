package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputsInputProcessingConfigurationInputLambdaProcessor struct {
	// The ARN of the Amazon Lambda function that operates on records in the stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisanalyticsv2_application#resource_arn Kinesisanalyticsv2Application#resource_arn}
	ResourceArn *string `field:"required" json:"resourceArn" yaml:"resourceArn"`
}

