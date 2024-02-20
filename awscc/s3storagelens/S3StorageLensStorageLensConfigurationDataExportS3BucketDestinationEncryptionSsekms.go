package s3storagelens


type S3StorageLensStorageLensConfigurationDataExportS3BucketDestinationEncryptionSsekms struct {
	// The ARN of the KMS key to use for encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/s3_storage_lens#key_id S3StorageLens#key_id}
	KeyId *string `field:"required" json:"keyId" yaml:"keyId"`
}

