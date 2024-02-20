package maciefindingsfilter


type MacieFindingsFilterFindingCriteriaCriterion struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/macie_findings_filter#eq MacieFindingsFilter#eq}.
	Eq *[]*string `field:"optional" json:"eq" yaml:"eq"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/macie_findings_filter#gt MacieFindingsFilter#gt}.
	Gt *float64 `field:"optional" json:"gt" yaml:"gt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/macie_findings_filter#gte MacieFindingsFilter#gte}.
	Gte *float64 `field:"optional" json:"gte" yaml:"gte"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/macie_findings_filter#lt MacieFindingsFilter#lt}.
	Lt *float64 `field:"optional" json:"lt" yaml:"lt"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/macie_findings_filter#lte MacieFindingsFilter#lte}.
	Lte *float64 `field:"optional" json:"lte" yaml:"lte"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/macie_findings_filter#neq MacieFindingsFilter#neq}.
	Neq *[]*string `field:"optional" json:"neq" yaml:"neq"`
}

