package sagemakermodelqualityjobdefinition


type SagemakerModelQualityJobDefinitionModelQualityJobOutputConfigMonitoringOutputs struct {
	// Information about where and how to store the results of a monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_quality_job_definition#s3_output SagemakerModelQualityJobDefinition#s3_output}
	S3Output *SagemakerModelQualityJobDefinitionModelQualityJobOutputConfigMonitoringOutputsS3Output `field:"required" json:"s3Output" yaml:"s3Output"`
}

