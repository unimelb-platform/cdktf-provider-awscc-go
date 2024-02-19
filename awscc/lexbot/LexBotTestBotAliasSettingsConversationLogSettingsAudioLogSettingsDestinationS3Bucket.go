package lexbot


type LexBotTestBotAliasSettingsConversationLogSettingsAudioLogSettingsDestinationS3Bucket struct {
	// The Amazon S3 key of the deployment package.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#log_prefix LexBot#log_prefix}
	LogPrefix *string `field:"required" json:"logPrefix" yaml:"logPrefix"`
	// The Amazon Resource Name (ARN) of an Amazon S3 bucket where audio log files are stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#s3_bucket_arn LexBot#s3_bucket_arn}
	S3BucketArn *string `field:"required" json:"s3BucketArn" yaml:"s3BucketArn"`
	// The Amazon Resource Name (ARN) of an AWS Key Management Service (KMS) key for encrypting audio log files stored in an S3 bucket.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/lex_bot#kms_key_arn LexBot#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

