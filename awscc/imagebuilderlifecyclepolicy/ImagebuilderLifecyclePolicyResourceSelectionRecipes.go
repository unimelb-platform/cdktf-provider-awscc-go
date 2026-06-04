package imagebuilderlifecyclepolicy


type ImagebuilderLifecyclePolicyResourceSelectionRecipes struct {
	// The recipe name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/imagebuilder_lifecycle_policy#name ImagebuilderLifecyclePolicy#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// The recipe version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/imagebuilder_lifecycle_policy#semantic_version ImagebuilderLifecyclePolicy#semantic_version}
	SemanticVersion *string `field:"optional" json:"semanticVersion" yaml:"semanticVersion"`
}

