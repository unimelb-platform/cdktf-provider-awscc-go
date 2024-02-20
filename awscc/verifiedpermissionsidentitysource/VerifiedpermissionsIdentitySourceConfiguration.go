package verifiedpermissionsidentitysource


type VerifiedpermissionsIdentitySourceConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/verifiedpermissions_identity_source#cognito_user_pool_configuration VerifiedpermissionsIdentitySource#cognito_user_pool_configuration}.
	CognitoUserPoolConfiguration *VerifiedpermissionsIdentitySourceConfigurationCognitoUserPoolConfiguration `field:"required" json:"cognitoUserPoolConfiguration" yaml:"cognitoUserPoolConfiguration"`
}

