package bedrockprompt


type BedrockPromptVariantsGenAiResource struct {
	// Target Agent to invoke with Prompt.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_prompt#agent BedrockPrompt#agent}
	Agent *BedrockPromptVariantsGenAiResourceAgent `field:"optional" json:"agent" yaml:"agent"`
}

