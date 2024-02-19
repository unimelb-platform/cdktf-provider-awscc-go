package kendradatasource


type KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#included_states KendraDataSource#included_states}.
	IncludedStates *[]*string `field:"required" json:"includedStates" yaml:"includedStates"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#custom_knowledge_article_type_configurations KendraDataSource#custom_knowledge_article_type_configurations}.
	CustomKnowledgeArticleTypeConfigurations interface{} `field:"optional" json:"customKnowledgeArticleTypeConfigurations" yaml:"customKnowledgeArticleTypeConfigurations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#standard_knowledge_article_type_configuration KendraDataSource#standard_knowledge_article_type_configuration}.
	StandardKnowledgeArticleTypeConfiguration *KendraDataSourceDataSourceConfigurationSalesforceConfigurationKnowledgeArticleConfigurationStandardKnowledgeArticleTypeConfiguration `field:"optional" json:"standardKnowledgeArticleTypeConfiguration" yaml:"standardKnowledgeArticleTypeConfiguration"`
}

