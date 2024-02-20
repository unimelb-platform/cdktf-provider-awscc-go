package imagebuilderlifecyclepolicy


type ImagebuilderLifecyclePolicyResourceSelection struct {
	// The recipes to select.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_lifecycle_policy#recipes ImagebuilderLifecyclePolicy#recipes}
	Recipes interface{} `field:"optional" json:"recipes" yaml:"recipes"`
	// The Image Builder resources to select by tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_lifecycle_policy#tag_map ImagebuilderLifecyclePolicy#tag_map}
	TagMap *map[string]*string `field:"optional" json:"tagMap" yaml:"tagMap"`
}

