package sagemakerdomain


type SagemakerDomainDefaultUserSettingsCustomFileSystemConfigs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_domain#efs_file_system_config SagemakerDomain#efs_file_system_config}.
	EfsFileSystemConfig *SagemakerDomainDefaultUserSettingsCustomFileSystemConfigsEfsFileSystemConfig `field:"optional" json:"efsFileSystemConfig" yaml:"efsFileSystemConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_domain#fsx_lustre_file_system_config SagemakerDomain#fsx_lustre_file_system_config}.
	FsxLustreFileSystemConfig *SagemakerDomainDefaultUserSettingsCustomFileSystemConfigsFsxLustreFileSystemConfig `field:"optional" json:"fsxLustreFileSystemConfig" yaml:"fsxLustreFileSystemConfig"`
}

