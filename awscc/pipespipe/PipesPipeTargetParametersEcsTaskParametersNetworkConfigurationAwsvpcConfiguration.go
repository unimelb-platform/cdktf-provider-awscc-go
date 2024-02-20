package pipespipe


type PipesPipeTargetParametersEcsTaskParametersNetworkConfigurationAwsvpcConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pipes_pipe#subnets PipesPipe#subnets}.
	Subnets *[]*string `field:"required" json:"subnets" yaml:"subnets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pipes_pipe#assign_public_ip PipesPipe#assign_public_ip}.
	AssignPublicIp *string `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pipes_pipe#security_groups PipesPipe#security_groups}.
	SecurityGroups *[]*string `field:"optional" json:"securityGroups" yaml:"securityGroups"`
}

