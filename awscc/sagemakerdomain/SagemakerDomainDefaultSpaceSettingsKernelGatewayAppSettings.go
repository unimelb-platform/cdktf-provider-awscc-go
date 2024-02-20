package sagemakerdomain


type SagemakerDomainDefaultSpaceSettingsKernelGatewayAppSettings struct {
	// A list of custom SageMaker images that are configured to run as a KernelGateway app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_domain#custom_images SagemakerDomain#custom_images}
	CustomImages interface{} `field:"optional" json:"customImages" yaml:"customImages"`
	// The default instance type and the Amazon Resource Name (ARN) of the default SageMaker image used by the KernelGateway app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_domain#default_resource_spec SagemakerDomain#default_resource_spec}
	DefaultResourceSpec *SagemakerDomainDefaultSpaceSettingsKernelGatewayAppSettingsDefaultResourceSpec `field:"optional" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
}

