package bedrockprompt


type BedrockPromptVariantsGenAiResourceAgent struct {
	// Arn representation of the Agent Alias.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_prompt#agent_identifier BedrockPrompt#agent_identifier}
	AgentIdentifier *string `field:"optional" json:"agentIdentifier" yaml:"agentIdentifier"`
}

