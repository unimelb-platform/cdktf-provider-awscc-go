package wisdomknowledgebase


type WisdomKnowledgeBaseVectorIngestionConfigurationChunkingConfigurationHierarchicalChunkingConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_knowledge_base#level_configurations WisdomKnowledgeBase#level_configurations}.
	LevelConfigurations interface{} `field:"optional" json:"levelConfigurations" yaml:"levelConfigurations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/wisdom_knowledge_base#overlap_tokens WisdomKnowledgeBase#overlap_tokens}.
	OverlapTokens *float64 `field:"optional" json:"overlapTokens" yaml:"overlapTokens"`
}

