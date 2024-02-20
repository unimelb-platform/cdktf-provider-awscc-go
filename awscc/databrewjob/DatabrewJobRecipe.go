package databrewjob


type DatabrewJobRecipe struct {
	// Recipe name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_job#name DatabrewJob#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Recipe version.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_job#version DatabrewJob#version}
	Version *string `field:"optional" json:"version" yaml:"version"`
}

