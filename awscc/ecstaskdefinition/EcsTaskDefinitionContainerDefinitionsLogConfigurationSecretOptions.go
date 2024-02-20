package ecstaskdefinition


type EcsTaskDefinitionContainerDefinitionsLogConfigurationSecretOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#name EcsTaskDefinition#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#value_from EcsTaskDefinition#value_from}.
	ValueFrom *string `field:"required" json:"valueFrom" yaml:"valueFrom"`
}

