package bedrockknowledgebase


type BedrockKnowledgeBaseKnowledgeBaseConfiguration struct {
	// The type of a knowledge base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_knowledge_base#type BedrockKnowledgeBase#type}
	Type *string `field:"required" json:"type" yaml:"type"`
	// Configurations for a Kendra knowledge base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_knowledge_base#kendra_knowledge_base_configuration BedrockKnowledgeBase#kendra_knowledge_base_configuration}
	KendraKnowledgeBaseConfiguration *BedrockKnowledgeBaseKnowledgeBaseConfigurationKendraKnowledgeBaseConfiguration `field:"optional" json:"kendraKnowledgeBaseConfiguration" yaml:"kendraKnowledgeBaseConfiguration"`
	// Configurations for a SQL knowledge base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_knowledge_base#sql_knowledge_base_configuration BedrockKnowledgeBase#sql_knowledge_base_configuration}
	SqlKnowledgeBaseConfiguration *BedrockKnowledgeBaseKnowledgeBaseConfigurationSqlKnowledgeBaseConfiguration `field:"optional" json:"sqlKnowledgeBaseConfiguration" yaml:"sqlKnowledgeBaseConfiguration"`
	// Contains details about the model used to create vector embeddings for the knowledge base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_knowledge_base#vector_knowledge_base_configuration BedrockKnowledgeBase#vector_knowledge_base_configuration}
	VectorKnowledgeBaseConfiguration *BedrockKnowledgeBaseKnowledgeBaseConfigurationVectorKnowledgeBaseConfiguration `field:"optional" json:"vectorKnowledgeBaseConfiguration" yaml:"vectorKnowledgeBaseConfiguration"`
}

