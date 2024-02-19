package batchjobdefinition


type BatchJobDefinitionRetryStrategyEvaluateOnExit struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#action BatchJobDefinition#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#on_exit_code BatchJobDefinition#on_exit_code}.
	OnExitCode *string `field:"optional" json:"onExitCode" yaml:"onExitCode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#on_reason BatchJobDefinition#on_reason}.
	OnReason *string `field:"optional" json:"onReason" yaml:"onReason"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#on_status_reason BatchJobDefinition#on_status_reason}.
	OnStatusReason *string `field:"optional" json:"onStatusReason" yaml:"onStatusReason"`
}

