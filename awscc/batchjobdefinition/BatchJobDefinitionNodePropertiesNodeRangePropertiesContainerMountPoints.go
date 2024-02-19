package batchjobdefinition


type BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerMountPoints struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#container_path BatchJobDefinition#container_path}.
	ContainerPath *string `field:"optional" json:"containerPath" yaml:"containerPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#read_only BatchJobDefinition#read_only}.
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#source_volume BatchJobDefinition#source_volume}.
	SourceVolume *string `field:"optional" json:"sourceVolume" yaml:"sourceVolume"`
}

