package datapipelinepipeline


type DatapipelinePipelinePipelineTags struct {
	// The key name of a tag.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datapipeline_pipeline#key DatapipelinePipeline#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value to associate with the key name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datapipeline_pipeline#value DatapipelinePipeline#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

