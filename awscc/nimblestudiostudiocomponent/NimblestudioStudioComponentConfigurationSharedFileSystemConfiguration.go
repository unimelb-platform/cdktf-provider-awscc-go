package nimblestudiostudiocomponent


type NimblestudioStudioComponentConfigurationSharedFileSystemConfiguration struct {
	// <p>The endpoint of the shared file system that is accessed by the studio component             resource.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio_component#endpoint NimblestudioStudioComponent#endpoint}
	Endpoint *string `field:"optional" json:"endpoint" yaml:"endpoint"`
	// <p>The unique identifier for a file system.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio_component#file_system_id NimblestudioStudioComponent#file_system_id}
	FileSystemId *string `field:"optional" json:"fileSystemId" yaml:"fileSystemId"`
	// <p>The mount location for a shared file system on a Linux virtual workstation.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio_component#linux_mount_point NimblestudioStudioComponent#linux_mount_point}
	LinuxMountPoint *string `field:"optional" json:"linuxMountPoint" yaml:"linuxMountPoint"`
	// <p>The name of the file share.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio_component#share_name NimblestudioStudioComponent#share_name}
	ShareName *string `field:"optional" json:"shareName" yaml:"shareName"`
	// <p>The mount location for a shared file system on a Windows virtual workstation.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/nimblestudio_studio_component#windows_mount_drive NimblestudioStudioComponent#windows_mount_drive}
	WindowsMountDrive *string `field:"optional" json:"windowsMountDrive" yaml:"windowsMountDrive"`
}

