package batchjobdefinition


type BatchJobDefinitionNodePropertiesNodeRangePropertiesContainerRuntimePlatform struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#cpu_architecture BatchJobDefinition#cpu_architecture}.
	CpuArchitecture *string `field:"optional" json:"cpuArchitecture" yaml:"cpuArchitecture"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#operating_system_family BatchJobDefinition#operating_system_family}.
	OperatingSystemFamily *string `field:"optional" json:"operatingSystemFamily" yaml:"operatingSystemFamily"`
}

