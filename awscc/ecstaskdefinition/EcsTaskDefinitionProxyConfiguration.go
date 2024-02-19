package ecstaskdefinition


type EcsTaskDefinitionProxyConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#container_name EcsTaskDefinition#container_name}.
	ContainerName *string `field:"required" json:"containerName" yaml:"containerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#proxy_configuration_properties EcsTaskDefinition#proxy_configuration_properties}.
	ProxyConfigurationProperties interface{} `field:"optional" json:"proxyConfigurationProperties" yaml:"proxyConfigurationProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_definition#type EcsTaskDefinition#type}.
	Type *string `field:"optional" json:"type" yaml:"type"`
}

