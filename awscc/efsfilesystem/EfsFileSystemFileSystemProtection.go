package efsfilesystem


type EfsFileSystemFileSystemProtection struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/efs_file_system#replication_overwrite_protection EfsFileSystem#replication_overwrite_protection}.
	ReplicationOverwriteProtection *string `field:"optional" json:"replicationOverwriteProtection" yaml:"replicationOverwriteProtection"`
}

