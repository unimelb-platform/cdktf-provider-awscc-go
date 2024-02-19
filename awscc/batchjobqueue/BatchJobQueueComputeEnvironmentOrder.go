package batchjobqueue


type BatchJobQueueComputeEnvironmentOrder struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_queue#compute_environment BatchJobQueue#compute_environment}.
	ComputeEnvironment *string `field:"required" json:"computeEnvironment" yaml:"computeEnvironment"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_job_queue#order BatchJobQueue#order}.
	Order *float64 `field:"required" json:"order" yaml:"order"`
}

