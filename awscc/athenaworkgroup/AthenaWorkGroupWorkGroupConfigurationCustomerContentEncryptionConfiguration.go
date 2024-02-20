package athenaworkgroup


type AthenaWorkGroupWorkGroupConfigurationCustomerContentEncryptionConfiguration struct {
	// For SSE-KMS and CSE-KMS, this is the KMS key ARN or ID.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/athena_work_group#kms_key AthenaWorkGroup#kms_key}
	KmsKey *string `field:"required" json:"kmsKey" yaml:"kmsKey"`
}

