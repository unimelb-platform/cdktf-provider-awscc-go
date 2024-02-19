package efsfilesystem


type EfsFileSystemReplicationConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/efs_file_system#destinations EfsFileSystem#destinations}.
	Destinations interface{} `field:"optional" json:"destinations" yaml:"destinations"`
}

