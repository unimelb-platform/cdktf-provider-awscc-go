package lexbot


type LexBotBotLocalesSlotTypesExternalSourceSettingGrammarSlotTypeSettingSource struct {
	// The name of the S3 bucket that contains the grammar source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#s3_bucket_name LexBot#s3_bucket_name}
	S3BucketName *string `field:"required" json:"s3BucketName" yaml:"s3BucketName"`
	// The path to the grammar in the S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#s3_object_key LexBot#s3_object_key}
	S3ObjectKey *string `field:"required" json:"s3ObjectKey" yaml:"s3ObjectKey"`
	// The Amazon KMS key required to decrypt the contents of the grammar, if any.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#kms_key_arn LexBot#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

