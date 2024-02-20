package nimblestudiostudio


type NimblestudioStudioStudioEncryptionConfiguration struct {
	// <p>The type of KMS key that is used to encrypt studio data.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio#key_type NimblestudioStudio#key_type}
	KeyType *string `field:"required" json:"keyType" yaml:"keyType"`
	// <p>The ARN for a KMS key that is used to encrypt studio data.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio#key_arn NimblestudioStudio#key_arn}
	KeyArn *string `field:"optional" json:"keyArn" yaml:"keyArn"`
}

