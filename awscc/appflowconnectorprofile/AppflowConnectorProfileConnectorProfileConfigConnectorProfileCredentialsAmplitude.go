package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsAmplitude struct {
	// A unique alphanumeric identi?er used to authenticate a user, developer, or calling program to your API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_connector_profile#api_key AppflowConnectorProfile#api_key}
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_connector_profile#secret_key AppflowConnectorProfile#secret_key}.
	SecretKey *string `field:"required" json:"secretKey" yaml:"secretKey"`
}

