package cleanroomsconfiguredtable


type CleanroomsConfiguredTableAnalysisRulesPolicyV1AggregationOutputConstraints struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_configured_table#column_name CleanroomsConfiguredTable#column_name}.
	ColumnName *string `field:"required" json:"columnName" yaml:"columnName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_configured_table#minimum CleanroomsConfiguredTable#minimum}.
	Minimum *float64 `field:"required" json:"minimum" yaml:"minimum"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_configured_table#type CleanroomsConfiguredTable#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

