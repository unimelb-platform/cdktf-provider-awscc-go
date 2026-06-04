package cognitouserpoolriskconfigurationattachment


type CognitoUserPoolRiskConfigurationAttachmentAccountTakeoverRiskConfigurationActions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cognito_user_pool_risk_configuration_attachment#high_action CognitoUserPoolRiskConfigurationAttachment#high_action}.
	HighAction *CognitoUserPoolRiskConfigurationAttachmentAccountTakeoverRiskConfigurationActionsHighAction `field:"optional" json:"highAction" yaml:"highAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cognito_user_pool_risk_configuration_attachment#low_action CognitoUserPoolRiskConfigurationAttachment#low_action}.
	LowAction *CognitoUserPoolRiskConfigurationAttachmentAccountTakeoverRiskConfigurationActionsLowAction `field:"optional" json:"lowAction" yaml:"lowAction"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cognito_user_pool_risk_configuration_attachment#medium_action CognitoUserPoolRiskConfigurationAttachment#medium_action}.
	MediumAction *CognitoUserPoolRiskConfigurationAttachmentAccountTakeoverRiskConfigurationActionsMediumAction `field:"optional" json:"mediumAction" yaml:"mediumAction"`
}

