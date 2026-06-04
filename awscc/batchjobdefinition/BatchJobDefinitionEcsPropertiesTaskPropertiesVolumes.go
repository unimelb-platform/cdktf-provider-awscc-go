package batchjobdefinition


type BatchJobDefinitionEcsPropertiesTaskPropertiesVolumes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/batch_job_definition#efs_volume_configuration BatchJobDefinition#efs_volume_configuration}.
	EfsVolumeConfiguration *BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesEfsVolumeConfiguration `field:"optional" json:"efsVolumeConfiguration" yaml:"efsVolumeConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/batch_job_definition#host BatchJobDefinition#host}.
	Host *BatchJobDefinitionEcsPropertiesTaskPropertiesVolumesHost `field:"optional" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/batch_job_definition#name BatchJobDefinition#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

