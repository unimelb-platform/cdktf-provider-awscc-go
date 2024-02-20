package sagemakermodelbiasjobdefinition


type SagemakerModelBiasJobDefinitionModelBiasJobInputBatchTransformInputDatasetFormat struct {
	// The CSV format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_bias_job_definition#csv SagemakerModelBiasJobDefinition#csv}
	Csv *SagemakerModelBiasJobDefinitionModelBiasJobInputBatchTransformInputDatasetFormatCsv `field:"optional" json:"csv" yaml:"csv"`
	// The Json format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_bias_job_definition#json SagemakerModelBiasJobDefinition#json}
	Json *SagemakerModelBiasJobDefinitionModelBiasJobInputBatchTransformInputDatasetFormatJson `field:"optional" json:"json" yaml:"json"`
	// A flag indicate if the dataset format is Parquet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_model_bias_job_definition#parquet SagemakerModelBiasJobDefinition#parquet}
	Parquet interface{} `field:"optional" json:"parquet" yaml:"parquet"`
}

