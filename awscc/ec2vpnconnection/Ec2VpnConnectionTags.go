package ec2vpnconnection


type Ec2VpnConnectionTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_vpn_connection#key Ec2VpnConnection#key}.
	Key *string `field:"required" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/ec2_vpn_connection#value Ec2VpnConnection#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

