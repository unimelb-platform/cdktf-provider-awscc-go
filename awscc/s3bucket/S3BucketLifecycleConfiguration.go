package s3bucket


type S3BucketLifecycleConfiguration struct {
	// A lifecycle rule for individual objects in an Amazon S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#rules S3Bucket#rules}
	Rules interface{} `field:"required" json:"rules" yaml:"rules"`
}

