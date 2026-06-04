package datasynclocationfsxontap


type DatasyncLocationFsxOntapProtocolSmb struct {
	// The name of the Windows domain that the SMB server belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datasync_location_fsx_ontap#domain DatasyncLocationFsxOntap#domain}
	Domain *string `field:"optional" json:"domain" yaml:"domain"`
	// The mount options used by DataSync to access the SMB server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datasync_location_fsx_ontap#mount_options DatasyncLocationFsxOntap#mount_options}
	MountOptions *DatasyncLocationFsxOntapProtocolSmbMountOptions `field:"optional" json:"mountOptions" yaml:"mountOptions"`
	// The password of the user who can mount the share and has the permissions to access files and folders in the SMB share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datasync_location_fsx_ontap#password DatasyncLocationFsxOntap#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The user who can mount the share, has the permissions to access files and folders in the SMB share.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/datasync_location_fsx_ontap#user DatasyncLocationFsxOntap#user}
	User *string `field:"optional" json:"user" yaml:"user"`
}

