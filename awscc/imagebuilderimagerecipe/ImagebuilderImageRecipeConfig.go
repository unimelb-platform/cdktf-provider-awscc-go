package imagebuilderimagerecipe

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ImagebuilderImageRecipeConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The components of the image recipe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_image_recipe#components ImagebuilderImageRecipe#components}
	Components interface{} `field:"required" json:"components" yaml:"components"`
	// The name of the image recipe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_image_recipe#name ImagebuilderImageRecipe#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The parent image of the image recipe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_image_recipe#parent_image ImagebuilderImageRecipe#parent_image}
	ParentImage *string `field:"required" json:"parentImage" yaml:"parentImage"`
	// The version of the image recipe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_image_recipe#version ImagebuilderImageRecipe#version}
	Version *string `field:"required" json:"version" yaml:"version"`
	// Specify additional settings and launch scripts for your build instances.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_image_recipe#additional_instance_configuration ImagebuilderImageRecipe#additional_instance_configuration}
	AdditionalInstanceConfiguration *ImagebuilderImageRecipeAdditionalInstanceConfiguration `field:"optional" json:"additionalInstanceConfiguration" yaml:"additionalInstanceConfiguration"`
	// The block device mappings to apply when creating images from this recipe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_image_recipe#block_device_mappings ImagebuilderImageRecipe#block_device_mappings}
	BlockDeviceMappings interface{} `field:"optional" json:"blockDeviceMappings" yaml:"blockDeviceMappings"`
	// The description of the image recipe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_image_recipe#description ImagebuilderImageRecipe#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// The tags of the image recipe.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_image_recipe#tags ImagebuilderImageRecipe#tags}
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// The working directory to be used during build and test workflows.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_image_recipe#working_directory ImagebuilderImageRecipe#working_directory}
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}

