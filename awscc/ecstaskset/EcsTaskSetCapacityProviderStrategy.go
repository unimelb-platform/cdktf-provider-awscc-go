package ecstaskset


type EcsTaskSetCapacityProviderStrategy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ecs_task_set#base EcsTaskSet#base}.
	Base *float64 `field:"optional" json:"base" yaml:"base"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ecs_task_set#capacity_provider EcsTaskSet#capacity_provider}.
	CapacityProvider *string `field:"optional" json:"capacityProvider" yaml:"capacityProvider"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ecs_task_set#weight EcsTaskSet#weight}.
	Weight *float64 `field:"optional" json:"weight" yaml:"weight"`
}

