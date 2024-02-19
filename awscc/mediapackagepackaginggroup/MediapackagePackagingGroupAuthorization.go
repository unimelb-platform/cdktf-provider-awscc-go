package mediapackagepackaginggroup


type MediapackagePackagingGroupAuthorization struct {
	// The Amazon Resource Name (ARN) for the secret in AWS Secrets Manager that is used for CDN authorization.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_packaging_group#cdn_identifier_secret MediapackagePackagingGroup#cdn_identifier_secret}
	CdnIdentifierSecret *string `field:"required" json:"cdnIdentifierSecret" yaml:"cdnIdentifierSecret"`
	// The Amazon Resource Name (ARN) for the IAM role that allows MediaPackage to communicate with AWS Secrets Manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_packaging_group#secrets_role_arn MediapackagePackagingGroup#secrets_role_arn}
	SecretsRoleArn *string `field:"required" json:"secretsRoleArn" yaml:"secretsRoleArn"`
}

