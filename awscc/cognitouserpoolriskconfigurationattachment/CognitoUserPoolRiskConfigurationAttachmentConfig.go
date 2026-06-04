package cognitouserpoolriskconfigurationattachment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CognitoUserPoolRiskConfigurationAttachmentConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cognito_user_pool_risk_configuration_attachment#client_id CognitoUserPoolRiskConfigurationAttachment#client_id}.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cognito_user_pool_risk_configuration_attachment#user_pool_id CognitoUserPoolRiskConfigurationAttachment#user_pool_id}.
	UserPoolId *string `field:"required" json:"userPoolId" yaml:"userPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cognito_user_pool_risk_configuration_attachment#account_takeover_risk_configuration CognitoUserPoolRiskConfigurationAttachment#account_takeover_risk_configuration}.
	AccountTakeoverRiskConfiguration *CognitoUserPoolRiskConfigurationAttachmentAccountTakeoverRiskConfiguration `field:"optional" json:"accountTakeoverRiskConfiguration" yaml:"accountTakeoverRiskConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cognito_user_pool_risk_configuration_attachment#compromised_credentials_risk_configuration CognitoUserPoolRiskConfigurationAttachment#compromised_credentials_risk_configuration}.
	CompromisedCredentialsRiskConfiguration *CognitoUserPoolRiskConfigurationAttachmentCompromisedCredentialsRiskConfiguration `field:"optional" json:"compromisedCredentialsRiskConfiguration" yaml:"compromisedCredentialsRiskConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cognito_user_pool_risk_configuration_attachment#risk_exception_configuration CognitoUserPoolRiskConfigurationAttachment#risk_exception_configuration}.
	RiskExceptionConfiguration *CognitoUserPoolRiskConfigurationAttachmentRiskExceptionConfiguration `field:"optional" json:"riskExceptionConfiguration" yaml:"riskExceptionConfiguration"`
}

