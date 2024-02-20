package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSapoData struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_connector_profile#basic_auth_credentials AppflowConnectorProfile#basic_auth_credentials}.
	BasicAuthCredentials *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSapoDataBasicAuthCredentials `field:"optional" json:"basicAuthCredentials" yaml:"basicAuthCredentials"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_connector_profile#o_auth_credentials AppflowConnectorProfile#o_auth_credentials}.
	OAuthCredentials *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSapoDataOAuthCredentials `field:"optional" json:"oAuthCredentials" yaml:"oAuthCredentials"`
}

