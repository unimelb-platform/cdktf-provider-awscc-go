package sagemakerdataqualityjobdefinition


type SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputDatasetFormat struct {
	// The CSV format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_data_quality_job_definition#csv SagemakerDataQualityJobDefinition#csv}
	Csv *SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputDatasetFormatCsv `field:"optional" json:"csv" yaml:"csv"`
	// The Json format.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_data_quality_job_definition#json SagemakerDataQualityJobDefinition#json}
	Json *SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputDatasetFormatJson `field:"optional" json:"json" yaml:"json"`
	// A flag indicate if the dataset format is Parquet.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_data_quality_job_definition#parquet SagemakerDataQualityJobDefinition#parquet}
	Parquet interface{} `field:"optional" json:"parquet" yaml:"parquet"`
}

