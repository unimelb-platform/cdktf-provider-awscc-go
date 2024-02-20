package kendradatasource


type KendraDataSourceDataSourceConfigurationServiceNowConfigurationKnowledgeArticleConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#document_data_field_name KendraDataSource#document_data_field_name}.
	DocumentDataFieldName *string `field:"required" json:"documentDataFieldName" yaml:"documentDataFieldName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#crawl_attachments KendraDataSource#crawl_attachments}.
	CrawlAttachments interface{} `field:"optional" json:"crawlAttachments" yaml:"crawlAttachments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#document_title_field_name KendraDataSource#document_title_field_name}.
	DocumentTitleFieldName *string `field:"optional" json:"documentTitleFieldName" yaml:"documentTitleFieldName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#exclude_attachment_file_patterns KendraDataSource#exclude_attachment_file_patterns}.
	ExcludeAttachmentFilePatterns *[]*string `field:"optional" json:"excludeAttachmentFilePatterns" yaml:"excludeAttachmentFilePatterns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#field_mappings KendraDataSource#field_mappings}.
	FieldMappings interface{} `field:"optional" json:"fieldMappings" yaml:"fieldMappings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#filter_query KendraDataSource#filter_query}.
	FilterQuery *string `field:"optional" json:"filterQuery" yaml:"filterQuery"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#include_attachment_file_patterns KendraDataSource#include_attachment_file_patterns}.
	IncludeAttachmentFilePatterns *[]*string `field:"optional" json:"includeAttachmentFilePatterns" yaml:"includeAttachmentFilePatterns"`
}

