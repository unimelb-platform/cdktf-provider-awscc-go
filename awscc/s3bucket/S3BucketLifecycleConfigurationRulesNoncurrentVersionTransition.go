package s3bucket


type S3BucketLifecycleConfigurationRulesNoncurrentVersionTransition struct {
	// The class of storage used to store the object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#storage_class S3Bucket#storage_class}
	StorageClass *string `field:"required" json:"storageClass" yaml:"storageClass"`
	// Specifies the number of days an object is noncurrent before Amazon S3 can perform the associated action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#transition_in_days S3Bucket#transition_in_days}
	TransitionInDays *float64 `field:"required" json:"transitionInDays" yaml:"transitionInDays"`
	// Specified the number of newer noncurrent and current versions that must exists before performing the associated action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#newer_noncurrent_versions S3Bucket#newer_noncurrent_versions}
	NewerNoncurrentVersions *float64 `field:"optional" json:"newerNoncurrentVersions" yaml:"newerNoncurrentVersions"`
}

