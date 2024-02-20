package ecstaskdefinition


type EcsTaskDefinitionVolumes struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#configured_at_launch EcsTaskDefinition#configured_at_launch}.
	ConfiguredAtLaunch interface{} `field:"optional" json:"configuredAtLaunch" yaml:"configuredAtLaunch"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#docker_volume_configuration EcsTaskDefinition#docker_volume_configuration}.
	DockerVolumeConfiguration *EcsTaskDefinitionVolumesDockerVolumeConfiguration `field:"optional" json:"dockerVolumeConfiguration" yaml:"dockerVolumeConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#efs_volume_configuration EcsTaskDefinition#efs_volume_configuration}.
	EfsVolumeConfiguration *EcsTaskDefinitionVolumesEfsVolumeConfiguration `field:"optional" json:"efsVolumeConfiguration" yaml:"efsVolumeConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#host EcsTaskDefinition#host}.
	Host *EcsTaskDefinitionVolumesHost `field:"optional" json:"host" yaml:"host"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#name EcsTaskDefinition#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}

