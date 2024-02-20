package efsfilesystem


type EfsFileSystemLifecyclePolicies struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/efs_file_system#transition_to_archive EfsFileSystem#transition_to_archive}.
	TransitionToArchive *string `field:"optional" json:"transitionToArchive" yaml:"transitionToArchive"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/efs_file_system#transition_to_ia EfsFileSystem#transition_to_ia}.
	TransitionToIa *string `field:"optional" json:"transitionToIa" yaml:"transitionToIa"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/efs_file_system#transition_to_primary_storage_class EfsFileSystem#transition_to_primary_storage_class}.
	TransitionToPrimaryStorageClass *string `field:"optional" json:"transitionToPrimaryStorageClass" yaml:"transitionToPrimaryStorageClass"`
}

