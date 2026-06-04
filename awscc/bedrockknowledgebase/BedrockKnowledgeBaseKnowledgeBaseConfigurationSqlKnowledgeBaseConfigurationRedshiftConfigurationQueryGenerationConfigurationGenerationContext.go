package bedrockknowledgebase


type BedrockKnowledgeBaseKnowledgeBaseConfigurationSqlKnowledgeBaseConfigurationRedshiftConfigurationQueryGenerationConfigurationGenerationContext struct {
	// List of example queries and results.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_knowledge_base#curated_queries BedrockKnowledgeBase#curated_queries}
	CuratedQueries interface{} `field:"optional" json:"curatedQueries" yaml:"curatedQueries"`
	// List of tables used for Redshift query generation context.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_knowledge_base#tables BedrockKnowledgeBase#tables}
	Tables interface{} `field:"optional" json:"tables" yaml:"tables"`
}

