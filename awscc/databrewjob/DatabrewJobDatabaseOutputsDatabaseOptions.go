package databrewjob


type DatabrewJobDatabaseOutputsDatabaseOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_job#table_name DatabrewJob#table_name}.
	TableName *string `field:"required" json:"tableName" yaml:"tableName"`
	// S3 Output location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_job#temp_directory DatabrewJob#temp_directory}
	TempDirectory *DatabrewJobDatabaseOutputsDatabaseOptionsTempDirectory `field:"optional" json:"tempDirectory" yaml:"tempDirectory"`
}

