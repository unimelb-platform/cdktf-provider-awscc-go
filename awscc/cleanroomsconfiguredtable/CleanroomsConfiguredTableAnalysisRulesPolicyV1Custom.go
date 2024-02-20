package cleanroomsconfiguredtable


type CleanroomsConfiguredTableAnalysisRulesPolicyV1Custom struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_configured_table#allowed_analyses CleanroomsConfiguredTable#allowed_analyses}.
	AllowedAnalyses *[]*string `field:"required" json:"allowedAnalyses" yaml:"allowedAnalyses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/cleanrooms_configured_table#allowed_analysis_providers CleanroomsConfiguredTable#allowed_analysis_providers}.
	AllowedAnalysisProviders *[]*string `field:"optional" json:"allowedAnalysisProviders" yaml:"allowedAnalysisProviders"`
}

