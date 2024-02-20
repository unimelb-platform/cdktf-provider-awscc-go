package sagemakermodelexplainabilityjobdefinition


type SagemakerModelExplainabilityJobDefinitionModelExplainabilityBaselineConfigConstraintsResource struct {
	// The Amazon S3 URI for baseline constraint file in Amazon S3 that the current monitoring job should validated against.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_explainability_job_definition#s3_uri SagemakerModelExplainabilityJobDefinition#s3_uri}
	S3Uri *string `field:"optional" json:"s3Uri" yaml:"s3Uri"`
}

