package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfig struct {
	// Connector specific configuration needed to create connector profile based on Authentication mechanism.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_connector_profile#connector_profile_credentials AppflowConnectorProfile#connector_profile_credentials}
	ConnectorProfileCredentials *AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentials `field:"optional" json:"connectorProfileCredentials" yaml:"connectorProfileCredentials"`
	// Connector specific properties needed to create connector profile - currently not needed for Amplitude, Trendmicro, Googleanalytics and Singular.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_connector_profile#connector_profile_properties AppflowConnectorProfile#connector_profile_properties}
	ConnectorProfileProperties *AppflowConnectorProfileConnectorProfileConfigConnectorProfileProperties `field:"optional" json:"connectorProfileProperties" yaml:"connectorProfileProperties"`
}

