package kendradatasource


type KendraDataSourceDataSourceConfigurationDatabaseConfigurationColumnConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#change_detecting_columns KendraDataSource#change_detecting_columns}.
	ChangeDetectingColumns *[]*string `field:"required" json:"changeDetectingColumns" yaml:"changeDetectingColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#document_data_column_name KendraDataSource#document_data_column_name}.
	DocumentDataColumnName *string `field:"required" json:"documentDataColumnName" yaml:"documentDataColumnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#document_id_column_name KendraDataSource#document_id_column_name}.
	DocumentIdColumnName *string `field:"required" json:"documentIdColumnName" yaml:"documentIdColumnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#document_title_column_name KendraDataSource#document_title_column_name}.
	DocumentTitleColumnName *string `field:"optional" json:"documentTitleColumnName" yaml:"documentTitleColumnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#field_mappings KendraDataSource#field_mappings}.
	FieldMappings interface{} `field:"optional" json:"fieldMappings" yaml:"fieldMappings"`
}

