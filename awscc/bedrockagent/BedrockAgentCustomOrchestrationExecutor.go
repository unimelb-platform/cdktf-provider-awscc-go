package bedrockagent


type BedrockAgentCustomOrchestrationExecutor struct {
	// ARN of a Lambda.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_agent#lambda BedrockAgent#lambda}
	Lambda *string `field:"optional" json:"lambda" yaml:"lambda"`
}

