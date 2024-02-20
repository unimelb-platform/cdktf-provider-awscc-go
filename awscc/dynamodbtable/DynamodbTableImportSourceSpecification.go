package dynamodbtable


type DynamodbTableImportSourceSpecification struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#input_format DynamodbTable#input_format}.
	InputFormat *string `field:"required" json:"inputFormat" yaml:"inputFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#s3_bucket_source DynamodbTable#s3_bucket_source}.
	S3BucketSource *DynamodbTableImportSourceSpecificationS3BucketSource `field:"required" json:"s3BucketSource" yaml:"s3BucketSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#input_compression_type DynamodbTable#input_compression_type}.
	InputCompressionType *string `field:"optional" json:"inputCompressionType" yaml:"inputCompressionType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/dynamodb_table#input_format_options DynamodbTable#input_format_options}.
	InputFormatOptions *DynamodbTableImportSourceSpecificationInputFormatOptions `field:"optional" json:"inputFormatOptions" yaml:"inputFormatOptions"`
}

