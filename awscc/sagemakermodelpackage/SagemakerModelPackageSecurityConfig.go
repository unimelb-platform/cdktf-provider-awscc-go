package sagemakermodelpackage


type SagemakerModelPackageSecurityConfig struct {
	// The AWS KMS Key ID (KMSKeyId) used for encryption of model package information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_model_package#kms_key_id SagemakerModelPackage#kms_key_id}
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
}

