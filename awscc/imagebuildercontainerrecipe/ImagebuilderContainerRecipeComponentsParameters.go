package imagebuildercontainerrecipe


type ImagebuilderContainerRecipeComponentsParameters struct {
	// The name of the component parameter to set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_container_recipe#name ImagebuilderContainerRecipe#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Sets the value for the named component parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_container_recipe#value ImagebuilderContainerRecipe#value}
	Value *[]*string `field:"required" json:"value" yaml:"value"`
}

