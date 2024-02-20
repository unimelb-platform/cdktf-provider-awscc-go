package sagemakerspace


type SagemakerSpaceSpaceSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_space#app_type SagemakerSpace#app_type}.
	AppType *string `field:"optional" json:"appType" yaml:"appType"`
	// The CodeEditor app settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_space#code_editor_app_settings SagemakerSpace#code_editor_app_settings}
	CodeEditorAppSettings *SagemakerSpaceSpaceSettingsCodeEditorAppSettings `field:"optional" json:"codeEditorAppSettings" yaml:"codeEditorAppSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_space#custom_file_systems SagemakerSpace#custom_file_systems}.
	CustomFileSystems interface{} `field:"optional" json:"customFileSystems" yaml:"customFileSystems"`
	// The JupyterLab app settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_space#jupyter_lab_app_settings SagemakerSpace#jupyter_lab_app_settings}
	JupyterLabAppSettings *SagemakerSpaceSpaceSettingsJupyterLabAppSettings `field:"optional" json:"jupyterLabAppSettings" yaml:"jupyterLabAppSettings"`
	// The Jupyter server's app settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_space#jupyter_server_app_settings SagemakerSpace#jupyter_server_app_settings}
	JupyterServerAppSettings *SagemakerSpaceSpaceSettingsJupyterServerAppSettings `field:"optional" json:"jupyterServerAppSettings" yaml:"jupyterServerAppSettings"`
	// The kernel gateway app settings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_space#kernel_gateway_app_settings SagemakerSpace#kernel_gateway_app_settings}
	KernelGatewayAppSettings *SagemakerSpaceSpaceSettingsKernelGatewayAppSettings `field:"optional" json:"kernelGatewayAppSettings" yaml:"kernelGatewayAppSettings"`
	// Default storage settings for a space.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_space#space_storage_settings SagemakerSpace#space_storage_settings}
	SpaceStorageSettings *SagemakerSpaceSpaceSettingsSpaceStorageSettings `field:"optional" json:"spaceStorageSettings" yaml:"spaceStorageSettings"`
}

