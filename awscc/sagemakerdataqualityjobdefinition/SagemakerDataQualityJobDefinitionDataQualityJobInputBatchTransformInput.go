package sagemakerdataqualityjobdefinition


type SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInput struct {
	// A URI that identifies the Amazon S3 storage location where Batch Transform Job captures data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_data_quality_job_definition#data_captured_destination_s3_uri SagemakerDataQualityJobDefinition#data_captured_destination_s3_uri}
	DataCapturedDestinationS3Uri *string `field:"required" json:"dataCapturedDestinationS3Uri" yaml:"dataCapturedDestinationS3Uri"`
	// The dataset format of the data to monitor.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_data_quality_job_definition#dataset_format SagemakerDataQualityJobDefinition#dataset_format}
	DatasetFormat *SagemakerDataQualityJobDefinitionDataQualityJobInputBatchTransformInputDatasetFormat `field:"required" json:"datasetFormat" yaml:"datasetFormat"`
	// Path to the filesystem where the endpoint data is available to the container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_data_quality_job_definition#local_path SagemakerDataQualityJobDefinition#local_path}
	LocalPath *string `field:"required" json:"localPath" yaml:"localPath"`
	// Indexes or names of the features to be excluded from analysis.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_data_quality_job_definition#exclude_features_attribute SagemakerDataQualityJobDefinition#exclude_features_attribute}
	ExcludeFeaturesAttribute *string `field:"optional" json:"excludeFeaturesAttribute" yaml:"excludeFeaturesAttribute"`
	// Whether input data distributed in Amazon S3 is fully replicated or sharded by an S3 key. Defauts to FullyReplicated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_data_quality_job_definition#s3_data_distribution_type SagemakerDataQualityJobDefinition#s3_data_distribution_type}
	S3DataDistributionType *string `field:"optional" json:"s3DataDistributionType" yaml:"s3DataDistributionType"`
	// Whether the Pipe or File is used as the input mode for transfering data for the monitoring job.
	//
	// Pipe mode is recommended for large datasets. File mode is useful for small files that fit in memory. Defaults to File.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_data_quality_job_definition#s3_input_mode SagemakerDataQualityJobDefinition#s3_input_mode}
	S3InputMode *string `field:"optional" json:"s3InputMode" yaml:"s3InputMode"`
}

