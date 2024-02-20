package ecstaskdefinition


type EcsTaskDefinitionEphemeralStorage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#size_in_gi_b EcsTaskDefinition#size_in_gi_b}.
	SizeInGiB *float64 `field:"optional" json:"sizeInGiB" yaml:"sizeInGiB"`
}

