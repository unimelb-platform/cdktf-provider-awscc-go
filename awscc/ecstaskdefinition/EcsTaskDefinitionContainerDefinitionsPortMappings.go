package ecstaskdefinition


type EcsTaskDefinitionContainerDefinitionsPortMappings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#app_protocol EcsTaskDefinition#app_protocol}.
	AppProtocol *string `field:"optional" json:"appProtocol" yaml:"appProtocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#container_port EcsTaskDefinition#container_port}.
	ContainerPort *float64 `field:"optional" json:"containerPort" yaml:"containerPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#container_port_range EcsTaskDefinition#container_port_range}.
	ContainerPortRange *string `field:"optional" json:"containerPortRange" yaml:"containerPortRange"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#host_port EcsTaskDefinition#host_port}.
	HostPort *float64 `field:"optional" json:"hostPort" yaml:"hostPort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#name EcsTaskDefinition#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#protocol EcsTaskDefinition#protocol}.
	Protocol *string `field:"optional" json:"protocol" yaml:"protocol"`
}

