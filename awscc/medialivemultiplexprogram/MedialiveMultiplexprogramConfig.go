package medialivemultiplexprogram

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MedialiveMultiplexprogramConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The MediaLive channel associated with the program.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/medialive_multiplexprogram#channel_id MedialiveMultiplexprogram#channel_id}
	ChannelId *string `field:"optional" json:"channelId" yaml:"channelId"`
	// The ID of the multiplex that the program belongs to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/medialive_multiplexprogram#multiplex_id MedialiveMultiplexprogram#multiplex_id}
	MultiplexId *string `field:"optional" json:"multiplexId" yaml:"multiplexId"`
	// The settings for this multiplex program.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/medialive_multiplexprogram#multiplex_program_settings MedialiveMultiplexprogram#multiplex_program_settings}
	MultiplexProgramSettings *MedialiveMultiplexprogramMultiplexProgramSettings `field:"optional" json:"multiplexProgramSettings" yaml:"multiplexProgramSettings"`
	// The packet identifier map for this multiplex program.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/medialive_multiplexprogram#packet_identifiers_map MedialiveMultiplexprogram#packet_identifiers_map}
	PacketIdentifiersMap *MedialiveMultiplexprogramPacketIdentifiersMap `field:"optional" json:"packetIdentifiersMap" yaml:"packetIdentifiersMap"`
	// Contains information about the current sources for the specified program in the specified multiplex.
	//
	// Keep in mind that each multiplex pipeline connects to both pipelines in a given source channel (the channel identified by the program). But only one of those channel pipelines is ever active at one time.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/medialive_multiplexprogram#pipeline_details MedialiveMultiplexprogram#pipeline_details}
	PipelineDetails interface{} `field:"optional" json:"pipelineDetails" yaml:"pipelineDetails"`
	// The settings for this multiplex program.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/medialive_multiplexprogram#preferred_channel_pipeline MedialiveMultiplexprogram#preferred_channel_pipeline}
	PreferredChannelPipeline *string `field:"optional" json:"preferredChannelPipeline" yaml:"preferredChannelPipeline"`
	// The name of the multiplex program.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/medialive_multiplexprogram#program_name MedialiveMultiplexprogram#program_name}
	ProgramName *string `field:"optional" json:"programName" yaml:"programName"`
}

