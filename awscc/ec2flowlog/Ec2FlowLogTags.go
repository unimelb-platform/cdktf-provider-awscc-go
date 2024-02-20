package ec2flowlog


type Ec2FlowLogTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_flow_log#key Ec2FlowLog#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_flow_log#value Ec2FlowLog#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

