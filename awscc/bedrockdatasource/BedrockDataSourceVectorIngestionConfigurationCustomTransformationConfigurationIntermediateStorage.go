package bedrockdatasource


type BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationIntermediateStorage struct {
	// An Amazon S3 location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_data_source#s3_location BedrockDataSource#s3_location}
	S3Location *BedrockDataSourceVectorIngestionConfigurationCustomTransformationConfigurationIntermediateStorageS3Location `field:"optional" json:"s3Location" yaml:"s3Location"`
}

