package s3bucket


type S3BucketNotificationConfiguration struct {
	// Describes the Amazon EventBridge notification configuration for an Amazon S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#event_bridge_configuration S3Bucket#event_bridge_configuration}
	EventBridgeConfiguration *S3BucketNotificationConfigurationEventBridgeConfiguration `field:"optional" json:"eventBridgeConfiguration" yaml:"eventBridgeConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#lambda_configurations S3Bucket#lambda_configurations}.
	LambdaConfigurations interface{} `field:"optional" json:"lambdaConfigurations" yaml:"lambdaConfigurations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#queue_configurations S3Bucket#queue_configurations}.
	QueueConfigurations interface{} `field:"optional" json:"queueConfigurations" yaml:"queueConfigurations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#topic_configurations S3Bucket#topic_configurations}.
	TopicConfigurations interface{} `field:"optional" json:"topicConfigurations" yaml:"topicConfigurations"`
}

