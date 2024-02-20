package sagemakeruserprofile


type SagemakerUserProfileUserSettingsJupyterLabAppSettings struct {
	// A list of CodeRepositories available for use with JupyterLab apps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_user_profile#code_repositories SagemakerUserProfile#code_repositories}
	CodeRepositories interface{} `field:"optional" json:"codeRepositories" yaml:"codeRepositories"`
	// A list of custom images available for use for JupyterLab apps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_user_profile#custom_images SagemakerUserProfile#custom_images}
	CustomImages interface{} `field:"optional" json:"customImages" yaml:"customImages"`
	// The default instance type and the Amazon Resource Name (ARN) of the default SageMaker image used by the JupyterLab app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_user_profile#default_resource_spec SagemakerUserProfile#default_resource_spec}
	DefaultResourceSpec *SagemakerUserProfileUserSettingsJupyterLabAppSettingsDefaultResourceSpec `field:"optional" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
	// A list of LifecycleConfigArns available for use with JupyterLab apps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_user_profile#lifecycle_config_arns SagemakerUserProfile#lifecycle_config_arns}
	LifecycleConfigArns *[]*string `field:"optional" json:"lifecycleConfigArns" yaml:"lifecycleConfigArns"`
}

