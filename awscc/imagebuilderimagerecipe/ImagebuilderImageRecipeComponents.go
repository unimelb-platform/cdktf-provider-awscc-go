package imagebuilderimagerecipe


type ImagebuilderImageRecipeComponents struct {
	// The Amazon Resource Name (ARN) of the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_image_recipe#component_arn ImagebuilderImageRecipe#component_arn}
	ComponentArn *string `field:"optional" json:"componentArn" yaml:"componentArn"`
	// A group of parameter settings that are used to configure the component for a specific recipe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_image_recipe#parameters ImagebuilderImageRecipe#parameters}
	Parameters interface{} `field:"optional" json:"parameters" yaml:"parameters"`
}

