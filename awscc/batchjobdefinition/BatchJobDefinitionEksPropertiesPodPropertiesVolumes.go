package batchjobdefinition


type BatchJobDefinitionEksPropertiesPodPropertiesVolumes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#name BatchJobDefinition#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#empty_dir BatchJobDefinition#empty_dir}.
	EmptyDir *BatchJobDefinitionEksPropertiesPodPropertiesVolumesEmptyDir `field:"optional" json:"emptyDir" yaml:"emptyDir"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#host_path BatchJobDefinition#host_path}.
	HostPath *BatchJobDefinitionEksPropertiesPodPropertiesVolumesHostPath `field:"optional" json:"hostPath" yaml:"hostPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#secret BatchJobDefinition#secret}.
	Secret *BatchJobDefinitionEksPropertiesPodPropertiesVolumesSecret `field:"optional" json:"secret" yaml:"secret"`
}

