package datapipelinepipeline


type DatapipelinePipelineParameterObjects struct {
	// The attributes of the parameter object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datapipeline_pipeline#attributes DatapipelinePipeline#attributes}
	Attributes interface{} `field:"required" json:"attributes" yaml:"attributes"`
	// The ID of the parameter object.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datapipeline_pipeline#id DatapipelinePipeline#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
}

