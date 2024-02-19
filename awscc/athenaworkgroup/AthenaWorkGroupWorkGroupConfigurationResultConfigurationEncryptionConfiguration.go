package athenaworkgroup


type AthenaWorkGroupWorkGroupConfigurationResultConfigurationEncryptionConfiguration struct {
	// Indicates whether Amazon S3 server-side encryption with Amazon S3-managed keys (SSE-S3), server-side encryption with KMS-managed keys (SSE-KMS), or client-side encryption with KMS-managed keys (CSE-KMS) is used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/athena_work_group#encryption_option AthenaWorkGroup#encryption_option}
	EncryptionOption *string `field:"required" json:"encryptionOption" yaml:"encryptionOption"`
	// For SSE-KMS and CSE-KMS, this is the KMS key ARN or ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/athena_work_group#kms_key AthenaWorkGroup#kms_key}
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

