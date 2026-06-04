package ecsservice


type EcsServiceDeploymentConfigurationDeploymentCircuitBreaker struct {
	// Determines whether to use the deployment circuit breaker logic for the service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ecs_service#enable EcsService#enable}
	Enable interface{} `field:"optional" json:"enable" yaml:"enable"`
	// Determines whether to configure Amazon ECS to roll back the service if a service deployment fails.
	//
	// If rollback is on, when a service deployment fails, the service is rolled back to the last deployment that completed successfully.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ecs_service#rollback EcsService#rollback}
	Rollback interface{} `field:"optional" json:"rollback" yaml:"rollback"`
}

