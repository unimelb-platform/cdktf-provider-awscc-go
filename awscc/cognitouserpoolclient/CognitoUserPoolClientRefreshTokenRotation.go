package cognitouserpoolclient


type CognitoUserPoolClientRefreshTokenRotation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cognito_user_pool_client#feature CognitoUserPoolClient#feature}.
	Feature *string `field:"optional" json:"feature" yaml:"feature"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/cognito_user_pool_client#retry_grace_period_seconds CognitoUserPoolClient#retry_grace_period_seconds}.
	RetryGracePeriodSeconds *float64 `field:"optional" json:"retryGracePeriodSeconds" yaml:"retryGracePeriodSeconds"`
}

