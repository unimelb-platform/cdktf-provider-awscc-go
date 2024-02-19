package ecstaskset


type EcsTaskSetScale struct {
	// The unit of measure for the scale value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_set#unit EcsTaskSet#unit}
	Unit *string `field:"optional" json:"unit" yaml:"unit"`
	// The value, specified as a percent total of a service's desiredCount, to scale the task set.
	//
	// Accepted values are numbers between 0 and 100.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_set#value EcsTaskSet#value}
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

