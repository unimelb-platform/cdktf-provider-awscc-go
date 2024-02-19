package maciefindingsfilter


type MacieFindingsFilterFindingCriteria struct {
	// Map of filter criteria.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/macie_findings_filter#criterion MacieFindingsFilter#criterion}
	Criterion interface{} `field:"optional" json:"criterion" yaml:"criterion"`
}

