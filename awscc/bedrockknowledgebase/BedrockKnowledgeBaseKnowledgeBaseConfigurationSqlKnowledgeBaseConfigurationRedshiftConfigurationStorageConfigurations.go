package bedrockknowledgebase


type BedrockKnowledgeBaseKnowledgeBaseConfigurationSqlKnowledgeBaseConfigurationRedshiftConfigurationStorageConfigurations struct {
	// Configurations for Redshift query engine AWS Data Catalog backed storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_knowledge_base#aws_data_catalog_configuration BedrockKnowledgeBase#aws_data_catalog_configuration}
	AwsDataCatalogConfiguration *BedrockKnowledgeBaseKnowledgeBaseConfigurationSqlKnowledgeBaseConfigurationRedshiftConfigurationStorageConfigurationsAwsDataCatalogConfiguration `field:"optional" json:"awsDataCatalogConfiguration" yaml:"awsDataCatalogConfiguration"`
	// Configurations for Redshift query engine Redshift backed storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_knowledge_base#redshift_configuration BedrockKnowledgeBase#redshift_configuration}
	RedshiftConfiguration *BedrockKnowledgeBaseKnowledgeBaseConfigurationSqlKnowledgeBaseConfigurationRedshiftConfigurationStorageConfigurationsRedshiftConfiguration `field:"optional" json:"redshiftConfiguration" yaml:"redshiftConfiguration"`
	// Redshift query engine storage type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/bedrock_knowledge_base#type BedrockKnowledgeBase#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

