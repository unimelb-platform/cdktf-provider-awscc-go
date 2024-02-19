package ecsservice


type EcsServiceNetworkConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_service#awsvpc_configuration EcsService#awsvpc_configuration}.
	AwsvpcConfiguration *EcsServiceNetworkConfigurationAwsvpcConfiguration `field:"optional" json:"awsvpcConfiguration" yaml:"awsvpcConfiguration"`
}

