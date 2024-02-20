package sagemakerpipeline


type SagemakerPipelinePipelineDefinition struct {
	// A specification that defines the pipeline in JSON format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_pipeline#pipeline_definition_body SagemakerPipeline#pipeline_definition_body}
	PipelineDefinitionBody *string `field:"optional" json:"pipelineDefinitionBody" yaml:"pipelineDefinitionBody"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_pipeline#pipeline_definition_s3_location SagemakerPipeline#pipeline_definition_s3_location}.
	PipelineDefinitionS3Location *SagemakerPipelinePipelineDefinitionPipelineDefinitionS3Location `field:"optional" json:"pipelineDefinitionS3Location" yaml:"pipelineDefinitionS3Location"`
}

