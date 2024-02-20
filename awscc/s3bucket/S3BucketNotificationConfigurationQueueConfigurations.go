package s3bucket


type S3BucketNotificationConfigurationQueueConfigurations struct {
	// The Amazon S3 bucket event about which you want to publish messages to Amazon SQS.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#event S3Bucket#event}
	Event *string `field:"required" json:"event" yaml:"event"`
	// The Amazon Resource Name (ARN) of the Amazon SQS queue to which Amazon S3 publishes a message when it detects events of the specified type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#queue S3Bucket#queue}
	Queue *string `field:"required" json:"queue" yaml:"queue"`
	// The filtering rules that determine which objects trigger notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#filter S3Bucket#filter}
	Filter *S3BucketNotificationConfigurationQueueConfigurationsFilter `field:"optional" json:"filter" yaml:"filter"`
}

