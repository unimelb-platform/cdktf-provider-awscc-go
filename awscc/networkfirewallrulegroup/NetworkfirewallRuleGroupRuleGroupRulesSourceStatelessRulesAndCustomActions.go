package networkfirewallrulegroup


type NetworkfirewallRuleGroupRuleGroupRulesSourceStatelessRulesAndCustomActions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkfirewall_rule_group#stateless_rules NetworkfirewallRuleGroup#stateless_rules}.
	StatelessRules interface{} `field:"required" json:"statelessRules" yaml:"statelessRules"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkfirewall_rule_group#custom_actions NetworkfirewallRuleGroup#custom_actions}.
	CustomActions interface{} `field:"optional" json:"customActions" yaml:"customActions"`
}

