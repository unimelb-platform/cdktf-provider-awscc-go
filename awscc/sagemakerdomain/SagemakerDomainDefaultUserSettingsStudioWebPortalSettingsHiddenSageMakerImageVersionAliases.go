package sagemakerdomain


type SagemakerDomainDefaultUserSettingsStudioWebPortalSettingsHiddenSageMakerImageVersionAliases struct {
	// The SageMaker image name that you are hiding from the Studio user interface.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_domain#sage_maker_image_name SagemakerDomain#sage_maker_image_name}
	SageMakerImageName *string `field:"optional" json:"sageMakerImageName" yaml:"sageMakerImageName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_domain#version_aliases SagemakerDomain#version_aliases}.
	VersionAliases *[]*string `field:"optional" json:"versionAliases" yaml:"versionAliases"`
}

