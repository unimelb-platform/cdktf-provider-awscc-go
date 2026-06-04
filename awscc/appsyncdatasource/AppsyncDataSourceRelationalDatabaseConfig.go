package appsyncdatasource


type AppsyncDataSourceRelationalDatabaseConfig struct {
	// Information about the Amazon RDS resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/appsync_data_source#rds_http_endpoint_config AppsyncDataSource#rds_http_endpoint_config}
	RdsHttpEndpointConfig *AppsyncDataSourceRelationalDatabaseConfigRdsHttpEndpointConfig `field:"optional" json:"rdsHttpEndpointConfig" yaml:"rdsHttpEndpointConfig"`
	// The type of relational data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/appsync_data_source#relational_database_source_type AppsyncDataSource#relational_database_source_type}
	RelationalDatabaseSourceType *string `field:"optional" json:"relationalDatabaseSourceType" yaml:"relationalDatabaseSourceType"`
}

