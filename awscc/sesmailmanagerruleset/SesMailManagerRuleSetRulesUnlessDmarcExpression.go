package sesmailmanagerruleset


type SesMailManagerRuleSetRulesUnlessDmarcExpression struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ses_mail_manager_rule_set#operator SesMailManagerRuleSet#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ses_mail_manager_rule_set#values SesMailManagerRuleSet#values}.
	Values *[]*string `field:"optional" json:"values" yaml:"values"`
}

