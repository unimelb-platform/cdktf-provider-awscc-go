package sagemakeruserprofile


type SagemakerUserProfileUserSettingsKernelGatewayAppSettingsCustomImages struct {
	// The Name of the AppImageConfig.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_user_profile#app_image_config_name SagemakerUserProfile#app_image_config_name}
	AppImageConfigName *string `field:"required" json:"appImageConfigName" yaml:"appImageConfigName"`
	// The name of the CustomImage. Must be unique to your account.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_user_profile#image_name SagemakerUserProfile#image_name}
	ImageName *string `field:"required" json:"imageName" yaml:"imageName"`
	// The version number of the CustomImage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_user_profile#image_version_number SagemakerUserProfile#image_version_number}
	ImageVersionNumber *float64 `field:"optional" json:"imageVersionNumber" yaml:"imageVersionNumber"`
}

