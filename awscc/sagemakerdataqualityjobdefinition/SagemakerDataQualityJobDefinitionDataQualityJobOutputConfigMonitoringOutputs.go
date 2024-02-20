package sagemakerdataqualityjobdefinition


type SagemakerDataQualityJobDefinitionDataQualityJobOutputConfigMonitoringOutputs struct {
	// Information about where and how to store the results of a monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_data_quality_job_definition#s3_output SagemakerDataQualityJobDefinition#s3_output}
	S3Output *SagemakerDataQualityJobDefinitionDataQualityJobOutputConfigMonitoringOutputsS3Output `field:"required" json:"s3Output" yaml:"s3Output"`
}

