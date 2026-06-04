package kinesisanalyticsv2application


type Kinesisanalyticsv2ApplicationApplicationConfigurationSqlApplicationConfigurationInputsKinesisStreamsInput struct {
	// The ARN of the input Kinesis data stream to read.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/kinesisanalyticsv2_application#resource_arn Kinesisanalyticsv2Application#resource_arn}
	ResourceArn *string `field:"optional" json:"resourceArn" yaml:"resourceArn"`
}

