package ecscluster


type EcsClusterConfigurationExecuteCommandConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_cluster#kms_key_id EcsCluster#kms_key_id}.
	KmsKeyId *string `field:"optional" json:"kmsKeyId" yaml:"kmsKeyId"`
	// The session logging configuration for ExecuteCommand.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_cluster#log_configuration EcsCluster#log_configuration}
	LogConfiguration *EcsClusterConfigurationExecuteCommandConfigurationLogConfiguration `field:"optional" json:"logConfiguration" yaml:"logConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_cluster#logging EcsCluster#logging}.
	Logging *string `field:"optional" json:"logging" yaml:"logging"`
}

