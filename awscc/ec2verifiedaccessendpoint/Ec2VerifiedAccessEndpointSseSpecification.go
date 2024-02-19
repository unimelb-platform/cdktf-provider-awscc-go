package ec2verifiedaccessendpoint


type Ec2VerifiedAccessEndpointSseSpecification struct {
	// Whether to encrypt the policy with the provided key or disable encryption.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_verified_access_endpoint#customer_managed_key_enabled Ec2VerifiedAccessEndpoint#customer_managed_key_enabled}
	CustomerManagedKeyEnabled interface{} `field:"optional" json:"customerManagedKeyEnabled" yaml:"customerManagedKeyEnabled"`
	// KMS Key Arn used to encrypt the group policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_verified_access_endpoint#kms_key_arn Ec2VerifiedAccessEndpoint#kms_key_arn}
	KmsKeyArn *string `field:"optional" json:"kmsKeyArn" yaml:"kmsKeyArn"`
}

