package sagemakermodelexplainabilityjobdefinition


type SagemakerModelExplainabilityJobDefinitionStoppingCondition struct {
	// The maximum runtime allowed in seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_explainability_job_definition#max_runtime_in_seconds SagemakerModelExplainabilityJobDefinition#max_runtime_in_seconds}
	MaxRuntimeInSeconds *float64 `field:"required" json:"maxRuntimeInSeconds" yaml:"maxRuntimeInSeconds"`
}

