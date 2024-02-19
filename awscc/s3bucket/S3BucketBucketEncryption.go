package s3bucket


type S3BucketBucketEncryption struct {
	// Specifies the default server-side-encryption configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_bucket#server_side_encryption_configuration S3Bucket#server_side_encryption_configuration}
	ServerSideEncryptionConfiguration interface{} `field:"required" json:"serverSideEncryptionConfiguration" yaml:"serverSideEncryptionConfiguration"`
}

