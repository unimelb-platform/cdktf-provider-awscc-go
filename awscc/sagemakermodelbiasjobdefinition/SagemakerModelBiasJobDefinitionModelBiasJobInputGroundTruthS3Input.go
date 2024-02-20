package sagemakermodelbiasjobdefinition


type SagemakerModelBiasJobDefinitionModelBiasJobInputGroundTruthS3Input struct {
	// A URI that identifies the Amazon S3 storage location where Amazon SageMaker saves the results of a monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_bias_job_definition#s3_uri SagemakerModelBiasJobDefinition#s3_uri}
	S3Uri *string `field:"required" json:"s3Uri" yaml:"s3Uri"`
}

