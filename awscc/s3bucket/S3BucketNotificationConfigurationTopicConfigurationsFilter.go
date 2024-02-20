package s3bucket


type S3BucketNotificationConfigurationTopicConfigurationsFilter struct {
	// A container for object key name prefix and suffix filtering rules.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#s3_key S3Bucket#s3_key}
	S3Key *S3BucketNotificationConfigurationTopicConfigurationsFilterS3Key `field:"required" json:"s3Key" yaml:"s3Key"`
}

