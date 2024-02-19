package cleanroomsconfiguredtable


type CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationAggregateColumns struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_configured_table#column_names CleanroomsConfiguredTable#column_names}.
	ColumnNames *[]*string `field:"required" json:"columnNames" yaml:"columnNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_configured_table#function CleanroomsConfiguredTable#function}.
	Function *string `field:"required" json:"function" yaml:"function"`
}

