package kendradatasource


type KendraDataSourceDataSourceConfigurationServiceNowConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#host_url KendraDataSource#host_url}.
	HostUrl *string `field:"required" json:"hostUrl" yaml:"hostUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#secret_arn KendraDataSource#secret_arn}.
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#service_now_build_version KendraDataSource#service_now_build_version}.
	ServiceNowBuildVersion *string `field:"required" json:"serviceNowBuildVersion" yaml:"serviceNowBuildVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#authentication_type KendraDataSource#authentication_type}.
	AuthenticationType *string `field:"optional" json:"authenticationType" yaml:"authenticationType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#knowledge_article_configuration KendraDataSource#knowledge_article_configuration}.
	KnowledgeArticleConfiguration *KendraDataSourceDataSourceConfigurationServiceNowConfigurationKnowledgeArticleConfiguration `field:"optional" json:"knowledgeArticleConfiguration" yaml:"knowledgeArticleConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#service_catalog_configuration KendraDataSource#service_catalog_configuration}.
	ServiceCatalogConfiguration *KendraDataSourceDataSourceConfigurationServiceNowConfigurationServiceCatalogConfiguration `field:"optional" json:"serviceCatalogConfiguration" yaml:"serviceCatalogConfiguration"`
}

