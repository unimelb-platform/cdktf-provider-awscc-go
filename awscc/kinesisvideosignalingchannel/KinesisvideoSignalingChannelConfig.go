package kinesisvideosignalingchannel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KinesisvideoSignalingChannelConfig struct {
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
	// The period of time a signaling channel retains undelivered messages before they are discarded.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisvideo_signaling_channel#message_ttl_seconds KinesisvideoSignalingChannel#message_ttl_seconds}
	MessageTtlSeconds *float64 `field:"optional" json:"messageTtlSeconds" yaml:"messageTtlSeconds"`
	// The name of the Kinesis Video Signaling Channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisvideo_signaling_channel#name KinesisvideoSignalingChannel#name}
	Name *string `field:"optional" json:"name" yaml:"name"`
	// An array of key-value pairs to apply to this resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisvideo_signaling_channel#tags KinesisvideoSignalingChannel#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
	// The type of the Kinesis Video Signaling Channel to create. Currently, SINGLE_MASTER is the only supported channel type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/kinesisvideo_signaling_channel#type KinesisvideoSignalingChannel#type}
	Type *string `field:"optional" json:"type" yaml:"type"`
}

