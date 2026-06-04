package cleanroomsprivacybudgettemplate


type CleanroomsPrivacyBudgetTemplateParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cleanrooms_privacy_budget_template#epsilon CleanroomsPrivacyBudgetTemplate#epsilon}.
	Epsilon *float64 `field:"required" json:"epsilon" yaml:"epsilon"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cleanrooms_privacy_budget_template#users_noise_per_query CleanroomsPrivacyBudgetTemplate#users_noise_per_query}.
	UsersNoisePerQuery *float64 `field:"required" json:"usersNoisePerQuery" yaml:"usersNoisePerQuery"`
}

