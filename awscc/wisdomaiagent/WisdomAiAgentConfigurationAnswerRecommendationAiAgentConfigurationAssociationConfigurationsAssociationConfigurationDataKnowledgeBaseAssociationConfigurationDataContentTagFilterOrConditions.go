package wisdomaiagent


type WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataKnowledgeBaseAssociationConfigurationDataContentTagFilterOrConditions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_ai_agent#and_conditions WisdomAiAgent#and_conditions}.
	AndConditions interface{} `field:"optional" json:"andConditions" yaml:"andConditions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_ai_agent#tag_condition WisdomAiAgent#tag_condition}.
	TagCondition *WisdomAiAgentConfigurationAnswerRecommendationAiAgentConfigurationAssociationConfigurationsAssociationConfigurationDataKnowledgeBaseAssociationConfigurationDataContentTagFilterOrConditionsTagCondition `field:"optional" json:"tagCondition" yaml:"tagCondition"`
}

