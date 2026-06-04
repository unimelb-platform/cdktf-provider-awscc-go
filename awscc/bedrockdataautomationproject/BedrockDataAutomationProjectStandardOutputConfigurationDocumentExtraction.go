package bedrockdataautomationproject


type BedrockDataAutomationProjectStandardOutputConfigurationDocumentExtraction struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_data_automation_project#bounding_box BedrockDataAutomationProject#bounding_box}.
	BoundingBox *BedrockDataAutomationProjectStandardOutputConfigurationDocumentExtractionBoundingBox `field:"optional" json:"boundingBox" yaml:"boundingBox"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_data_automation_project#granularity BedrockDataAutomationProject#granularity}.
	Granularity *BedrockDataAutomationProjectStandardOutputConfigurationDocumentExtractionGranularity `field:"optional" json:"granularity" yaml:"granularity"`
}

