package bedrockagent


type BedrockAgentCustomOrchestration struct {
	// Types of executors for custom orchestration strategy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_agent#executor BedrockAgent#executor}
	Executor *BedrockAgentCustomOrchestrationExecutor `field:"optional" json:"executor" yaml:"executor"`
}

