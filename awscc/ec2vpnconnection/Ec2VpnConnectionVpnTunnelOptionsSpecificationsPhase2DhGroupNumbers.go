package ec2vpnconnection


type Ec2VpnConnectionVpnTunnelOptionsSpecificationsPhase2DhGroupNumbers struct {
	// The Diffie-Hellmann group number.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ec2_vpn_connection#value Ec2VpnConnection#value}
	Value *float64 `field:"optional" json:"value" yaml:"value"`
}

