package appflowflow


type AppflowFlowTasks struct {
	// Source fields on which particular task will be applied.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_flow#source_fields AppflowFlow#source_fields}
	SourceFields *[]*string `field:"required" json:"sourceFields" yaml:"sourceFields"`
	// Type of task.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_flow#task_type AppflowFlow#task_type}
	TaskType *string `field:"required" json:"taskType" yaml:"taskType"`
	// Operation to be performed on provided source fields.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_flow#connector_operator AppflowFlow#connector_operator}
	ConnectorOperator *AppflowFlowTasksConnectorOperator `field:"optional" json:"connectorOperator" yaml:"connectorOperator"`
	// A field value on which source field should be validated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_flow#destination_field AppflowFlow#destination_field}
	DestinationField *string `field:"optional" json:"destinationField" yaml:"destinationField"`
	// A Map used to store task related info.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/appflow_flow#task_properties AppflowFlow#task_properties}
	TaskProperties interface{} `field:"optional" json:"taskProperties" yaml:"taskProperties"`
}

