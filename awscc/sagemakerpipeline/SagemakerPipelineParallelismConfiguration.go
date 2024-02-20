package sagemakerpipeline


type SagemakerPipelineParallelismConfiguration struct {
	// Maximum parallel execution steps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/sagemaker_pipeline#max_parallel_execution_steps SagemakerPipeline#max_parallel_execution_steps}
	MaxParallelExecutionSteps *float64 `field:"required" json:"maxParallelExecutionSteps" yaml:"maxParallelExecutionSteps"`
}

