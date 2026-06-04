package sesmailmanagerruleset


type SesMailManagerRuleSetRulesActionsArchive struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ses_mail_manager_rule_set#action_failure_policy SesMailManagerRuleSet#action_failure_policy}.
	ActionFailurePolicy *string `field:"optional" json:"actionFailurePolicy" yaml:"actionFailurePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ses_mail_manager_rule_set#target_archive SesMailManagerRuleSet#target_archive}.
	TargetArchive *string `field:"optional" json:"targetArchive" yaml:"targetArchive"`
}

