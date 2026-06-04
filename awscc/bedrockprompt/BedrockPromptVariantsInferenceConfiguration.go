package bedrockprompt


type BedrockPromptVariantsInferenceConfiguration struct {
	// Prompt model inference configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_prompt#text BedrockPrompt#text}
	Text *BedrockPromptVariantsInferenceConfigurationText `field:"optional" json:"text" yaml:"text"`
}

