package datazonedatasource


type DatazoneDataSourceConfigurationRedshiftRunConfigurationRedshiftStorageRedshiftClusterSource struct {
	// The name of an Amazon Redshift cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_data_source#cluster_name DatazoneDataSource#cluster_name}
	ClusterName *string `field:"optional" json:"clusterName" yaml:"clusterName"`
}

