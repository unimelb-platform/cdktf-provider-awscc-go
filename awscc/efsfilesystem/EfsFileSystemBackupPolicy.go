package efsfilesystem


type EfsFileSystemBackupPolicy struct {
	// Set the backup policy status for the file system.
	//
	// +  *ENABLED* - Turns automatic backups on for the file system.
	//   +  *DISABLED* - Turns automatic backups off for the file system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/efs_file_system#status EfsFileSystem#status}
	Status *string `field:"optional" json:"status" yaml:"status"`
}

