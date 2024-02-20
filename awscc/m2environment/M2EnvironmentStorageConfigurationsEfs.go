package m2environment


type M2EnvironmentStorageConfigurationsEfs struct {
	// The file system identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/m2_environment#file_system_id M2Environment#file_system_id}
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// The mount point for the file system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/m2_environment#mount_point M2Environment#mount_point}
	MountPoint *string `field:"required" json:"mountPoint" yaml:"mountPoint"`
}

