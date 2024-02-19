package ecstaskdefinition


type EcsTaskDefinitionContainerDefinitionsDependsOn struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#condition EcsTaskDefinition#condition}.
	Condition *string `field:"optional" json:"condition" yaml:"condition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#container_name EcsTaskDefinition#container_name}.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
}

