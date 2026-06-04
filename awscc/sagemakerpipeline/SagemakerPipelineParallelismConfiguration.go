package sagemakerpipeline


type SagemakerPipelineParallelismConfiguration struct {
	// Maximum parallel execution steps.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/sagemaker_pipeline#max_parallel_execution_steps SagemakerPipeline#max_parallel_execution_steps}
	MaxParallelExecutionSteps *float64 `field:"optional" json:"maxParallelExecutionSteps" yaml:"maxParallelExecutionSteps"`
}

