package bedrockdatasource


type BedrockDataSourceVectorIngestionConfigurationParsingConfigurationBedrockDataAutomationConfiguration struct {
	// Determine how will parsed content be stored.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_data_source#parsing_modality BedrockDataSource#parsing_modality}
	ParsingModality *string `field:"optional" json:"parsingModality" yaml:"parsingModality"`
}

