package customerprofilesintegration


type CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigSourceConnectorPropertiesS3 struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_integration#bucket_name CustomerprofilesIntegration#bucket_name}.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_integration#bucket_prefix CustomerprofilesIntegration#bucket_prefix}.
	BucketPrefix *string `field:"optional" json:"bucketPrefix" yaml:"bucketPrefix"`
}

