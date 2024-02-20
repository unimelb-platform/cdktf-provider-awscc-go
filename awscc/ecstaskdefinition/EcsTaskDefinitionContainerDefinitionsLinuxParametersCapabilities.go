package ecstaskdefinition


type EcsTaskDefinitionContainerDefinitionsLinuxParametersCapabilities struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#add EcsTaskDefinition#add}.
	Add *[]*string `field:"optional" json:"add" yaml:"add"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#drop EcsTaskDefinition#drop}.
	Drop *[]*string `field:"optional" json:"drop" yaml:"drop"`
}

