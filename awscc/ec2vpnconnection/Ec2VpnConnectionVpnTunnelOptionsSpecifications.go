package ec2vpnconnection


type Ec2VpnConnectionVpnTunnelOptionsSpecifications struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_vpn_connection#pre_shared_key Ec2VpnConnection#pre_shared_key}.
	PreSharedKey *string `field:"optional" json:"preSharedKey" yaml:"preSharedKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_vpn_connection#tunnel_inside_cidr Ec2VpnConnection#tunnel_inside_cidr}.
	TunnelInsideCidr *string `field:"optional" json:"tunnelInsideCidr" yaml:"tunnelInsideCidr"`
}

