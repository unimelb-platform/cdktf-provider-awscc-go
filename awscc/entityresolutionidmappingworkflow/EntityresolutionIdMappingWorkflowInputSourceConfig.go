package entityresolutionidmappingworkflow


type EntityresolutionIdMappingWorkflowInputSourceConfig struct {
	// An Glue table ARN for the input source table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/entityresolution_id_mapping_workflow#input_source_arn EntityresolutionIdMappingWorkflow#input_source_arn}
	InputSourceArn *string `field:"required" json:"inputSourceArn" yaml:"inputSourceArn"`
	// The SchemaMapping arn associated with the Schema.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/entityresolution_id_mapping_workflow#schema_arn EntityresolutionIdMappingWorkflow#schema_arn}
	SchemaArn *string `field:"required" json:"schemaArn" yaml:"schemaArn"`
}

