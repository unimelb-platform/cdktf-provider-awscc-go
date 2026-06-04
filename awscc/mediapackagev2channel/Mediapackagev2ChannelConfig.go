package mediapackagev2channel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Mediapackagev2ChannelConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mediapackagev2_channel#channel_group_name Mediapackagev2Channel#channel_group_name}.
	ChannelGroupName *string `field:"required" json:"channelGroupName" yaml:"channelGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mediapackagev2_channel#channel_name Mediapackagev2Channel#channel_name}.
	ChannelName *string `field:"required" json:"channelName" yaml:"channelName"`
	// <p>Enter any descriptive text that helps you to identify the channel.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mediapackagev2_channel#description Mediapackagev2Channel#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// <p>The configuration for input switching based on the media quality confidence score (MQCS) as provided from AWS Elemental MediaLive.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mediapackagev2_channel#input_switch_configuration Mediapackagev2Channel#input_switch_configuration}
	InputSwitchConfiguration *Mediapackagev2ChannelInputSwitchConfiguration `field:"optional" json:"inputSwitchConfiguration" yaml:"inputSwitchConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mediapackagev2_channel#input_type Mediapackagev2Channel#input_type}.
	InputType *string `field:"optional" json:"inputType" yaml:"inputType"`
	// <p>The settings for what common media server data (CMSD) headers AWS Elemental MediaPackage includes in responses to the CDN.</p>.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mediapackagev2_channel#output_header_configuration Mediapackagev2Channel#output_header_configuration}
	OutputHeaderConfiguration *Mediapackagev2ChannelOutputHeaderConfiguration `field:"optional" json:"outputHeaderConfiguration" yaml:"outputHeaderConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/mediapackagev2_channel#tags Mediapackagev2Channel#tags}.
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

