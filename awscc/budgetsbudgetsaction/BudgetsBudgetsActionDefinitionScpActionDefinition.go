package budgetsbudgetsaction


type BudgetsBudgetsActionDefinitionScpActionDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/budgets_budgets_action#policy_id BudgetsBudgetsAction#policy_id}.
	PolicyId *string `field:"required" json:"policyId" yaml:"policyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/budgets_budgets_action#target_ids BudgetsBudgetsAction#target_ids}.
	TargetIds *[]*string `field:"required" json:"targetIds" yaml:"targetIds"`
}

