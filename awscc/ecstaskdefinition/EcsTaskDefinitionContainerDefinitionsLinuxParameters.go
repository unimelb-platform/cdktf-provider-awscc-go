package ecstaskdefinition


type EcsTaskDefinitionContainerDefinitionsLinuxParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#capabilities EcsTaskDefinition#capabilities}.
	Capabilities *EcsTaskDefinitionContainerDefinitionsLinuxParametersCapabilities `field:"optional" json:"capabilities" yaml:"capabilities"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#devices EcsTaskDefinition#devices}.
	Devices interface{} `field:"optional" json:"devices" yaml:"devices"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#init_process_enabled EcsTaskDefinition#init_process_enabled}.
	InitProcessEnabled interface{} `field:"optional" json:"initProcessEnabled" yaml:"initProcessEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#max_swap EcsTaskDefinition#max_swap}.
	MaxSwap *float64 `field:"optional" json:"maxSwap" yaml:"maxSwap"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#shared_memory_size EcsTaskDefinition#shared_memory_size}.
	SharedMemorySize *float64 `field:"optional" json:"sharedMemorySize" yaml:"sharedMemorySize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#swappiness EcsTaskDefinition#swappiness}.
	Swappiness *float64 `field:"optional" json:"swappiness" yaml:"swappiness"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#tmpfs EcsTaskDefinition#tmpfs}.
	Tmpfs interface{} `field:"optional" json:"tmpfs" yaml:"tmpfs"`
}

