package datazonedatasource


type DatazoneDataSourceConfigurationGlueRunConfiguration struct {
	// The relational filter configurations included in the configuration details of the AWS Glue data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datazone_data_source#relational_filter_configurations DatazoneDataSource#relational_filter_configurations}
	RelationalFilterConfigurations interface{} `field:"required" json:"relationalFilterConfigurations" yaml:"relationalFilterConfigurations"`
	// The data access role included in the configuration details of the AWS Glue data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datazone_data_source#data_access_role DatazoneDataSource#data_access_role}
	DataAccessRole *string `field:"optional" json:"dataAccessRole" yaml:"dataAccessRole"`
}

