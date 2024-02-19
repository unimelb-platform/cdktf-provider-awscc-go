package s3outpostsbucket


type S3OutpostsBucketLifecycleConfigurationRulesAbortIncompleteMultipartUpload struct {
	// Specifies the number of days after which Amazon S3Outposts aborts an incomplete multipart upload.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3outposts_bucket#days_after_initiation S3OutpostsBucket#days_after_initiation}
	DaysAfterInitiation *float64 `field:"required" json:"daysAfterInitiation" yaml:"daysAfterInitiation"`
}

