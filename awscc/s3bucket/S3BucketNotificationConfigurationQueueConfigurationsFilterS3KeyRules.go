package s3bucket


type S3BucketNotificationConfigurationQueueConfigurationsFilterS3KeyRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#name S3Bucket#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#value S3Bucket#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

