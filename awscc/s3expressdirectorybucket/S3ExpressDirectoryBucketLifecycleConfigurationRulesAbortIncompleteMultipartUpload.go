package s3expressdirectorybucket


type S3ExpressDirectoryBucketLifecycleConfigurationRulesAbortIncompleteMultipartUpload struct {
	// Specifies the number of days after which Amazon S3 aborts an incomplete multipart upload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/s3express_directory_bucket#days_after_initiation S3ExpressDirectoryBucket#days_after_initiation}
	DaysAfterInitiation *float64 `field:"optional" json:"daysAfterInitiation" yaml:"daysAfterInitiation"`
}

