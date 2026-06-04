package sesmailmanagerruleset


type SesMailManagerRuleSetRulesUnlessBooleanExpression struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ses_mail_manager_rule_set#evaluate SesMailManagerRuleSet#evaluate}.
	Evaluate *SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluate `field:"optional" json:"evaluate" yaml:"evaluate"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ses_mail_manager_rule_set#operator SesMailManagerRuleSet#operator}.
	Operator *string `field:"optional" json:"operator" yaml:"operator"`
}

