package sesmailmanagerruleset


type SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ses_mail_manager_rule_set#analysis SesMailManagerRuleSet#analysis}.
	Analysis *SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateAnalysis `field:"optional" json:"analysis" yaml:"analysis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ses_mail_manager_rule_set#attribute SesMailManagerRuleSet#attribute}.
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ses_mail_manager_rule_set#is_in_address_list SesMailManagerRuleSet#is_in_address_list}.
	IsInAddressList *SesMailManagerRuleSetRulesUnlessBooleanExpressionEvaluateIsInAddressListStruct `field:"optional" json:"isInAddressList" yaml:"isInAddressList"`
}

