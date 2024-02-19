package fmspolicy


type FmsPolicySecurityServicePolicyDataPolicyOption struct {
	// Network firewall policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/fms_policy#network_firewall_policy FmsPolicy#network_firewall_policy}
	NetworkFirewallPolicy *FmsPolicySecurityServicePolicyDataPolicyOptionNetworkFirewallPolicy `field:"optional" json:"networkFirewallPolicy" yaml:"networkFirewallPolicy"`
	// Third party firewall policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/fms_policy#third_party_firewall_policy FmsPolicy#third_party_firewall_policy}
	ThirdPartyFirewallPolicy *FmsPolicySecurityServicePolicyDataPolicyOptionThirdPartyFirewallPolicy `field:"optional" json:"thirdPartyFirewallPolicy" yaml:"thirdPartyFirewallPolicy"`
}

