package s3bucket


type S3BucketReplicationConfigurationRulesDestinationMetrics struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#status S3Bucket#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#event_threshold S3Bucket#event_threshold}.
	EventThreshold *S3BucketReplicationConfigurationRulesDestinationMetricsEventThreshold `field:"optional" json:"eventThreshold" yaml:"eventThreshold"`
}

