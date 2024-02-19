package datasynclocationsmb

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatasyncLocationSmbConfig struct {
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
	// The Amazon Resource Names (ARNs) of agents to use for a Simple Message Block (SMB) location.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_smb#agent_arns DatasyncLocationSmb#agent_arns}
	AgentArns *[]*string `field:"required" json:"agentArns" yaml:"agentArns"`
	// The user who can mount the share, has the permissions to access files and folders in the SMB share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_smb#user DatasyncLocationSmb#user}
	User *string `field:"required" json:"user" yaml:"user"`
	// The name of the Windows domain that the SMB server belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_smb#domain DatasyncLocationSmb#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// The mount options used by DataSync to access the SMB server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_smb#mount_options DatasyncLocationSmb#mount_options}
	MountOptions *DatasyncLocationSmbMountOptions `field:"optional" json:"mountOptions" yaml:"mountOptions"`
	// The password of the user who can mount the share and has the permissions to access files and folders in the SMB share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_smb#password DatasyncLocationSmb#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The name of the SMB server.
	//
	// This value is the IP address or Domain Name Service (DNS) name of the SMB server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_smb#server_hostname DatasyncLocationSmb#server_hostname}
	ServerHostname *string `field:"optional" json:"serverHostname" yaml:"serverHostname"`
	// The subdirectory in the SMB file system that is used to read data from the SMB source location or write data to the SMB destination.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_smb#subdirectory DatasyncLocationSmb#subdirectory}
	Subdirectory *string `field:"optional" json:"subdirectory" yaml:"subdirectory"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/datasync_location_smb#tags DatasyncLocationSmb#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

