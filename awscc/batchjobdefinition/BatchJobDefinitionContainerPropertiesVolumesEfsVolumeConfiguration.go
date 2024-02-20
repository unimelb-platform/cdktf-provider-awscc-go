package batchjobdefinition


type BatchJobDefinitionContainerPropertiesVolumesEfsVolumeConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#file_system_id BatchJobDefinition#file_system_id}.
	FileSystemId *string `field:"required" json:"fileSystemId" yaml:"fileSystemId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#authorization_config BatchJobDefinition#authorization_config}.
	AuthorizationConfig *BatchJobDefinitionContainerPropertiesVolumesEfsVolumeConfigurationAuthorizationConfig `field:"optional" json:"authorizationConfig" yaml:"authorizationConfig"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#root_directory BatchJobDefinition#root_directory}.
	RootDirectory *string `field:"optional" json:"rootDirectory" yaml:"rootDirectory"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#transit_encryption BatchJobDefinition#transit_encryption}.
	TransitEncryption *string `field:"optional" json:"transitEncryption" yaml:"transitEncryption"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#transit_encryption_port BatchJobDefinition#transit_encryption_port}.
	TransitEncryptionPort *float64 `field:"optional" json:"transitEncryptionPort" yaml:"transitEncryptionPort"`
}

