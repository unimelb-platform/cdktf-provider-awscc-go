package dynamodbtable


type DynamodbTableImportSourceSpecificationS3BucketSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#s3_bucket DynamodbTable#s3_bucket}.
	S3Bucket *string `field:"required" json:"s3Bucket" yaml:"s3Bucket"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#s3_bucket_owner DynamodbTable#s3_bucket_owner}.
	S3BucketOwner *string `field:"optional" json:"s3BucketOwner" yaml:"s3BucketOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#s3_key_prefix DynamodbTable#s3_key_prefix}.
	S3KeyPrefix *string `field:"optional" json:"s3KeyPrefix" yaml:"s3KeyPrefix"`
}

