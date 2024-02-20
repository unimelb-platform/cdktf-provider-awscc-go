package cleanroomsconfiguredtable


type CleanroomsConfiguredTableAnalysisRulesPolicyV1ListStruct struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_configured_table#join_columns CleanroomsConfiguredTable#join_columns}.
	JoinColumns *[]*string `field:"required" json:"joinColumns" yaml:"joinColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_configured_table#list_columns CleanroomsConfiguredTable#list_columns}.
	ListColumns *[]*string `field:"required" json:"listColumns" yaml:"listColumns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_configured_table#allowed_join_operators CleanroomsConfiguredTable#allowed_join_operators}.
	AllowedJoinOperators *[]*string `field:"optional" json:"allowedJoinOperators" yaml:"allowedJoinOperators"`
}

