package networkfirewallrulegroup


type NetworkfirewallRuleGroupRuleGroupRuleVariables struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkfirewall_rule_group#ip_sets NetworkfirewallRuleGroup#ip_sets}.
	IpSets interface{} `field:"optional" json:"ipSets" yaml:"ipSets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkfirewall_rule_group#port_sets NetworkfirewallRuleGroup#port_sets}.
	PortSets interface{} `field:"optional" json:"portSets" yaml:"portSets"`
}

