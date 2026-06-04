package gluejob


type GlueJobExecutionProperty struct {
	// The maximum number of concurrent runs allowed for the job.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/glue_job#max_concurrent_runs GlueJob#max_concurrent_runs}
	MaxConcurrentRuns *float64 `field:"optional" json:"maxConcurrentRuns" yaml:"maxConcurrentRuns"`
}

