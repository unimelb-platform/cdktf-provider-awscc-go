package batchjobdefinition


type BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerLinuxParametersTmpfs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#container_path BatchJobDefinition#container_path}.
	ContainerPath *string `field:"required" json:"containerPath" yaml:"containerPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#size BatchJobDefinition#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#mount_options BatchJobDefinition#mount_options}.
	MountOptions *[]*string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
}

