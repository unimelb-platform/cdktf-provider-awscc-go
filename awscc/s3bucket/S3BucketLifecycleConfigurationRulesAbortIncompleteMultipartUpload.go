package s3bucket


type S3BucketLifecycleConfigurationRulesAbortIncompleteMultipartUpload struct {
	// Specifies the number of days after which Amazon S3 aborts an incomplete multipart upload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#days_after_initiation S3Bucket#days_after_initiation}
	DaysAfterInitiation *float64 `field:"required" json:"daysAfterInitiation" yaml:"daysAfterInitiation"`
}

