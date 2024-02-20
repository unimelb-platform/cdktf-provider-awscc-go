package budgetsbudgetsaction


type BudgetsBudgetsActionDefinitionIamActionDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/budgets_budgets_action#policy_arn BudgetsBudgetsAction#policy_arn}.
	PolicyArn *string `field:"required" json:"policyArn" yaml:"policyArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/budgets_budgets_action#groups BudgetsBudgetsAction#groups}.
	Groups *[]*string `field:"optional" json:"groups" yaml:"groups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/budgets_budgets_action#roles BudgetsBudgetsAction#roles}.
	Roles *[]*string `field:"optional" json:"roles" yaml:"roles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/budgets_budgets_action#users BudgetsBudgetsAction#users}.
	Users *[]*string `field:"optional" json:"users" yaml:"users"`
}

