package apsworkspace


type ApsWorkspaceQueryLoggingConfigurationDestinations struct {
	// Represents a cloudwatch logs destination for query logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/aps_workspace#cloudwatch_logs ApsWorkspace#cloudwatch_logs}
	CloudwatchLogs *ApsWorkspaceQueryLoggingConfigurationDestinationsCloudwatchLogs `field:"optional" json:"cloudwatchLogs" yaml:"cloudwatchLogs"`
	// Filters for logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/aps_workspace#filters ApsWorkspace#filters}
	Filters *ApsWorkspaceQueryLoggingConfigurationDestinationsFilters `field:"optional" json:"filters" yaml:"filters"`
}

