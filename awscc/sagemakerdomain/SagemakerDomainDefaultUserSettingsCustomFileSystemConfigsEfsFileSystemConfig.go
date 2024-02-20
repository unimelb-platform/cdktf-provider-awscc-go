package sagemakerdomain


type SagemakerDomainDefaultUserSettingsCustomFileSystemConfigsEfsFileSystemConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_domain#file_system_id SagemakerDomain#file_system_id}.
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_domain#file_system_path SagemakerDomain#file_system_path}.
	FileSystemPath *string `field:"optional" json:"fileSystemPath" yaml:"fileSystemPath"`
}

