package apsworkspace


type ApsWorkspaceQueryLoggingConfigurationDestinationsFilters struct {
	// Query logs with QSP above this limit are vended.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/aps_workspace#qsp_threshold ApsWorkspace#qsp_threshold}
	QspThreshold *float64 `field:"optional" json:"qspThreshold" yaml:"qspThreshold"`
}

