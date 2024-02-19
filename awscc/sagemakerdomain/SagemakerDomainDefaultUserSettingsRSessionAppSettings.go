package sagemakerdomain


type SagemakerDomainDefaultUserSettingsRSessionAppSettings struct {
	// A list of custom SageMaker images that are configured to run as a KernelGateway app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_domain#custom_images SagemakerDomain#custom_images}
	CustomImages interface{} `field:"optional" json:"customImages" yaml:"customImages"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_domain#default_resource_spec SagemakerDomain#default_resource_spec}.
	DefaultResourceSpec *SagemakerDomainDefaultUserSettingsRSessionAppSettingsDefaultResourceSpec `field:"optional" json:"defaultResourceSpec" yaml:"defaultResourceSpec"`
}

