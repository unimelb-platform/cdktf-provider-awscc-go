package apsworkspace


type ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsLimits struct {
	// The maximum number of active series that can be ingested for this label set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/aps_workspace#max_series ApsWorkspace#max_series}
	MaxSeries *float64 `field:"optional" json:"maxSeries" yaml:"maxSeries"`
}

