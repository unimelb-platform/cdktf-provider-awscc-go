package ecscluster


type EcsClusterConfiguration struct {
	// The configuration for ExecuteCommand.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_cluster#execute_command_configuration EcsCluster#execute_command_configuration}
	ExecuteCommandConfiguration *EcsClusterConfigurationExecuteCommandConfiguration `field:"optional" json:"executeCommandConfiguration" yaml:"executeCommandConfiguration"`
}

