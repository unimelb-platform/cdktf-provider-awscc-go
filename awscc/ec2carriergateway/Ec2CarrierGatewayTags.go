package ec2carriergateway


type Ec2CarrierGatewayTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_carrier_gateway#key Ec2CarrierGateway#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_carrier_gateway#value Ec2CarrierGateway#value}.
	Value *string `field:"optional" json:"value" yaml:"value"`
}

