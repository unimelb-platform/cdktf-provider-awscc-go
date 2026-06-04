package entityresolutionmatchingworkflow


type EntityresolutionMatchingWorkflowOutputSourceConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/entityresolution_matching_workflow#output EntityresolutionMatchingWorkflow#output}.
	Output interface{} `field:"required" json:"output" yaml:"output"`
	// The S3 path to which Entity Resolution will write the output table.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/entityresolution_matching_workflow#output_s3_path EntityresolutionMatchingWorkflow#output_s3_path}
	OutputS3Path *string `field:"required" json:"outputS3Path" yaml:"outputS3Path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/entityresolution_matching_workflow#apply_normalization EntityresolutionMatchingWorkflow#apply_normalization}.
	ApplyNormalization interface{} `field:"optional" json:"applyNormalization" yaml:"applyNormalization"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/entityresolution_matching_workflow#kms_arn EntityresolutionMatchingWorkflow#kms_arn}.
	KmsArn *string `field:"optional" json:"kmsArn" yaml:"kmsArn"`
}

