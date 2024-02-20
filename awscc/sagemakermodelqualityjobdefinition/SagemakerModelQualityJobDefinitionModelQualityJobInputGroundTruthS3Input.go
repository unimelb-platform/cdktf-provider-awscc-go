package sagemakermodelqualityjobdefinition


type SagemakerModelQualityJobDefinitionModelQualityJobInputGroundTruthS3Input struct {
	// A URI that identifies the Amazon S3 storage location where Amazon SageMaker saves the results of a monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_quality_job_definition#s3_uri SagemakerModelQualityJobDefinition#s3_uri}
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
}

