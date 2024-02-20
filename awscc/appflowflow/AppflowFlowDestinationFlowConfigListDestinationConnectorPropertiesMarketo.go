package appflowflow


type AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesMarketo struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_flow#object AppflowFlow#object}.
	Object *string `field:"required" json:"object" yaml:"object"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_flow#error_handling_config AppflowFlow#error_handling_config}.
	ErrorHandlingConfig *AppflowFlowDestinationFlowConfigListDestinationConnectorPropertiesMarketoErrorHandlingConfig `field:"optional" json:"errorHandlingConfig" yaml:"errorHandlingConfig"`
}

