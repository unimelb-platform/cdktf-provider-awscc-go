package cleanroomsconfiguredtable


type CleanroomsConfiguredTableTableReferenceGlue struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cleanrooms_configured_table#database_name CleanroomsConfiguredTable#database_name}.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cleanrooms_configured_table#table_name CleanroomsConfiguredTable#table_name}.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
}

