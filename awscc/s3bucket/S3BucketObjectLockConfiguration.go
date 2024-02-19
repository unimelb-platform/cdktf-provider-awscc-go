package s3bucket


type S3BucketObjectLockConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#object_lock_enabled S3Bucket#object_lock_enabled}.
	ObjectLockEnabled *string `field:"optional" json:"objectLockEnabled" yaml:"objectLockEnabled"`
	// The Object Lock rule in place for the specified object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#rule S3Bucket#rule}
	Rule *S3BucketObjectLockConfigurationRule `field:"optional" json:"rule" yaml:"rule"`
}

