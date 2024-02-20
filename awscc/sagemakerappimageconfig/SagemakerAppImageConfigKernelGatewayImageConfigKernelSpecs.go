package sagemakerappimageconfig


type SagemakerAppImageConfigKernelGatewayImageConfigKernelSpecs struct {
	// The name of the kernel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_app_image_config#name SagemakerAppImageConfig#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The display name of the kernel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_app_image_config#display_name SagemakerAppImageConfig#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
}

