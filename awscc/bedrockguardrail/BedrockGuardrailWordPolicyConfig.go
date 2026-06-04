package bedrockguardrail


type BedrockGuardrailWordPolicyConfig struct {
	// A config for the list of managed words.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_guardrail#managed_word_lists_config BedrockGuardrail#managed_word_lists_config}
	ManagedWordListsConfig interface{} `field:"optional" json:"managedWordListsConfig" yaml:"managedWordListsConfig"`
	// List of custom word configs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_guardrail#words_config BedrockGuardrail#words_config}
	WordsConfig interface{} `field:"optional" json:"wordsConfig" yaml:"wordsConfig"`
}

