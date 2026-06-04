package apsworkspace


type ApsWorkspaceWorkspaceConfigurationLimitsPerLabelSetsLabelSet struct {
	// Name of the label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/aps_workspace#name ApsWorkspace#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Value of the label.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/aps_workspace#value ApsWorkspace#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

