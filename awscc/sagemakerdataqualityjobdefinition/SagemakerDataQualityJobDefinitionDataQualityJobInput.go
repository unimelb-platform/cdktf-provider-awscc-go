package sagemakerdataqualityjobdefinition


type SagemakerDataQualityJobDefinitionDataQualityJobInput struct {
	// The batch transform input for a monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_data_quality_job_definition#batch_transform_input SagemakerDataQualityJobDefinition#batch_transform_input}
	BatchTransformInput *SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInput `field:"optional" json:"batchTransformInput" yaml:"batchTransformInput"`
	// The endpoint for a monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_data_quality_job_definition#endpoint_input SagemakerDataQualityJobDefinition#endpoint_input}
	EndpointInput *SagemakerDataQualityJobDefinitionDataQualityJobInputEndpointInput `field:"optional" json:"endpointInput" yaml:"endpointInput"`
}

