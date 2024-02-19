package datazonedatasource


type DatazoneDataSourceConfigurationRedshiftRunConfigurationRedshiftStorage struct {
	// The name of an Amazon Redshift cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datazone_data_source#redshift_cluster_source DatazoneDataSource#redshift_cluster_source}
	RedshiftClusterSource *DatazoneDataSourceConfigurationRedshiftRunConfigurationRedshiftStorageRedshiftClusterSource `field:"optional" json:"redshiftClusterSource" yaml:"redshiftClusterSource"`
	// The details of the Amazon Redshift Serverless workgroup storage.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datazone_data_source#redshift_serverless_source DatazoneDataSource#redshift_serverless_source}
	RedshiftServerlessSource *DatazoneDataSourceConfigurationRedshiftRunConfigurationRedshiftStorageRedshiftServerlessSource `field:"optional" json:"redshiftServerlessSource" yaml:"redshiftServerlessSource"`
}

