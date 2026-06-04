package bedrockknowledgebase


type BedrockKnowledgeBaseKnowledgeBaseConfigurationVectorKnowledgeBaseConfigurationEmbeddingModelConfigurationBedrockEmbeddingModelConfiguration struct {
	// The dimensions details for the vector configuration used on the Bedrock embeddings model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_knowledge_base#dimensions BedrockKnowledgeBase#dimensions}
	Dimensions *float64 `field:"optional" json:"dimensions" yaml:"dimensions"`
	// The data type for the vectors when using a model to convert text into vector embeddings.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_knowledge_base#embedding_data_type BedrockKnowledgeBase#embedding_data_type}
	EmbeddingDataType *string `field:"optional" json:"embeddingDataType" yaml:"embeddingDataType"`
}

