package bedrockagent


type BedrockAgentMemoryConfigurationSessionSummaryConfiguration struct {
	// Maximum number of Sessions to Summarize.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_agent#max_recent_sessions BedrockAgent#max_recent_sessions}
	MaxRecentSessions *float64 `field:"optional" json:"maxRecentSessions" yaml:"maxRecentSessions"`
}

