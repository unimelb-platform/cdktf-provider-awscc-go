package ecstaskdefinition


type EcsTaskDefinitionRuntimePlatform struct {
	// The CPU architecture.
	//
	// You can run your Linux tasks on an ARM-based platform by setting the value to ``ARM64``. This option is available for tasks that run on Linux Amazon EC2 instance or Linux containers on Fargate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ecs_task_definition#cpu_architecture EcsTaskDefinition#cpu_architecture}
	CpuArchitecture *string `field:"optional" json:"cpuArchitecture" yaml:"cpuArchitecture"`
	// The operating system.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ecs_task_definition#operating_system_family EcsTaskDefinition#operating_system_family}
	OperatingSystemFamily *string `field:"optional" json:"operatingSystemFamily" yaml:"operatingSystemFamily"`
}

