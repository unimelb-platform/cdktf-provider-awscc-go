package sagemakerspace


type SagemakerSpaceSpaceSettingsCustomFileSystems struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_space#efs_file_system SagemakerSpace#efs_file_system}.
	EfsFileSystem *SagemakerSpaceSpaceSettingsCustomFileSystemsEfsFileSystem `field:"optional" json:"efsFileSystem" yaml:"efsFileSystem"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_space#fsx_lustre_file_system SagemakerSpace#fsx_lustre_file_system}.
	FsxLustreFileSystem *SagemakerSpaceSpaceSettingsCustomFileSystemsFsxLustreFileSystem `field:"optional" json:"fsxLustreFileSystem" yaml:"fsxLustreFileSystem"`
}

