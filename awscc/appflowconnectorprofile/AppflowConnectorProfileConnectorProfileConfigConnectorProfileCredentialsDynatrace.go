package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsDynatrace struct {
	// The API tokens used by Dynatrace API to authenticate various API calls.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_connector_profile#api_token AppflowConnectorProfile#api_token}
	ApiToken *string `field:"required" json:"apiToken" yaml:"apiToken"`
}

