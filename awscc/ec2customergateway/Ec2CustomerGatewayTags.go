package ec2customergateway


type Ec2CustomerGatewayTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_customer_gateway#key Ec2CustomerGateway#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_customer_gateway#value Ec2CustomerGateway#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

