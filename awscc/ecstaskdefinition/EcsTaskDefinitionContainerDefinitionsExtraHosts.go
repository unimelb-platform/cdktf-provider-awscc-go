package ecstaskdefinition


type EcsTaskDefinitionContainerDefinitionsExtraHosts struct {
	// The hostname to use in the ``/etc/hosts`` entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ecs_task_definition#hostname EcsTaskDefinition#hostname}
	Hostname *string `field:"optional" json:"hostname" yaml:"hostname"`
	// The IP address to use in the ``/etc/hosts`` entry.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ecs_task_definition#ip_address EcsTaskDefinition#ip_address}
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
}

