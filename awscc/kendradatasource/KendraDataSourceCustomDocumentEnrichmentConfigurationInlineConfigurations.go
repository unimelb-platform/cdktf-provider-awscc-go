package kendradatasource


type KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#condition KendraDataSource#condition}.
	Condition *KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsCondition `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#document_content_deletion KendraDataSource#document_content_deletion}.
	DocumentContentDeletion interface{} `field:"optional" json:"documentContentDeletion" yaml:"documentContentDeletion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#target KendraDataSource#target}.
	Target *KendraDataSourceCustomDocumentEnrichmentConfigurationInlineConfigurationsTarget `field:"optional" json:"target" yaml:"target"`
}

