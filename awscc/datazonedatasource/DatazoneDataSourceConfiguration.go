package datazonedatasource


type DatazoneDataSourceConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datazone_data_source#glue_run_configuration DatazoneDataSource#glue_run_configuration}.
	GlueRunConfiguration *DatazoneDataSourceConfigurationGlueRunConfiguration `field:"optional" json:"glueRunConfiguration" yaml:"glueRunConfiguration"`
	// The configuration details of the Amazon Redshift data source.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datazone_data_source#redshift_run_configuration DatazoneDataSource#redshift_run_configuration}
	RedshiftRunConfiguration *DatazoneDataSourceConfigurationRedshiftRunConfiguration `field:"optional" json:"redshiftRunConfiguration" yaml:"redshiftRunConfiguration"`
}

