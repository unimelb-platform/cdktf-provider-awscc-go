package customerprofilesintegration


type CustomerprofilesIntegrationFlowDefinitionTasks struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_integration#source_fields CustomerprofilesIntegration#source_fields}.
	SourceFields *[]*string `field:"required" json:"sourceFields" yaml:"sourceFields"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_integration#task_type CustomerprofilesIntegration#task_type}.
	TaskType *string `field:"required" json:"taskType" yaml:"taskType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_integration#connector_operator CustomerprofilesIntegration#connector_operator}.
	ConnectorOperator *CustomerprofilesIntegrationFlowDefinitionTasksConnectorOperator `field:"optional" json:"connectorOperator" yaml:"connectorOperator"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_integration#destination_field CustomerprofilesIntegration#destination_field}.
	DestinationField *string `field:"optional" json:"destinationField" yaml:"destinationField"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/customerprofiles_integration#task_properties CustomerprofilesIntegration#task_properties}.
	TaskProperties interface{} `field:"optional" json:"taskProperties" yaml:"taskProperties"`
}

