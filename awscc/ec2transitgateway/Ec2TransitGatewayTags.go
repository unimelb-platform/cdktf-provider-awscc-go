package ec2transitgateway


type Ec2TransitGatewayTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_transit_gateway#key Ec2TransitGateway#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_transit_gateway#value Ec2TransitGateway#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

