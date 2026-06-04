package apsworkspace


type ApsWorkspaceQueryLoggingConfiguration struct {
	// The destinations configuration for query logging.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/aps_workspace#destinations ApsWorkspace#destinations}
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
}

