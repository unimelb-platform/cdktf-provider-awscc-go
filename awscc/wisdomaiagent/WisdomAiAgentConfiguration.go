package wisdomaiagent


type WisdomAiAgentConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_ai_agent#answer_recommendation_ai_agent_configuration WisdomAiAgent#answer_recommendation_ai_agent_configuration}.
	AnswerRecommendationAiAgentConfiguration *WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfiguration `field:"optional" json:"answerRecommendationAiAgentConfiguration" yaml:"answerRecommendationAiAgentConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_ai_agent#manual_search_ai_agent_configuration WisdomAiAgent#manual_search_ai_agent_configuration}.
	ManualSearchAiAgentConfiguration *WisdomAiAgentConfigurationManualSearchAiAgentConfiguration `field:"optional" json:"manualSearchAiAgentConfiguration" yaml:"manualSearchAiAgentConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_ai_agent#self_service_ai_agent_configuration WisdomAiAgent#self_service_ai_agent_configuration}.
	SelfServiceAiAgentConfiguration *WisdomAiAgentConfigurationSelfServiceAiAgentConfiguration `field:"optional" json:"selfServiceAiAgentConfiguration" yaml:"selfServiceAiAgentConfiguration"`
}

