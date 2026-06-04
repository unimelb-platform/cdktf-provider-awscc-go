package apsworkspace


type ApsWorkspaceWorkspaceConfiguration struct {
	// An array of label set and associated limits.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/aps_workspace#limits_per_label_sets ApsWorkspace#limits_per_label_sets}
	LimitsPerLabelSets interface{} `field:"optional" json:"limitsPerLabelSets" yaml:"limitsPerLabelSets"`
	// How many days that metrics are retained in the workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/aps_workspace#retention_period_in_days ApsWorkspace#retention_period_in_days}
	RetentionPeriodInDays *float64 `field:"optional" json:"retentionPeriodInDays" yaml:"retentionPeriodInDays"`
}

