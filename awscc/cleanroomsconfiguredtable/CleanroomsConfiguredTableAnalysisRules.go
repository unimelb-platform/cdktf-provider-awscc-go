package cleanroomsconfiguredtable


type CleanroomsConfiguredTableAnalysisRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_configured_table#policy CleanroomsConfiguredTable#policy}.
	Policy *CleanroomsConfiguredTableAnalysisRulesPolicy `field:"required" json:"policy" yaml:"policy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_configured_table#type CleanroomsConfiguredTable#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

