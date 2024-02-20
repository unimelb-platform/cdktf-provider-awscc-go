package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsCustomConnectorCustom struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_connector_profile#custom_authentication_type AppflowConnectorProfile#custom_authentication_type}.
	CustomAuthenticationType *string `field:"required" json:"customAuthenticationType" yaml:"customAuthenticationType"`
	// A map for properties for custom authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_connector_profile#credentials_map AppflowConnectorProfile#credentials_map}
	CredentialsMap *map[string]*string `field:"optional" json:"credentialsMap" yaml:"credentialsMap"`
}

