package ecstaskset


type EcsTaskSetNetworkConfiguration struct {
	// The VPC subnets and security groups associated with a task.
	//
	// All specified subnets and security groups must be from the same VPC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ecs_task_set#aws_vpc_configuration EcsTaskSet#aws_vpc_configuration}
	AwsVpcConfiguration *EcsTaskSetNetworkConfigurationAwsVpcConfiguration `field:"optional" json:"awsVpcConfiguration" yaml:"awsVpcConfiguration"`
}

