package ec2verifiedaccessgroup


type Ec2VerifiedAccessGroupSseSpecification struct {
	// Whether to encrypt the policy with the provided key or disable encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_verified_access_group#customer_managed_key_enabled Ec2VerifiedAccessGroup#customer_managed_key_enabled}
	CustomerManagedKeyEnabled interface{} `field:"optional" json:"customerManagedKeyEnabled" yaml:"customerManagedKeyEnabled"`
	// KMS Key Arn used to encrypt the group policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_verified_access_group#kms_key_arn Ec2VerifiedAccessGroup#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

