package bedrockknowledgebase


type BedrockKnowledgeBaseKnowledgeBaseConfigurationSqlKnowledgeBaseConfiguration struct {
	// Configurations for a Redshift knowledge base.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_knowledge_base#redshift_configuration BedrockKnowledgeBase#redshift_configuration}
	RedshiftConfiguration *BedrockKnowledgeBaseKnowledgeBaseConfigurationSqlKnowledgeBaseConfigurationRedshiftConfiguration `field:"optional" json:"redshiftConfiguration" yaml:"redshiftConfiguration"`
	// SQL query engine type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_knowledge_base#type BedrockKnowledgeBase#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

