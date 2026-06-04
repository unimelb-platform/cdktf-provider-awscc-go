package ecsservice


type EcsServiceServiceConnectConfigurationServicesTlsIssuerCertificateAuthority struct {
	// The ARN of the AWS Private Certificate Authority certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ecs_service#aws_pca_authority_arn EcsService#aws_pca_authority_arn}
	AwsPcaAuthorityArn *string `field:"optional" json:"awsPcaAuthorityArn" yaml:"awsPcaAuthorityArn"`
}

