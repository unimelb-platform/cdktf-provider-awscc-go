package networkfirewallfirewall


type NetworkfirewallFirewallSubnetMappings struct {
	// A SubnetId.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkfirewall_firewall#subnet_id NetworkfirewallFirewall#subnet_id}
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// A IPAddressType.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkfirewall_firewall#ip_address_type NetworkfirewallFirewall#ip_address_type}
	IpAddressType *string `field:"optional" json:"ipAddressType" yaml:"ipAddressType"`
}

