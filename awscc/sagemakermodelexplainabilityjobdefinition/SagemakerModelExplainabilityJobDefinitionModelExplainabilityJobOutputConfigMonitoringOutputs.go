package sagemakermodelexplainabilityjobdefinition


type SagemakerModelExplainabilityJobDefinitionModelExplainabilityJobOutputConfigMonitoringOutputs struct {
	// Information about where and how to store the results of a monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_explainability_job_definition#s3_output SagemakerModelExplainabilityJobDefinition#s3_output}
	S3Output *SagemakerModelExplainabilityJobDefinitionModelExplainabilityJobOutputConfigMonitoringOutputsS3Output `field:"required" json:"s3Output" yaml:"s3Output"`
}

