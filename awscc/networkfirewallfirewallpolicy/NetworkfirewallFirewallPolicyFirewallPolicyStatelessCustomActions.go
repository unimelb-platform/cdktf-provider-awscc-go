package networkfirewallfirewallpolicy


type NetworkfirewallFirewallPolicyFirewallPolicyStatelessCustomActions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkfirewall_firewall_policy#action_definition NetworkfirewallFirewallPolicy#action_definition}.
	ActionDefinition *NetworkfirewallFirewallPolicyFirewallPolicyStatelessCustomActionsActionDefinition `field:"required" json:"actionDefinition" yaml:"actionDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkfirewall_firewall_policy#action_name NetworkfirewallFirewallPolicy#action_name}.
	ActionName *string `field:"required" json:"actionName" yaml:"actionName"`
}

