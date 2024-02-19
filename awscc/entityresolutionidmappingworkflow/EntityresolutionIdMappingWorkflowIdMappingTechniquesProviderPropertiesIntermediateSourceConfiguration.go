package entityresolutionidmappingworkflow


type EntityresolutionIdMappingWorkflowIdMappingTechniquesProviderPropertiesIntermediateSourceConfiguration struct {
	// The s3 path that would be used to stage the intermediate data being generated during workflow execution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/entityresolution_id_mapping_workflow#intermediate_s3_path EntityresolutionIdMappingWorkflow#intermediate_s3_path}
	IntermediateS3Path *string `field:"required" json:"intermediateS3Path" yaml:"intermediateS3Path"`
}

