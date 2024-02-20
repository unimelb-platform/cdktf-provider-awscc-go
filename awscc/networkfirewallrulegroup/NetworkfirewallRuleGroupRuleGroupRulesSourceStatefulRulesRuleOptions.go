package networkfirewallrulegroup


type NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesRuleOptions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkfirewall_rule_group#keyword NetworkfirewallRuleGroup#keyword}.
	Keyword *string `field:"required" json:"keyword" yaml:"keyword"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkfirewall_rule_group#settings NetworkfirewallRuleGroup#settings}.
	Settings *[]*string `field:"optional" json:"settings" yaml:"settings"`
}

