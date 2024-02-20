package omicsworkflow


type OmicsWorkflowParameterTemplate struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/omics_workflow#description OmicsWorkflow#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/omics_workflow#optional OmicsWorkflow#optional}.
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}

