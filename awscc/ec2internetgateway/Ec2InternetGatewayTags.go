package ec2internetgateway


type Ec2InternetGatewayTags struct {
	// The tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_internet_gateway#key Ec2InternetGateway#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_internet_gateway#value Ec2InternetGateway#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

