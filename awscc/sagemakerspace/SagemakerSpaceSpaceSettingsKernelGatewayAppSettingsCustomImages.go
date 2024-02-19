package sagemakerspace


type SagemakerSpaceSpaceSettingsKernelGatewayAppSettingsCustomImages struct {
	// The Name of the AppImageConfig.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_space#app_image_config_name SagemakerSpace#app_image_config_name}
	AppImageConfigName *string `field:"required" json:"appImageConfigName" yaml:"appImageConfigName"`
	// The name of the CustomImage. Must be unique to your account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_space#image_name SagemakerSpace#image_name}
	ImageName *string `field:"required" json:"imageName" yaml:"imageName"`
	// The version number of the CustomImage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_space#image_version_number SagemakerSpace#image_version_number}
	ImageVersionNumber *float64 `field:"optional" json:"imageVersionNumber" yaml:"imageVersionNumber"`
}

