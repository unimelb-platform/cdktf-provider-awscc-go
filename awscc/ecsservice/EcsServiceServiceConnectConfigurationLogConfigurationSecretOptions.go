package ecsservice


type EcsServiceServiceConnectConfigurationLogConfigurationSecretOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_service#name EcsService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_service#value_from EcsService#value_from}.
	ValueFrom *string `field:"required" json:"valueFrom" yaml:"valueFrom"`
}

