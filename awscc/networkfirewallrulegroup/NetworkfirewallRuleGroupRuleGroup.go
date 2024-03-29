package networkfirewallrulegroup


type NetworkfirewallRuleGroupRuleGroup struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkfirewall_rule_group#rules_source NetworkfirewallRuleGroup#rules_source}.
	RulesSource *NetworkfirewallRuleGroupRuleGroupRulesSource `field:"required" json:"rulesSource" yaml:"rulesSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkfirewall_rule_group#reference_sets NetworkfirewallRuleGroup#reference_sets}.
	ReferenceSets *NetworkfirewallRuleGroupRuleGroupReferenceSets `field:"optional" json:"referenceSets" yaml:"referenceSets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkfirewall_rule_group#rule_variables NetworkfirewallRuleGroup#rule_variables}.
	RuleVariables *NetworkfirewallRuleGroupRuleGroupRuleVariables `field:"optional" json:"ruleVariables" yaml:"ruleVariables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/networkfirewall_rule_group#stateful_rule_options NetworkfirewallRuleGroup#stateful_rule_options}.
	StatefulRuleOptions *NetworkfirewallRuleGroupRuleGroupStatefulRuleOptions `field:"optional" json:"statefulRuleOptions" yaml:"statefulRuleOptions"`
}

