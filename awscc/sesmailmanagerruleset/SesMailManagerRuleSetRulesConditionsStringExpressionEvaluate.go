package sesmailmanagerruleset


type SesMailManagerRuleSetRulesConditionsStringExpressionEvaluate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ses_mail_manager_rule_set#analysis SesMailManagerRuleSet#analysis}.
	Analysis *SesMailManagerRuleSetRulesConditionsStringExpressionEvaluateAnalysis `field:"optional" json:"analysis" yaml:"analysis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ses_mail_manager_rule_set#attribute SesMailManagerRuleSet#attribute}.
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ses_mail_manager_rule_set#mime_header_attribute SesMailManagerRuleSet#mime_header_attribute}.
	MimeHeaderAttribute *string `field:"optional" json:"mimeHeaderAttribute" yaml:"mimeHeaderAttribute"`
}

