package s3bucket


type S3BucketLifecycleConfigurationRulesTransition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#storage_class S3Bucket#storage_class}.
	StorageClass *string `field:"required" json:"storageClass" yaml:"storageClass"`
	// The date value in ISO 8601 format. The timezone is always UTC. (YYYY-MM-DDThh:mm:ssZ).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#transition_date S3Bucket#transition_date}
	TransitionDate *string `field:"optional" json:"transitionDate" yaml:"transitionDate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#transition_in_days S3Bucket#transition_in_days}.
	TransitionInDays *float64 `field:"optional" json:"transitionInDays" yaml:"transitionInDays"`
}

