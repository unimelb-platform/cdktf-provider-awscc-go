package imagebuildercontainerrecipe


type ImagebuilderContainerRecipeInstanceConfiguration struct {
	// Defines the block devices to attach for building an instance from this Image Builder AMI.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_container_recipe#block_device_mappings ImagebuilderContainerRecipe#block_device_mappings}
	BlockDeviceMappings interface{} `field:"optional" json:"blockDeviceMappings" yaml:"blockDeviceMappings"`
	// The AMI ID to use as the base image for a container build and test instance.
	//
	// If not specified, Image Builder will use the appropriate ECS-optimized AMI as a base image.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_container_recipe#image ImagebuilderContainerRecipe#image}
	Image *string `field:"optional" json:"image" yaml:"image"`
}

