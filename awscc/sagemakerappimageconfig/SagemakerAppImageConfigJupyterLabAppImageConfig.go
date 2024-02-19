package sagemakerappimageconfig


type SagemakerAppImageConfigJupyterLabAppImageConfig struct {
	// The container configuration for a SageMaker image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_app_image_config#container_config SagemakerAppImageConfig#container_config}
	ContainerConfig *SagemakerAppImageConfigJupyterLabAppImageConfigContainerConfig `field:"optional" json:"containerConfig" yaml:"containerConfig"`
}

