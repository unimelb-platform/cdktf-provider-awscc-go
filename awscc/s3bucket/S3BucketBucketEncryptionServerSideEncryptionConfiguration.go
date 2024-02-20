package s3bucket


type S3BucketBucketEncryptionServerSideEncryptionConfiguration struct {
	// Specifies whether Amazon S3 should use an S3 Bucket Key with server-side encryption using KMS (SSE-KMS) for new objects in the bucket.
	//
	// Existing objects are not affected. Setting the BucketKeyEnabled element to true causes Amazon S3 to use an S3 Bucket Key. By default, S3 Bucket Key is not enabled.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#bucket_key_enabled S3Bucket#bucket_key_enabled}
	BucketKeyEnabled interface{} `field:"optional" json:"bucketKeyEnabled" yaml:"bucketKeyEnabled"`
	// Specifies the default server-side encryption to apply to new objects in the bucket.
	//
	// If a PUT Object request doesn't specify any server-side encryption, this default encryption will be applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#server_side_encryption_by_default S3Bucket#server_side_encryption_by_default}
	ServerSideEncryptionByDefault *S3BucketBucketEncryptionServerSideEncryptionConfigurationServerSideEncryptionByDefault `field:"optional" json:"serverSideEncryptionByDefault" yaml:"serverSideEncryptionByDefault"`
}

