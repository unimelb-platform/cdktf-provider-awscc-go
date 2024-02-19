package appflowconnectorprofile


type AppflowConnectorProfileConnectorProfileConfigConnectorProfileCredentialsRedshift struct {
	// The password that corresponds to the username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_connector_profile#password AppflowConnectorProfile#password}
	Password *string `field:"optional" json:"password" yaml:"password"`
	// The name of the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_connector_profile#username AppflowConnectorProfile#username}
	Username *string `field:"optional" json:"username" yaml:"username"`
}

