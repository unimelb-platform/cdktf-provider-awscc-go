package sagemakermodelexplainabilityjobdefinition


type SagemakerModelExplainabilityJobDefinitionModelExplainabilityAppSpecification struct {
	// The S3 URI to an analysis configuration file.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_explainability_job_definition#config_uri SagemakerModelExplainabilityJobDefinition#config_uri}
	ConfigUri *string `field:"required" json:"configUri" yaml:"configUri"`
	// The container image to be run by the monitoring job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_explainability_job_definition#image_uri SagemakerModelExplainabilityJobDefinition#image_uri}
	ImageUri *string `field:"required" json:"imageUri" yaml:"imageUri"`
	// Sets the environment variables in the Docker container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_explainability_job_definition#environment SagemakerModelExplainabilityJobDefinition#environment}
	Environment *map[string]*string `field:"optional" json:"environment" yaml:"environment"`
}

