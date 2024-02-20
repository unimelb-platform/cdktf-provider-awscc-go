package efsfilesystem


type EfsFileSystemBackupPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/efs_file_system#status EfsFileSystem#status}.
	Status *string `field:"required" json:"status" yaml:"status"`
}

