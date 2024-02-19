package ecsservice


type EcsServiceServiceConnectConfigurationServicesTls struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_service#issuer_certificate_authority EcsService#issuer_certificate_authority}.
	IssuerCertificateAuthority *EcsServiceServiceConnectConfigurationServicesTlsIssuerCertificateAuthority `field:"required" json:"issuerCertificateAuthority" yaml:"issuerCertificateAuthority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_service#kms_key EcsService#kms_key}.
	KmsKey *string `field:"optional" json:"kmsKey" yaml:"kmsKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_service#role_arn EcsService#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

