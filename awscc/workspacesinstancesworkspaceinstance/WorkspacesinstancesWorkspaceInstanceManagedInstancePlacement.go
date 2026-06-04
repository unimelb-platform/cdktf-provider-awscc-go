package workspacesinstancesworkspaceinstance


type WorkspacesinstancesWorkspaceInstanceManagedInstancePlacement struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/workspacesinstances_workspace_instance#availability_zone WorkspacesinstancesWorkspaceInstance#availability_zone}.
	AvailabilityZone *string `field:"optional" json:"availabilityZone" yaml:"availabilityZone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/workspacesinstances_workspace_instance#group_name WorkspacesinstancesWorkspaceInstance#group_name}.
	GroupName *string `field:"optional" json:"groupName" yaml:"groupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/workspacesinstances_workspace_instance#tenancy WorkspacesinstancesWorkspaceInstance#tenancy}.
	Tenancy *string `field:"optional" json:"tenancy" yaml:"tenancy"`
}

