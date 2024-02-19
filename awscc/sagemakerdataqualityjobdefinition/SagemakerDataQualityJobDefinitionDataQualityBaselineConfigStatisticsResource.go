package sagemakerdataqualityjobdefinition


type SagemakerDataQualityJobDefinitionDataQualityBaselineConfigStatisticsResource struct {
	// The Amazon S3 URI for the baseline statistics file in Amazon S3 that the current monitoring job should be validated against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_data_quality_job_definition#s3_uri SagemakerDataQualityJobDefinition#s3_uri}
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
}

