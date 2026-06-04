package ec2natgateway


type Ec2NatGatewayTags struct {
	// The tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_nat_gateway#key Ec2NatGateway#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_nat_gateway#value Ec2NatGateway#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

