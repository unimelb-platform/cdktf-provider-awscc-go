package datapipelinepipeline


type DatapipelinePipelineParameterObjectsAttributes struct {
	// The field identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datapipeline_pipeline#key DatapipelinePipeline#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The field value, expressed as a String.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datapipeline_pipeline#string_value DatapipelinePipeline#string_value}
	StringValue *string `field:"required" json:"stringValue" yaml:"stringValue"`
}

