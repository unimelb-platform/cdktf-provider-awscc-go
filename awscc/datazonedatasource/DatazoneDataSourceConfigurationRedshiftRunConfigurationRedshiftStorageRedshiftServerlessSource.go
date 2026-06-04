package datazonedatasource


type DatazoneDataSourceConfigurationRedshiftRunConfigurationRedshiftStorageRedshiftServerlessSource struct {
	// The name of the Amazon Redshift Serverless workgroup.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datazone_data_source#workgroup_name DatazoneDataSource#workgroup_name}
	WorkgroupName *string `field:"optional" json:"workgroupName" yaml:"workgroupName"`
}

