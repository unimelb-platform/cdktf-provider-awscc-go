package greengrassv2componentversion


type Greengrassv2ComponentVersionLambdaFunctionComponentLambdaParametersLinuxProcessParamsContainerParamsVolumes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/greengrassv2_component_version#add_group_owner Greengrassv2ComponentVersion#add_group_owner}.
	AddGroupOwner interface{} `field:"optional" json:"addGroupOwner" yaml:"addGroupOwner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/greengrassv2_component_version#destination_path Greengrassv2ComponentVersion#destination_path}.
	DestinationPath *string `field:"optional" json:"destinationPath" yaml:"destinationPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/greengrassv2_component_version#permission Greengrassv2ComponentVersion#permission}.
	Permission *string `field:"optional" json:"permission" yaml:"permission"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/greengrassv2_component_version#source_path Greengrassv2ComponentVersion#source_path}.
	SourcePath *string `field:"optional" json:"sourcePath" yaml:"sourcePath"`
}

