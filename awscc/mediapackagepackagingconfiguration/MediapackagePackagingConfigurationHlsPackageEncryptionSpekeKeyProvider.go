package mediapackagepackagingconfiguration


type MediapackagePackagingConfigurationHlsPackageEncryptionSpekeKeyProvider struct {
	// An Amazon Resource Name (ARN) of an IAM role that AWS Elemental MediaPackage will assume when accessing the key provider service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_packaging_configuration#role_arn MediapackagePackagingConfiguration#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The system IDs to include in key requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_packaging_configuration#system_ids MediapackagePackagingConfiguration#system_ids}
	SystemIds *[]*string `field:"required" json:"systemIds" yaml:"systemIds"`
	// The URL of the external key provider service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_packaging_configuration#url MediapackagePackagingConfiguration#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// The configuration to use for encrypting one or more content tracks separately for endpoints that use SPEKE 2.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_packaging_configuration#encryption_contract_configuration MediapackagePackagingConfiguration#encryption_contract_configuration}
	EncryptionContractConfiguration *MediapackagePackagingConfigurationHlsPackageEncryptionSpekeKeyProviderEncryptionContractConfiguration `field:"optional" json:"encryptionContractConfiguration" yaml:"encryptionContractConfiguration"`
}

