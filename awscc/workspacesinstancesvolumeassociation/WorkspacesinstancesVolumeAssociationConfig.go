package workspacesinstancesvolumeassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WorkspacesinstancesVolumeAssociationConfig struct {
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
	// The device name for the volume attachment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/workspacesinstances_volume_association#device WorkspacesinstancesVolumeAssociation#device}
	Device *string `field:"required" json:"device" yaml:"device"`
	// ID of the volume to attach to the workspace instance.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/workspacesinstances_volume_association#volume_id WorkspacesinstancesVolumeAssociation#volume_id}
	VolumeId *string `field:"required" json:"volumeId" yaml:"volumeId"`
	// ID of the workspace instance to associate with the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/workspacesinstances_volume_association#workspace_instance_id WorkspacesinstancesVolumeAssociation#workspace_instance_id}
	WorkspaceInstanceId *string `field:"required" json:"workspaceInstanceId" yaml:"workspaceInstanceId"`
	// Mode to use when disassociating the volume.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/workspacesinstances_volume_association#disassociate_mode WorkspacesinstancesVolumeAssociation#disassociate_mode}
	DisassociateMode *string `field:"optional" json:"disassociateMode" yaml:"disassociateMode"`
}

