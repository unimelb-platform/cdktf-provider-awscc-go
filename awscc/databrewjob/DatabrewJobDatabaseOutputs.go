package databrewjob


type DatabrewJobDatabaseOutputs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_job#database_options DatabrewJob#database_options}.
	DatabaseOptions *DatabrewJobDatabaseOutputsDatabaseOptions `field:"required" json:"databaseOptions" yaml:"databaseOptions"`
	// Glue connection name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_job#glue_connection_name DatabrewJob#glue_connection_name}
	GlueConnectionName *string `field:"required" json:"glueConnectionName" yaml:"glueConnectionName"`
	// Database table name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/databrew_job#database_output_mode DatabrewJob#database_output_mode}
	DatabaseOutputMode *string `field:"optional" json:"databaseOutputMode" yaml:"databaseOutputMode"`
}

