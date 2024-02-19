package pipespipe


type PipesPipeSourceParametersSelfManagedKafkaParametersVpc struct {
	// List of SecurityGroupId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pipes_pipe#security_group PipesPipe#security_group}
	SecurityGroup *[]*string `field:"optional" json:"securityGroup" yaml:"securityGroup"`
	// List of SubnetId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pipes_pipe#subnets PipesPipe#subnets}
	Subnets *[]*string `field:"optional" json:"subnets" yaml:"subnets"`
}

