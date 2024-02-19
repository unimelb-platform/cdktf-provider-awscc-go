package batchjobdefinition


type BatchJobDefinitionRetryStrategy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#attempts BatchJobDefinition#attempts}.
	Attempts *float64 `field:"optional" json:"attempts" yaml:"attempts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_definition#evaluate_on_exit BatchJobDefinition#evaluate_on_exit}.
	EvaluateOnExit interface{} `field:"optional" json:"evaluateOnExit" yaml:"evaluateOnExit"`
}

