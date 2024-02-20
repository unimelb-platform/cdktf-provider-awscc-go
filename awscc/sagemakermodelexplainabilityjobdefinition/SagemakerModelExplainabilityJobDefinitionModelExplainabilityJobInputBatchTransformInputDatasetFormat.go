package sagemakermodelexplainabilityjobdefinition


type SagemakerModelExplainabilityJobDefinitionModelExplainabilityJobInputBatchTransformInputDatasetFormat struct {
	// The CSV format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_explainability_job_definition#csv SagemakerModelExplainabilityJobDefinition#csv}
	Csv *SagemakerModelExplainabilityJobDefinitionModelExplainabilityJobInputBatchTransformInputDatasetFormatCsv `field:"optional" json:"csv" yaml:"csv"`
	// The Json format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_explainability_job_definition#json SagemakerModelExplainabilityJobDefinition#json}
	Json *SagemakerModelExplainabilityJobDefinitionModelExplainabilityJobInputBatchTransformInputDatasetFormatJson `field:"optional" json:"json" yaml:"json"`
	// A flag indicating if the dataset format is Parquet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_explainability_job_definition#parquet SagemakerModelExplainabilityJobDefinition#parquet}
	Parquet interface{} `field:"optional" json:"parquet" yaml:"parquet"`
}

