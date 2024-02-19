package customerprofilesintegration


type CustomerprofilesIntegrationFlowDefinitionSourceFlowConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_integration#connector_type CustomerprofilesIntegration#connector_type}.
	ConnectorType *string `field:"required" json:"connectorType" yaml:"connectorType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_integration#source_connector_properties CustomerprofilesIntegration#source_connector_properties}.
	SourceConnectorProperties *CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigSourceConnectorProperties `field:"required" json:"sourceConnectorProperties" yaml:"sourceConnectorProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_integration#connector_profile_name CustomerprofilesIntegration#connector_profile_name}.
	ConnectorProfileName *string `field:"optional" json:"connectorProfileName" yaml:"connectorProfileName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_integration#incremental_pull_config CustomerprofilesIntegration#incremental_pull_config}.
	IncrementalPullConfig *CustomerprofilesIntegrationFlowDefinitionSourceFlowConfigIncrementalPullConfig `field:"optional" json:"incrementalPullConfig" yaml:"incrementalPullConfig"`
}

