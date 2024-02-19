package networkfirewallrulegroup


type NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRules struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkfirewall_rule_group#action NetworkfirewallRuleGroup#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkfirewall_rule_group#header NetworkfirewallRuleGroup#header}.
	Header *NetworkfirewallRuleGroupRuleGroupRulesSourceStatefulRulesHeader `field:"required" json:"header" yaml:"header"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkfirewall_rule_group#rule_options NetworkfirewallRuleGroup#rule_options}.
	RuleOptions interface{} `field:"required" json:"ruleOptions" yaml:"ruleOptions"`
}

