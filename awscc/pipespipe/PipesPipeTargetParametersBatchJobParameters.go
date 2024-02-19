package pipespipe


type PipesPipeTargetParametersBatchJobParameters struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pipes_pipe#job_definition PipesPipe#job_definition}.
	JobDefinition *string `field:"required" json:"jobDefinition" yaml:"jobDefinition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pipes_pipe#job_name PipesPipe#job_name}.
	JobName *string `field:"required" json:"jobName" yaml:"jobName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pipes_pipe#array_properties PipesPipe#array_properties}.
	ArrayProperties *PipesPipeTargetParametersBatchJobParametersArrayProperties `field:"optional" json:"arrayProperties" yaml:"arrayProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pipes_pipe#container_overrides PipesPipe#container_overrides}.
	ContainerOverrides *PipesPipeTargetParametersBatchJobParametersContainerOverrides `field:"optional" json:"containerOverrides" yaml:"containerOverrides"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pipes_pipe#depends_on PipesPipe#depends_on}.
	DependsOn interface{} `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pipes_pipe#parameters PipesPipe#parameters}.
	Parameters *map[string]*string `field:"optional" json:"parameters" yaml:"parameters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/pipes_pipe#retry_strategy PipesPipe#retry_strategy}.
	RetryStrategy *PipesPipeTargetParametersBatchJobParametersRetryStrategy `field:"optional" json:"retryStrategy" yaml:"retryStrategy"`
}

