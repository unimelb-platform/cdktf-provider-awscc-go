package greengrassv2componentversion


type Greengrassv2ComponentVersionLambdaFunctionComponentDependencies struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/greengrassv2_component_version#dependency_type Greengrassv2ComponentVersion#dependency_type}.
	DependencyType *string `field:"optional" json:"dependencyType" yaml:"dependencyType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/greengrassv2_component_version#version_requirement Greengrassv2ComponentVersion#version_requirement}.
	VersionRequirement *string `field:"optional" json:"versionRequirement" yaml:"versionRequirement"`
}

