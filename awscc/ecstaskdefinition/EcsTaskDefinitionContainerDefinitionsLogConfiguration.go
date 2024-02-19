package ecstaskdefinition


type EcsTaskDefinitionContainerDefinitionsLogConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#log_driver EcsTaskDefinition#log_driver}.
	LogDriver *string `field:"required" json:"logDriver" yaml:"logDriver"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#options EcsTaskDefinition#options}.
	Options *map[string]*string `field:"optional" json:"options" yaml:"options"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#secret_options EcsTaskDefinition#secret_options}.
	SecretOptions interface{} `field:"optional" json:"secretOptions" yaml:"secretOptions"`
}

