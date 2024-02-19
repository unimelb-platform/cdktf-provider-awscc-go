package grafanaworkspace


type GrafanaWorkspaceSamlConfigurationRoleValues struct {
	// List of SAML roles which will be mapped into the Grafana Admin role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/grafana_workspace#admin GrafanaWorkspace#admin}
	Admin *[]*string `field:"optional" json:"admin" yaml:"admin"`
	// List of SAML roles which will be mapped into the Grafana Editor role.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/grafana_workspace#editor GrafanaWorkspace#editor}
	Editor *[]*string `field:"optional" json:"editor" yaml:"editor"`
}

