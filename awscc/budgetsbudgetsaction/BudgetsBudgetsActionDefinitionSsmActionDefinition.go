package budgetsbudgetsaction


type BudgetsBudgetsActionDefinitionSsmActionDefinition struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/budgets_budgets_action#instance_ids BudgetsBudgetsAction#instance_ids}.
	InstanceIds *[]*string `field:"required" json:"instanceIds" yaml:"instanceIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/budgets_budgets_action#region BudgetsBudgetsAction#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/budgets_budgets_action#subtype BudgetsBudgetsAction#subtype}.
	Subtype *string `field:"required" json:"subtype" yaml:"subtype"`
}

