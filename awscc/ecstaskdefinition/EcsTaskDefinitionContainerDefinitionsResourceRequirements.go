package ecstaskdefinition


type EcsTaskDefinitionContainerDefinitionsResourceRequirements struct {
	// The type of resource to assign to a container.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ecs_task_definition#type EcsTaskDefinition#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
	// The value for the specified resource type.
	//
	// When the type is ``GPU``, the value is the number of physical ``GPUs`` the Amazon ECS container agent reserves for the container. The number of GPUs that's reserved for all containers in a task can't exceed the number of available GPUs on the container instance that the task is launched on.
	//  When the type is ``InferenceAccelerator``, the ``value`` matches the ``deviceName`` for an [InferenceAccelerator](https://docs.aws.amazon.com/AmazonECS/latest/APIReference/API_InferenceAccelerator.html) specified in a task definition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ecs_task_definition#value EcsTaskDefinition#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

