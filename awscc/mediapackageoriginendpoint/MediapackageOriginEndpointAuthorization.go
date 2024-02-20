package mediapackageoriginendpoint


type MediapackageOriginEndpointAuthorization struct {
	// The Amazon Resource Name (ARN) for the secret in Secrets Manager that your Content Distribution Network (CDN) uses for authorization to access your endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#cdn_identifier_secret MediapackageOriginEndpoint#cdn_identifier_secret}
	CdnIdentifierSecret *string `field:"required" json:"cdnIdentifierSecret" yaml:"cdnIdentifierSecret"`
	// The Amazon Resource Name (ARN) for the IAM role that allows MediaPackage to communicate with AWS Secrets Manager.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#secrets_role_arn MediapackageOriginEndpoint#secrets_role_arn}
	SecretsRoleArn *string `field:"required" json:"secretsRoleArn" yaml:"secretsRoleArn"`
}

