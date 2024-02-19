package cloudfrontpublickey


type CloudfrontPublicKeyPublicKeyConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_public_key#caller_reference CloudfrontPublicKey#caller_reference}.
	CallerReference *string `field:"required" json:"callerReference" yaml:"callerReference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_public_key#encoded_key CloudfrontPublicKey#encoded_key}.
	EncodedKey *string `field:"required" json:"encodedKey" yaml:"encodedKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_public_key#name CloudfrontPublicKey#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cloudfront_public_key#comment CloudfrontPublicKey#comment}.
	Comment *string `field:"optional" json:"comment" yaml:"comment"`
}

