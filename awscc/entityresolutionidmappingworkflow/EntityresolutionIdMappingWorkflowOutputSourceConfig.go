package entityresolutionidmappingworkflow


type EntityresolutionIdMappingWorkflowOutputSourceConfig struct {
	// The S3 path to which Entity Resolution will write the output table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/entityresolution_id_mapping_workflow#output_s3_path EntityresolutionIdMappingWorkflow#output_s3_path}
	OutputS3Path *string `field:"required" json:"outputS3Path" yaml:"outputS3Path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/entityresolution_id_mapping_workflow#kms_arn EntityresolutionIdMappingWorkflow#kms_arn}.
	KmsArn *string `field:"optional" json:"kmsArn" yaml:"kmsArn"`
}

