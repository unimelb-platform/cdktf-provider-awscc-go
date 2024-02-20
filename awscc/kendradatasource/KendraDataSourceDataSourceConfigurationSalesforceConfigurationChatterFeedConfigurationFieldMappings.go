package kendradatasource


type KendraDataSourceDataSourceConfigurationSalesforceConfigurationChatterFeedConfigurationFieldMappings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#data_source_field_name KendraDataSource#data_source_field_name}.
	DataSourceFieldName *string `field:"required" json:"dataSourceFieldName" yaml:"dataSourceFieldName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#index_field_name KendraDataSource#index_field_name}.
	IndexFieldName *string `field:"required" json:"indexFieldName" yaml:"indexFieldName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#date_field_format KendraDataSource#date_field_format}.
	DateFieldFormat *string `field:"optional" json:"dateFieldFormat" yaml:"dateFieldFormat"`
}

