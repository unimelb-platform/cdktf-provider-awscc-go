package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsGoogleAnalyticsConnectorOAuthRequest struct {
	// The code provided by the connector when it has been authenticated via the connected app.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_connector_profile#auth_code AppflowConnectorProfile#auth_code}
	AuthCode *string `field:"optional" json:"authCode" yaml:"authCode"`
	// The URL to which the authentication server redirects the browser after authorization has been granted.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_connector_profile#redirect_uri AppflowConnectorProfile#redirect_uri}
	RedirectUri *string `field:"optional" json:"redirectUri" yaml:"redirectUri"`
}

