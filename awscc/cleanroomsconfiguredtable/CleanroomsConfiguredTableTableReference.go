package cleanroomsconfiguredtable


type CleanroomsConfiguredTableTableReference struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cleanrooms_configured_table#athena CleanroomsConfiguredTable#athena}.
	Athena *CleanroomsConfiguredTableTableReferenceAthena `field:"optional" json:"athena" yaml:"athena"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cleanrooms_configured_table#glue CleanroomsConfiguredTable#glue}.
	Glue *CleanroomsConfiguredTableTableReferenceGlue `field:"optional" json:"glue" yaml:"glue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cleanrooms_configured_table#snowflake CleanroomsConfiguredTable#snowflake}.
	Snowflake *CleanroomsConfiguredTableTableReferenceSnowflake `field:"optional" json:"snowflake" yaml:"snowflake"`
}

