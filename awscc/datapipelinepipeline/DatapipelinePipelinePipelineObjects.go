package datapipelinepipeline


type DatapipelinePipelinePipelineObjects struct {
	// Key-value pairs that define the properties of the object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datapipeline_pipeline#fields DatapipelinePipeline#fields}
	Fields interface{} `field:"required" json:"fields" yaml:"fields"`
	// The ID of the object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datapipeline_pipeline#id DatapipelinePipeline#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// The name of the object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datapipeline_pipeline#name DatapipelinePipeline#name}
	Name *string `field:"required" json:"name" yaml:"name"`
}

