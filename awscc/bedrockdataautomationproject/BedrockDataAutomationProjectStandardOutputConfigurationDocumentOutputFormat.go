package bedrockdataautomationproject


type BedrockDataAutomationProjectStandardOutputConfigurationDocumentOutputFormat struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_data_automation_project#additional_file_format BedrockDataAutomationProject#additional_file_format}.
	AdditionalFileFormat *BedrockDataAutomationProjectStandardOutputConfigurationDocumentOutputFormatAdditionalFileFormat `field:"optional" json:"additionalFileFormat" yaml:"additionalFileFormat"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_data_automation_project#text_format BedrockDataAutomationProject#text_format}.
	TextFormat *BedrockDataAutomationProjectStandardOutputConfigurationDocumentOutputFormatTextFormat `field:"optional" json:"textFormat" yaml:"textFormat"`
}

