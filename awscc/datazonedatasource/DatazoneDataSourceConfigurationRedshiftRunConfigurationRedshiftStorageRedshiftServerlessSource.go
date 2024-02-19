package datazonedatasource


type DatazoneDataSourceConfigurationRedshiftRunConfigurationRedshiftStorageRedshiftServerlessSource struct {
	// The name of the Amazon Redshift Serverless workgroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datazone_data_source#workgroup_name DatazoneDataSource#workgroup_name}
	WorkgroupName *string `field:"required" json:"workgroupName" yaml:"workgroupName"`
}

