package bedrockprompt


type BedrockPromptVariantsMetadata struct {
	// The key of a metadata tag for a prompt variant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_prompt#key BedrockPrompt#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The value of a metadata tag for a prompt variant.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_prompt#value BedrockPrompt#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

