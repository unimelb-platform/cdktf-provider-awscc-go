package connectcampaignsv2campaign

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type Connectcampaignsv2CampaignConfig struct {
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
	// The possible types of channel subtype config parameters.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connectcampaignsv2_campaign#channel_subtype_config Connectcampaignsv2Campaign#channel_subtype_config}
	ChannelSubtypeConfig *Connectcampaignsv2CampaignChannelSubtypeConfig `field:"required" json:"channelSubtypeConfig" yaml:"channelSubtypeConfig"`
	// Amazon Connect Instance Id.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connectcampaignsv2_campaign#connect_instance_id Connectcampaignsv2Campaign#connect_instance_id}
	ConnectInstanceId *string `field:"required" json:"connectInstanceId" yaml:"connectInstanceId"`
	// Campaign name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connectcampaignsv2_campaign#name Connectcampaignsv2Campaign#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Communication limits config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connectcampaignsv2_campaign#communication_limits_override Connectcampaignsv2Campaign#communication_limits_override}
	CommunicationLimitsOverride *Connectcampaignsv2CampaignCommunicationLimitsOverride `field:"optional" json:"communicationLimitsOverride" yaml:"communicationLimitsOverride"`
	// Campaign communication time config.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connectcampaignsv2_campaign#communication_time_config Connectcampaignsv2Campaign#communication_time_config}
	CommunicationTimeConfig *Connectcampaignsv2CampaignCommunicationTimeConfig `field:"optional" json:"communicationTimeConfig" yaml:"communicationTimeConfig"`
	// Arn.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connectcampaignsv2_campaign#connect_campaign_flow_arn Connectcampaignsv2Campaign#connect_campaign_flow_arn}
	ConnectCampaignFlowArn *string `field:"optional" json:"connectCampaignFlowArn" yaml:"connectCampaignFlowArn"`
	// Campaign schedule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connectcampaignsv2_campaign#schedule Connectcampaignsv2Campaign#schedule}
	Schedule *Connectcampaignsv2CampaignSchedule `field:"optional" json:"schedule" yaml:"schedule"`
	// The possible source of the campaign.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connectcampaignsv2_campaign#source Connectcampaignsv2Campaign#source}
	Source *Connectcampaignsv2CampaignSource `field:"optional" json:"source" yaml:"source"`
	// One or more tags.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/connectcampaignsv2_campaign#tags Connectcampaignsv2Campaign#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

