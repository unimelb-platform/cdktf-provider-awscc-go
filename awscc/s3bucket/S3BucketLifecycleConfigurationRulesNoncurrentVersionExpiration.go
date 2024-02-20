package s3bucket


type S3BucketLifecycleConfigurationRulesNoncurrentVersionExpiration struct {
	// Specified the number of days an object is noncurrent before Amazon S3 can perform the associated action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#noncurrent_days S3Bucket#noncurrent_days}
	NoncurrentDays *float64 `field:"required" json:"noncurrentDays" yaml:"noncurrentDays"`
	// Specified the number of newer noncurrent and current versions that must exists before performing the associated action.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#newer_noncurrent_versions S3Bucket#newer_noncurrent_versions}
	NewerNoncurrentVersions *float64 `field:"optional" json:"newerNoncurrentVersions" yaml:"newerNoncurrentVersions"`
}

