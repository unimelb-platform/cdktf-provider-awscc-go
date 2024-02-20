package sesconfigurationseteventdestination


type SesConfigurationSetEventDestinationEventDestinationKinesisFirehoseDestination struct {
	// The ARN of the Amazon Kinesis Firehose stream that email sending events should be published to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_configuration_set_event_destination#delivery_stream_arn SesConfigurationSetEventDestination#delivery_stream_arn}
	DeliveryStreamArn *string `field:"required" json:"deliveryStreamArn" yaml:"deliveryStreamArn"`
	// The ARN of the IAM role under which Amazon SES publishes email sending events to the Amazon Kinesis Firehose stream.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ses_configuration_set_event_destination#iam_role_arn SesConfigurationSetEventDestination#iam_role_arn}
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
}

