package imagebuilderimagepipeline


type ImagebuilderImagePipelineSchedule struct {
	// The condition configures when the pipeline should trigger a new image build.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_image_pipeline#pipeline_execution_start_condition ImagebuilderImagePipeline#pipeline_execution_start_condition}
	PipelineExecutionStartCondition *string `field:"optional" json:"pipelineExecutionStartCondition" yaml:"pipelineExecutionStartCondition"`
	// The expression determines how often EC2 Image Builder evaluates your pipelineExecutionStartCondition.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/imagebuilder_image_pipeline#schedule_expression ImagebuilderImagePipeline#schedule_expression}
	ScheduleExpression *string `field:"optional" json:"scheduleExpression" yaml:"scheduleExpression"`
}

