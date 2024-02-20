package ecstaskdefinition


type EcsTaskDefinitionContainerDefinitionsHealthCheck struct {
	// A string array representing the command that the container runs to determine if it is healthy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#command EcsTaskDefinition#command}
	Command *[]*string `field:"optional" json:"command" yaml:"command"`
	// The time period in seconds between each health check execution.
	//
	// You may specify between 5 and 300 seconds. The default value is 30 seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#interval EcsTaskDefinition#interval}
	Interval *float64 `field:"optional" json:"interval" yaml:"interval"`
	// The number of times to retry a failed health check before the container is considered unhealthy.
	//
	// You may specify between 1 and 10 retries. The default value is three retries.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#retries EcsTaskDefinition#retries}
	Retries *float64 `field:"optional" json:"retries" yaml:"retries"`
	// The optional grace period within which to provide containers time to bootstrap before failed health checks count towards the maximum number of retries.
	//
	// You may specify between 0 and 300 seconds. The startPeriod is disabled by default.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#start_period EcsTaskDefinition#start_period}
	StartPeriod *float64 `field:"optional" json:"startPeriod" yaml:"startPeriod"`
	// The time period in seconds to wait for a health check to succeed before it is considered a failure.
	//
	// You may specify between 2 and 60 seconds. The default value is 5 seconds.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#timeout EcsTaskDefinition#timeout}
	Timeout *float64 `field:"optional" json:"timeout" yaml:"timeout"`
}

