package ecstaskdefinition


type EcsTaskDefinitionContainerDefinitionsLinuxParametersDevices struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#container_path EcsTaskDefinition#container_path}.
	ContainerPath *string `field:"optional" json:"containerPath" yaml:"containerPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#host_path EcsTaskDefinition#host_path}.
	HostPath *string `field:"optional" json:"hostPath" yaml:"hostPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#permissions EcsTaskDefinition#permissions}.
	Permissions *[]*string `field:"optional" json:"permissions" yaml:"permissions"`
}

