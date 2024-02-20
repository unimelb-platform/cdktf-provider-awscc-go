package ecrrepository


type EcrRepositoryEncryptionConfiguration struct {
	// The encryption type to use.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecr_repository#encryption_type EcrRepository#encryption_type}
	EncryptionType *string `field:"required" json:"encryptionType" yaml:"encryptionType"`
	// If you use the KMS encryption type, specify the CMK to use for encryption.
	//
	// The alias, key ID, or full ARN of the CMK can be specified. The key must exist in the same Region as the repository. If no key is specified, the default AWS managed CMK for Amazon ECR will be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecr_repository#kms_key EcrRepository#kms_key}
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
}

