package cleanroomsconfiguredtable


type CleanroomsConfiguredTableAnalysisRulesPolicyV1Aggregation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_configured_table#aggregate_columns CleanroomsConfiguredTable#aggregate_columns}.
	AggregateColumns interface{} `field:"required" json:"aggregateColumns" yaml:"aggregateColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_configured_table#dimension_columns CleanroomsConfiguredTable#dimension_columns}.
	DimensionColumns *[]*string `field:"required" json:"dimensionColumns" yaml:"dimensionColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_configured_table#join_columns CleanroomsConfiguredTable#join_columns}.
	JoinColumns *[]*string `field:"required" json:"joinColumns" yaml:"joinColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_configured_table#output_constraints CleanroomsConfiguredTable#output_constraints}.
	OutputConstraints interface{} `field:"required" json:"outputConstraints" yaml:"outputConstraints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_configured_table#scalar_functions CleanroomsConfiguredTable#scalar_functions}.
	ScalarFunctions *[]*string `field:"required" json:"scalarFunctions" yaml:"scalarFunctions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_configured_table#allowed_join_operators CleanroomsConfiguredTable#allowed_join_operators}.
	AllowedJoinOperators *[]*string `field:"optional" json:"allowedJoinOperators" yaml:"allowedJoinOperators"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_configured_table#join_required CleanroomsConfiguredTable#join_required}.
	JoinRequired *string `field:"optional" json:"joinRequired" yaml:"joinRequired"`
}

