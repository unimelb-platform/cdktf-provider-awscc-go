package sesmailmanagertrafficpolicy


type SesMailManagerTrafficPolicyPolicyStatementsConditionsStringExpressionEvaluate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ses_mail_manager_traffic_policy#analysis SesMailManagerTrafficPolicy#analysis}.
	Analysis *SesMailManagerTrafficPolicyPolicyStatementsConditionsStringExpressionEvaluateAnalysis `field:"optional" json:"analysis" yaml:"analysis"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/ses_mail_manager_traffic_policy#attribute SesMailManagerTrafficPolicy#attribute}.
	Attribute *string `field:"optional" json:"attribute" yaml:"attribute"`
}

