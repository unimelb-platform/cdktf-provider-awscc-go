package bedrockagent


type BedrockAgentActionGroupsFunctionSchema struct {
	// List of Function definitions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_agent#functions BedrockAgent#functions}
	Functions interface{} `field:"optional" json:"functions" yaml:"functions"`
}

