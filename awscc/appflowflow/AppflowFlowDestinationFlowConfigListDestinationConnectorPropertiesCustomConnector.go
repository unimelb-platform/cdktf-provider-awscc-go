package appflowflow


type AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesCustomConnector struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_flow#entity_name AppflowFlow#entity_name}.
	EntityName *string `field:"required" json:"entityName" yaml:"entityName"`
	// A map for properties for custom connector.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_flow#custom_properties AppflowFlow#custom_properties}
	CustomProperties *map[string]*string `field:"optional" json:"customProperties" yaml:"customProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_flow#error_handling_config AppflowFlow#error_handling_config}.
	ErrorHandlingConfig *AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesCustomConnectorErrorHandlingConfig `field:"optional" json:"errorHandlingConfig" yaml:"errorHandlingConfig"`
	// List of fields used as ID when performing a write operation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_flow#id_field_names AppflowFlow#id_field_names}
	IdFieldNames *[]*string `field:"optional" json:"idFieldNames" yaml:"idFieldNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_flow#write_operation_type AppflowFlow#write_operation_type}.
	WriteOperationType *string `field:"optional" json:"writeOperationType" yaml:"writeOperationType"`
}

