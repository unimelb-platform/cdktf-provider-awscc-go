package mediapackageoriginendpoint


type MediapackageOriginEndpointCmafPackageEncryptionSpekeKeyProvider struct {
	// The resource ID to include in key requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#resource_id MediapackageOriginEndpoint#resource_id}
	ResourceId *string `field:"required" json:"resourceId" yaml:"resourceId"`
	// An Amazon Resource Name (ARN) of an IAM role that AWS Elemental MediaPackage will assume when accessing the key provider service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#role_arn MediapackageOriginEndpoint#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// The system IDs to include in key requests.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#system_ids MediapackageOriginEndpoint#system_ids}
	SystemIds *[]*string `field:"required" json:"systemIds" yaml:"systemIds"`
	// The URL of the external key provider service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#url MediapackageOriginEndpoint#url}
	Url *string `field:"required" json:"url" yaml:"url"`
	// An Amazon Resource Name (ARN) of a Certificate Manager certificate that MediaPackage will use for enforcing secure end-to-end data transfer with the key provider service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#certificate_arn MediapackageOriginEndpoint#certificate_arn}
	CertificateArn *string `field:"optional" json:"certificateArn" yaml:"certificateArn"`
	// The configuration to use for encrypting one or more content tracks separately for endpoints that use SPEKE 2.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/mediapackage_origin_endpoint#encryption_contract_configuration MediapackageOriginEndpoint#encryption_contract_configuration}
	EncryptionContractConfiguration *MediapackageOriginEndpointCmafPackageEncryptionSpekeKeyProviderEncryptionContractConfiguration `field:"optional" json:"encryptionContractConfiguration" yaml:"encryptionContractConfiguration"`
}

