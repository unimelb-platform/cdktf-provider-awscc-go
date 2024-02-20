package sagemakeruserprofile


type SagemakerUserProfileUserSettingsCustomFileSystemConfigsEfsFileSystemConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_user_profile#file_system_id SagemakerUserProfile#file_system_id}.
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_user_profile#file_system_path SagemakerUserProfile#file_system_path}.
	FileSystemPath *string `field:"optional" json:"fileSystemPath" yaml:"fileSystemPath"`
}

