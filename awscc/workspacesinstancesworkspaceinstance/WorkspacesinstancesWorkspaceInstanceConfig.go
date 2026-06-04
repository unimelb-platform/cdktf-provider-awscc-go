package workspacesinstancesworkspaceinstance

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkspacesinstancesWorkspaceInstanceConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/workspacesinstances_workspace_instance#managed_instance WorkspacesinstancesWorkspaceInstance#managed_instance}.
	ManagedInstance *WorkspacesinstancesWorkspaceInstanceManagedInstance `field:"optional" json:"managedInstance" yaml:"managedInstance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/workspacesinstances_workspace_instance#tags WorkspacesinstancesWorkspaceInstance#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

