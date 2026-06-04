package workspacesinstancesworkspaceinstance


type WorkspacesinstancesWorkspaceInstanceManagedInstanceCpuOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/workspacesinstances_workspace_instance#core_count WorkspacesinstancesWorkspaceInstance#core_count}.
	CoreCount *float64 `field:"optional" json:"coreCount" yaml:"coreCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/workspacesinstances_workspace_instance#threads_per_core WorkspacesinstancesWorkspaceInstance#threads_per_core}.
	ThreadsPerCore *float64 `field:"optional" json:"threadsPerCore" yaml:"threadsPerCore"`
}

