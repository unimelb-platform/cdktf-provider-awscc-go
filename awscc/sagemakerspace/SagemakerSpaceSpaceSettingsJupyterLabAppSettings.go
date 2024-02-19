package sagemakerspace


type SagemakerSpaceSpaceSettingsJupyterLabAppSettings struct {
	// A list of CodeRepositories available for use with JupyterLab apps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_space#code_repositories SagemakerSpace#code_repositories}
	CodeRepositories interface{} `field:"optional" json:"codeRepositories" yaml:"codeRepositories"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_space#default_resource_spec SagemakerSpace#default_resource_spec}.
	DefaultResourceSpec *SagemakerSpaceSpaceSettingsJupyterLabAppSettingsDefaultResourceSpec `field:"optional" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
}

