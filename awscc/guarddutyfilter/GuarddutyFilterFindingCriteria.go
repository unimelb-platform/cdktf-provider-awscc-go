package guarddutyfilter


type GuarddutyFilterFindingCriteria struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/guardduty_filter#criterion GuarddutyFilter#criterion}.
	Criterion interface{} `field:"optional" json:"criterion" yaml:"criterion"`
}

