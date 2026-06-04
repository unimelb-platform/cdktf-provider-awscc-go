package verifiedpermissionsidentitysource


type VerifiedpermissionsIdentitySourceConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/verifiedpermissions_identity_source#cognito_user_pool_configuration VerifiedpermissionsIdentitySource#cognito_user_pool_configuration}.
	CognitoUserPoolConfiguration *VerifiedpermissionsIdentitySourceConfigurationCognitoUserPoolConfiguration `field:"optional" json:"cognitoUserPoolConfiguration" yaml:"cognitoUserPoolConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/verifiedpermissions_identity_source#open_id_connect_configuration VerifiedpermissionsIdentitySource#open_id_connect_configuration}.
	OpenIdConnectConfiguration *VerifiedpermissionsIdentitySourceConfigurationOpenIdConnectConfiguration `field:"optional" json:"openIdConnectConfiguration" yaml:"openIdConnectConfiguration"`
}

