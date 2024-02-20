package kendradatasource


type KendraDataSourceDataSourceConfigurationSalesforceConfigurationStandardObjectConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#document_data_field_name KendraDataSource#document_data_field_name}.
	DocumentDataFieldName *string `field:"required" json:"documentDataFieldName" yaml:"documentDataFieldName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#name KendraDataSource#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#document_title_field_name KendraDataSource#document_title_field_name}.
	DocumentTitleFieldName *string `field:"optional" json:"documentTitleFieldName" yaml:"documentTitleFieldName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#field_mappings KendraDataSource#field_mappings}.
	FieldMappings interface{} `field:"optional" json:"fieldMappings" yaml:"fieldMappings"`
}

