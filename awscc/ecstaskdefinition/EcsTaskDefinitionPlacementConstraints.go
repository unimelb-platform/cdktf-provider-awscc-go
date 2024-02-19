package ecstaskdefinition


type EcsTaskDefinitionPlacementConstraints struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#type EcsTaskDefinition#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#expression EcsTaskDefinition#expression}.
	Expression *string `field:"optional" json:"expression" yaml:"expression"`
}

