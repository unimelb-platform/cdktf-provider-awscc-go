package transferworkflow


type TransferWorkflowOnExceptionStepsDeleteStepDetails struct {
	// The name of the step, used as an identifier.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_workflow#name TransferWorkflow#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// Specifies which file to use as input to the workflow step.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_workflow#source_file_location TransferWorkflow#source_file_location}
	SourceFileLocation *string `field:"optional" json:"sourceFileLocation" yaml:"sourceFileLocation"`
}

