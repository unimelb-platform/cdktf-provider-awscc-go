package transferworkflow


type TransferWorkflowStepsTagStepDetailsTags struct {
	// The name assigned to the tag that you create.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_workflow#key TransferWorkflow#key}
	Key *string `field:"required" json:"key" yaml:"key"`
	// The value that corresponds to the key.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/transfer_workflow#value TransferWorkflow#value}
	Value *string `field:"required" json:"value" yaml:"value"`
}

