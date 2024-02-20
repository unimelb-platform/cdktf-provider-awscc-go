package s3bucket


type S3BucketNotificationConfigurationTopicConfigurations struct {
	// The Amazon S3 bucket event about which to send notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#event S3Bucket#event}
	Event *string `field:"required" json:"event" yaml:"event"`
	// The Amazon Resource Name (ARN) of the Amazon SNS topic to which Amazon S3 publishes a message when it detects events of the specified type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#topic S3Bucket#topic}
	Topic *string `field:"required" json:"topic" yaml:"topic"`
	// The filtering rules that determine for which objects to send notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#filter S3Bucket#filter}
	Filter *S3BucketNotificationConfigurationTopicConfigurationsFilter `field:"optional" json:"filter" yaml:"filter"`
}

