package wisdomknowledgebase


type WisdomKnowledgeBaseVectorIngestionConfigurationChunkingConfigurationFixedSizeChunkingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_knowledge_base#max_tokens WisdomKnowledgeBase#max_tokens}.
	MaxTokens *float64 `field:"optional" json:"maxTokens" yaml:"maxTokens"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_knowledge_base#overlap_percentage WisdomKnowledgeBase#overlap_percentage}.
	OverlapPercentage *float64 `field:"optional" json:"overlapPercentage" yaml:"overlapPercentage"`
}

