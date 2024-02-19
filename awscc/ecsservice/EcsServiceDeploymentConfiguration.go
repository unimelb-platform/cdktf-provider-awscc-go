package ecsservice


type EcsServiceDeploymentConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_service#alarms EcsService#alarms}.
	Alarms *EcsServiceDeploymentConfigurationAlarms `field:"optional" json:"alarms" yaml:"alarms"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_service#deployment_circuit_breaker EcsService#deployment_circuit_breaker}.
	DeploymentCircuitBreaker *EcsServiceDeploymentConfigurationDeploymentCircuitBreaker `field:"optional" json:"deploymentCircuitBreaker" yaml:"deploymentCircuitBreaker"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_service#maximum_percent EcsService#maximum_percent}.
	MaximumPercent *float64 `field:"optional" json:"maximumPercent" yaml:"maximumPercent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_service#minimum_healthy_percent EcsService#minimum_healthy_percent}.
	MinimumHealthyPercent *float64 `field:"optional" json:"minimumHealthyPercent" yaml:"minimumHealthyPercent"`
}

