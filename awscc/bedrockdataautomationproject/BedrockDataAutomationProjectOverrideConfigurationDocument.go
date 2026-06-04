package bedrockdataautomationproject


type BedrockDataAutomationProjectOverrideConfigurationDocument struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_data_automation_project#modality_processing BedrockDataAutomationProject#modality_processing}.
	ModalityProcessing *BedrockDataAutomationProjectOverrideConfigurationDocumentModalityProcessing `field:"optional" json:"modalityProcessing" yaml:"modalityProcessing"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_data_automation_project#splitter BedrockDataAutomationProject#splitter}.
	Splitter *BedrockDataAutomationProjectOverrideConfigurationDocumentSplitter `field:"optional" json:"splitter" yaml:"splitter"`
}

