package ecstaskdefinition


type EcsTaskDefinitionContainerDefinitionsLinuxParametersTmpfs struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#size EcsTaskDefinition#size}.
	Size *float64 `field:"required" json:"size" yaml:"size"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#container_path EcsTaskDefinition#container_path}.
	ContainerPath *string `field:"optional" json:"containerPath" yaml:"containerPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#mount_options EcsTaskDefinition#mount_options}.
	MountOptions *[]*string `field:"optional" json:"mountOptions" yaml:"mountOptions"`
}

