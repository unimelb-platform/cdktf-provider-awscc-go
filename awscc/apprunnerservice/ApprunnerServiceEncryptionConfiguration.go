package apprunnerservice


type ApprunnerServiceEncryptionConfiguration struct {
	// The KMS Key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_service#kms_key ApprunnerService#kms_key}
	KmsKey *string `field:"required" json:"kmsKey" yaml:"kmsKey"`
}

