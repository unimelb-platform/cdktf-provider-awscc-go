package networkfirewallrulegroup


type NetworkfirewallRuleGroupRuleGroupRulesSourceRulesSourceListStruct struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/networkfirewall_rule_group#generated_rules_type NetworkfirewallRuleGroup#generated_rules_type}.
	GeneratedRulesType *string `field:"optional" json:"generatedRulesType" yaml:"generatedRulesType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/networkfirewall_rule_group#targets NetworkfirewallRuleGroup#targets}.
	Targets *[]*string `field:"optional" json:"targets" yaml:"targets"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/networkfirewall_rule_group#target_types NetworkfirewallRuleGroup#target_types}.
	TargetTypes *[]*string `field:"optional" json:"targetTypes" yaml:"targetTypes"`
}

