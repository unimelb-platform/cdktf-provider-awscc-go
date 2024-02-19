package ecsservice


type EcsServiceVolumeConfigurations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_service#name EcsService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_service#managed_ebs_volume EcsService#managed_ebs_volume}.
	ManagedEbsVolume *EcsServiceVolumeConfigurationsManagedEbsVolume `field:"optional" json:"managedEbsVolume" yaml:"managedEbsVolume"`
}

