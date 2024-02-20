package apprunnerservice


type ApprunnerServiceSourceConfigurationImageRepository struct {
	// Image Identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_service#image_identifier ApprunnerService#image_identifier}
	ImageIdentifier *string `field:"required" json:"imageIdentifier" yaml:"imageIdentifier"`
	// Image Repository Type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_service#image_repository_type ApprunnerService#image_repository_type}
	ImageRepositoryType *string `field:"required" json:"imageRepositoryType" yaml:"imageRepositoryType"`
	// Image Configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/apprunner_service#image_configuration ApprunnerService#image_configuration}
	ImageConfiguration *ApprunnerServiceSourceConfigurationImageRepositoryImageConfiguration `field:"optional" json:"imageConfiguration" yaml:"imageConfiguration"`
}

