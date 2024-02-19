package ec2internetgateway


type Ec2InternetGatewayTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_internet_gateway#key Ec2InternetGateway#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_internet_gateway#value Ec2InternetGateway#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

