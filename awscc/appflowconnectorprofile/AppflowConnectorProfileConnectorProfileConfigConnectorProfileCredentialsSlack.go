package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSlack struct {
	// The identi?er for the desired client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_connector_profile#client_id AppflowConnectorProfile#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The client secret used by the oauth client to authenticate to the authorization server.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_connector_profile#client_secret AppflowConnectorProfile#client_secret}
	ClientSecret *string `field:"required" json:"clientSecret" yaml:"clientSecret"`
	// The credentials used to access protected resources.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_connector_profile#access_token AppflowConnectorProfile#access_token}
	AccessToken *string `field:"optional" json:"accessToken" yaml:"accessToken"`
	// The oauth needed to request security tokens from the connector endpoint.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_connector_profile#connector_o_auth_request AppflowConnectorProfile#connector_o_auth_request}
	ConnectorOAuthRequest *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsSlackConnectorOAuthRequest `field:"optional" json:"connectorOAuthRequest" yaml:"connectorOAuthRequest"`
}

