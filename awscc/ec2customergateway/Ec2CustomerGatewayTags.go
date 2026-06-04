package ec2customergateway


type Ec2CustomerGatewayTags struct {
	// The tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_customer_gateway#key Ec2CustomerGateway#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_customer_gateway#value Ec2CustomerGateway#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

