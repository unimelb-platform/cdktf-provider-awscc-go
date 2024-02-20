package sagemakerdomain


type SagemakerDomainDefaultUserSettingsCodeEditorAppSettings struct {
	// The default instance type and the Amazon Resource Name (ARN) of the default SageMaker image used by the CodeEditor app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_domain#default_resource_spec SagemakerDomain#default_resource_spec}
	DefaultResourceSpec *SagemakerDomainDefaultUserSettingsCodeEditorAppSettingsDefaultResourceSpec `field:"optional" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
	// A list of LifecycleConfigArns available for use with CodeEditor apps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_domain#lifecycle_config_arns SagemakerDomain#lifecycle_config_arns}
	LifecycleConfigArns *[]*string `field:"optional" json:"lifecycleConfigArns" yaml:"lifecycleConfigArns"`
}

