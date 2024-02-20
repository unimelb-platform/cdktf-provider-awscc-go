package ecsservice


type EcsServiceServiceRegistries struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_service#container_name EcsService#container_name}.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_service#container_port EcsService#container_port}.
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_service#port EcsService#port}.
	Port *float64 `field:"optional" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_service#registry_arn EcsService#registry_arn}.
	RegistryArn *string `field:"optional" json:"registryArn" yaml:"registryArn"`
}

