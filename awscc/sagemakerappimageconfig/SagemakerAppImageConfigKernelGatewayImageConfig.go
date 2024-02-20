package sagemakerappimageconfig


type SagemakerAppImageConfigKernelGatewayImageConfig struct {
	// The specification of the Jupyter kernels in the image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_app_image_config#kernel_specs SagemakerAppImageConfig#kernel_specs}
	KernelSpecs interface{} `field:"required" json:"kernelSpecs" yaml:"kernelSpecs"`
	// The Amazon Elastic File System (EFS) storage configuration for a SageMaker image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_app_image_config#file_system_config SagemakerAppImageConfig#file_system_config}
	FileSystemConfig *SagemakerAppImageConfigKernelGatewayImageConfigFileSystemConfig `field:"optional" json:"fileSystemConfig" yaml:"fileSystemConfig"`
}

