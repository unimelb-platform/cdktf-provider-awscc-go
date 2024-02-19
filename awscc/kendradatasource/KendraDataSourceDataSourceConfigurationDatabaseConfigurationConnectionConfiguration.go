package kendradatasource


type KendraDataSourceDataSourceConfigurationDatabaseConfigurationConnectionConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#database_host KendraDataSource#database_host}.
	DatabaseHost *string `field:"required" json:"databaseHost" yaml:"databaseHost"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#database_name KendraDataSource#database_name}.
	DatabaseName *string `field:"required" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#database_port KendraDataSource#database_port}.
	DatabasePort *float64 `field:"required" json:"databasePort" yaml:"databasePort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#secret_arn KendraDataSource#secret_arn}.
	SecretArn *string `field:"required" json:"secretArn" yaml:"secretArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kendra_data_source#table_name KendraDataSource#table_name}.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
}

