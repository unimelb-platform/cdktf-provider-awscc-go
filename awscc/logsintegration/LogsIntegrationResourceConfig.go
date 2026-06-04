package logsintegration


type LogsIntegrationResourceConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/logs_integration#open_search_resource_config LogsIntegration#open_search_resource_config}.
	OpenSearchResourceConfig *LogsIntegrationResourceConfigOpenSearchResourceConfig `field:"optional" json:"openSearchResourceConfig" yaml:"openSearchResourceConfig"`
}

