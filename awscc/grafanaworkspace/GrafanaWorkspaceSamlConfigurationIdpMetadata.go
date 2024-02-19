package grafanaworkspace


type GrafanaWorkspaceSamlConfigurationIdpMetadata struct {
	// URL that vends the IdPs metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/grafana_workspace#url GrafanaWorkspace#url}
	Url *string `field:"optional" json:"url" yaml:"url"`
	// XML blob of the IdPs metadata.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/grafana_workspace#xml GrafanaWorkspace#xml}
	Xml *string `field:"optional" json:"xml" yaml:"xml"`
}

