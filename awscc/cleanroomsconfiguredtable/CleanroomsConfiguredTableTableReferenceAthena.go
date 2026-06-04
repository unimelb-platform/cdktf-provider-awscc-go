package cleanroomsconfiguredtable


type CleanroomsConfiguredTableTableReferenceAthena struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cleanrooms_configured_table#database_name CleanroomsConfiguredTable#database_name}.
	DatabaseName *string `field:"optional" json:"databaseName" yaml:"databaseName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cleanrooms_configured_table#output_location CleanroomsConfiguredTable#output_location}.
	OutputLocation *string `field:"optional" json:"outputLocation" yaml:"outputLocation"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cleanrooms_configured_table#table_name CleanroomsConfiguredTable#table_name}.
	TableName *string `field:"optional" json:"tableName" yaml:"tableName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cleanrooms_configured_table#work_group CleanroomsConfiguredTable#work_group}.
	WorkGroup *string `field:"optional" json:"workGroup" yaml:"workGroup"`
}

