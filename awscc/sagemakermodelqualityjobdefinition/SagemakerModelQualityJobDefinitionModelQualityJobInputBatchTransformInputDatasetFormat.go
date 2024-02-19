package sagemakermodelqualityjobdefinition


type SagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputDatasetFormat struct {
	// The CSV format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_quality_job_definition#csv SagemakerModelQualityJobDefinition#csv}
	Csv *SagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputDatasetFormatCsv `field:"optional" json:"csv" yaml:"csv"`
	// The Json format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_quality_job_definition#json SagemakerModelQualityJobDefinition#json}
	Json *SagemakerModelQualityJobDefinitionModelQualityJobInputBatchTransformInputDatasetFormatJson `field:"optional" json:"json" yaml:"json"`
	// A flag indicating if the dataset format is Parquet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_quality_job_definition#parquet SagemakerModelQualityJobDefinition#parquet}
	Parquet interface{} `field:"optional" json:"parquet" yaml:"parquet"`
}

