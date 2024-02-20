package batchcomputeenvironment


type BatchComputeEnvironmentUpdatePolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_compute_environment#job_execution_timeout_minutes BatchComputeEnvironment#job_execution_timeout_minutes}.
	JobExecutionTimeoutMinutes *float64 `field:"optional" json:"jobExecutionTimeoutMinutes" yaml:"jobExecutionTimeoutMinutes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/batch_compute_environment#terminate_jobs_on_update BatchComputeEnvironment#terminate_jobs_on_update}.
	TerminateJobsOnUpdate interface{} `field:"optional" json:"terminateJobsOnUpdate" yaml:"terminateJobsOnUpdate"`
}

