package batchjobdefinition


type BatchJobDefinitionContainerPropertiesLinuxParametersDevices struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#container_path BatchJobDefinition#container_path}.
	ContainerPath *string `field:"optional" json:"containerPath" yaml:"containerPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#host_path BatchJobDefinition#host_path}.
	HostPath *string `field:"optional" json:"hostPath" yaml:"hostPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#permissions BatchJobDefinition#permissions}.
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
}

