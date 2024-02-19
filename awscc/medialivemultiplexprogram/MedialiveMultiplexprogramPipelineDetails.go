package medialivemultiplexprogram


type MedialiveMultiplexprogramPipelineDetails struct {
	// Identifies the channel pipeline that is currently active for the pipeline (identified by PipelineId) in the multiplex.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/medialive_multiplexprogram#active_channel_pipeline MedialiveMultiplexprogram#active_channel_pipeline}
	ActiveChannelPipeline *string `field:"optional" json:"activeChannelPipeline" yaml:"activeChannelPipeline"`
	// Identifies a specific pipeline in the multiplex.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/medialive_multiplexprogram#pipeline_id MedialiveMultiplexprogram#pipeline_id}
	PipelineId *string `field:"optional" json:"pipelineId" yaml:"pipelineId"`
}

