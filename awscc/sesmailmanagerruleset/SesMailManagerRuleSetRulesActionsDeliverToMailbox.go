package sesmailmanagerruleset


type SesMailManagerRuleSetRulesActionsDeliverToMailbox struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ses_mail_manager_rule_set#action_failure_policy SesMailManagerRuleSet#action_failure_policy}.
	ActionFailurePolicy *string `field:"optional" json:"actionFailurePolicy" yaml:"actionFailurePolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ses_mail_manager_rule_set#mailbox_arn SesMailManagerRuleSet#mailbox_arn}.
	MailboxArn *string `field:"optional" json:"mailboxArn" yaml:"mailboxArn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ses_mail_manager_rule_set#role_arn SesMailManagerRuleSet#role_arn}.
	RoleArn *string `field:"optional" json:"roleArn" yaml:"roleArn"`
}

