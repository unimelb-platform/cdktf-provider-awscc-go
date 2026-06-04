package cognitouserpoolriskconfigurationattachment


type CognitoUserPoolRiskConfigurationAttachmentAccountTakeoverRiskConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cognito_user_pool_risk_configuration_attachment#actions CognitoUserPoolRiskConfigurationAttachment#actions}.
	Actions *CognitoUserPoolRiskConfigurationAttachmentAccountTakeoverRiskConfigurationActions `field:"optional" json:"actions" yaml:"actions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cognito_user_pool_risk_configuration_attachment#notify_configuration CognitoUserPoolRiskConfigurationAttachment#notify_configuration}.
	NotifyConfiguration *CognitoUserPoolRiskConfigurationAttachmentAccountTakeoverRiskConfigurationNotifyConfiguration `field:"optional" json:"notifyConfiguration" yaml:"notifyConfiguration"`
}

