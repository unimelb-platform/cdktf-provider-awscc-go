package ec2vpngateway


type Ec2VpnGatewayTags struct {
	// The tag key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_vpn_gateway#key Ec2VpnGateway#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// The tag value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_vpn_gateway#value Ec2VpnGateway#value}
	Value *string `field:"optional" json:"value" yaml:"value"`
}

