package ecsservice


type EcsServiceServiceConnectConfigurationServices struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_service#port_name EcsService#port_name}.
	PortName *string `field:"required" json:"portName" yaml:"portName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_service#client_aliases EcsService#client_aliases}.
	ClientAliases interface{} `field:"optional" json:"clientAliases" yaml:"clientAliases"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_service#discovery_name EcsService#discovery_name}.
	DiscoveryName *string `field:"optional" json:"discoveryName" yaml:"discoveryName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_service#ingress_port_override EcsService#ingress_port_override}.
	IngressPortOverride *float64 `field:"optional" json:"ingressPortOverride" yaml:"ingressPortOverride"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_service#timeout EcsService#timeout}.
	Timeout *EcsServiceServiceConnectConfigurationServicesTimeout `field:"optional" json:"timeout" yaml:"timeout"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_service#tls EcsService#tls}.
	Tls *EcsServiceServiceConnectConfigurationServicesTls `field:"optional" json:"tls" yaml:"tls"`
}

