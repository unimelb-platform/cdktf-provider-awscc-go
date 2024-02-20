package ec2natgateway


type Ec2NatGatewayTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_nat_gateway#key Ec2NatGateway#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_nat_gateway#value Ec2NatGateway#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

