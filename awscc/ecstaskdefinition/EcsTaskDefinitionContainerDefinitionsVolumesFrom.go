package ecstaskdefinition


type EcsTaskDefinitionContainerDefinitionsVolumesFrom struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#read_only EcsTaskDefinition#read_only}.
	ReadOnly interface{} `field:"optional" json:"readOnly" yaml:"readOnly"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#source_container EcsTaskDefinition#source_container}.
	SourceContainer *string `field:"optional" json:"sourceContainer" yaml:"sourceContainer"`
}

