package s3bucket


type S3BucketObjectLockConfigurationRule struct {
	// The default retention period that you want to apply to new objects placed in the specified bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#default_retention S3Bucket#default_retention}
	DefaultRetention *S3BucketObjectLockConfigurationRuleDefaultRetention `field:"optional" json:"defaultRetention" yaml:"defaultRetention"`
}

