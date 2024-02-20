package ecstaskdefinition


type EcsTaskDefinitionContainerDefinitionsUlimits struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#hard_limit EcsTaskDefinition#hard_limit}.
	HardLimit *float64 `field:"required" json:"hardLimit" yaml:"hardLimit"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#name EcsTaskDefinition#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#soft_limit EcsTaskDefinition#soft_limit}.
	SoftLimit *float64 `field:"required" json:"softLimit" yaml:"softLimit"`
}

