package appflowflow


type AppflowFlowSourceFlowConfigSourceConnectorPropertiesCustomConnector struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_flow#entity_name AppflowFlow#entity_name}.
	EntityName *string `field:"required" json:"entityName" yaml:"entityName"`
	// A map for properties for custom connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_flow#custom_properties AppflowFlow#custom_properties}
	CustomProperties *map[string]*string `field:"optional" json:"customProperties" yaml:"customProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_flow#data_transfer_api AppflowFlow#data_transfer_api}.
	DataTransferApi *AppflowFlowSourceFlowConfigSourceConnectorPropertiesCustomConnectorDataTransferApi `field:"optional" json:"dataTransferApi" yaml:"dataTransferApi"`
}

