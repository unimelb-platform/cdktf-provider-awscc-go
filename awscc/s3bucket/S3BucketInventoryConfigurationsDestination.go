package s3bucket


type S3BucketInventoryConfigurationsDestination struct {
	// The Amazon Resource Name (ARN) of the bucket to which data is exported.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#bucket_arn S3Bucket#bucket_arn}
	BucketArn *string `field:"required" json:"bucketArn" yaml:"bucketArn"`
	// Specifies the file format used when exporting data to Amazon S3.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#format S3Bucket#format}
	Format *string `field:"required" json:"format" yaml:"format"`
	// The account ID that owns the destination S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#bucket_account_id S3Bucket#bucket_account_id}
	BucketAccountId *string `field:"optional" json:"bucketAccountId" yaml:"bucketAccountId"`
	// The prefix to use when exporting data. The prefix is prepended to all results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#prefix S3Bucket#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

