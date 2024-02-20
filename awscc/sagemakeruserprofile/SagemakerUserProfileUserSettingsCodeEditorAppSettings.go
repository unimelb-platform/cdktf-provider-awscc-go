package sagemakeruserprofile


type SagemakerUserProfileUserSettingsCodeEditorAppSettings struct {
	// The default instance type and the Amazon Resource Name (ARN) of the default SageMaker image used by the CodeEditor app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_user_profile#default_resource_spec SagemakerUserProfile#default_resource_spec}
	DefaultResourceSpec *SagemakerUserProfileUserSettingsCodeEditorAppSettingsDefaultResourceSpec `field:"optional" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
	// A list of LifecycleConfigArns available for use with CodeEditor apps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_user_profile#lifecycle_config_arns SagemakerUserProfile#lifecycle_config_arns}
	LifecycleConfigArns *[]*string `field:"optional" json:"lifecycleConfigArns" yaml:"lifecycleConfigArns"`
}

