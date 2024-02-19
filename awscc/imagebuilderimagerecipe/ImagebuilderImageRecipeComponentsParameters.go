package imagebuilderimagerecipe


type ImagebuilderImageRecipeComponentsParameters struct {
	// The name of the component parameter to set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_image_recipe#name ImagebuilderImageRecipe#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Sets the value for the named component parameter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_image_recipe#value ImagebuilderImageRecipe#value}
	Value *[]*string `field:"required" json:"value" yaml:"value"`
}

