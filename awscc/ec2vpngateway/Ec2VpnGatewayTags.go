package ec2vpngateway


type Ec2VpnGatewayTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_vpn_gateway#key Ec2VpnGateway#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_vpn_gateway#value Ec2VpnGateway#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

