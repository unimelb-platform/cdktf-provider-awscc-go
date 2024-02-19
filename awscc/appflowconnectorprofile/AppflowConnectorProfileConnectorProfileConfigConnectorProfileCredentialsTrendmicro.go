package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsTrendmicro struct {
	// The Secret Access Key portion of the credentials.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_connector_profile#api_secret_key AppflowConnectorProfile#api_secret_key}
	ApiSecretKey *string `field:"required" json:"apiSecretKey" yaml:"apiSecretKey"`
}

