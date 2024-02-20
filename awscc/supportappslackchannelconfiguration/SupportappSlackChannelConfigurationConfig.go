package supportappslackchannelconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SupportappSlackChannelConfigurationConfig struct {
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
	// The channel ID in Slack, which identifies a channel within a workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/supportapp_slack_channel_configuration#channel_id SupportappSlackChannelConfiguration#channel_id}
	ChannelId *string `field:"required" json:"channelId" yaml:"channelId"`
	// The Amazon Resource Name (ARN) of an IAM role that grants the AWS Support App access to perform operations for AWS services.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/supportapp_slack_channel_configuration#channel_role_arn SupportappSlackChannelConfiguration#channel_role_arn}
	ChannelRoleArn *string `field:"required" json:"channelRoleArn" yaml:"channelRoleArn"`
	// The severity level of a support case that a customer wants to get notified for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/supportapp_slack_channel_configuration#notify_on_case_severity SupportappSlackChannelConfiguration#notify_on_case_severity}
	NotifyOnCaseSeverity *string `field:"required" json:"notifyOnCaseSeverity" yaml:"notifyOnCaseSeverity"`
	// The team ID in Slack, which uniquely identifies a workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/supportapp_slack_channel_configuration#team_id SupportappSlackChannelConfiguration#team_id}
	TeamId *string `field:"required" json:"teamId" yaml:"teamId"`
	// The channel name in Slack.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/supportapp_slack_channel_configuration#channel_name SupportappSlackChannelConfiguration#channel_name}
	ChannelName *string `field:"optional" json:"channelName" yaml:"channelName"`
	// Whether to notify when a correspondence is added to a case.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/supportapp_slack_channel_configuration#notify_on_add_correspondence_to_case SupportappSlackChannelConfiguration#notify_on_add_correspondence_to_case}
	NotifyOnAddCorrespondenceToCase interface{} `field:"optional" json:"notifyOnAddCorrespondenceToCase" yaml:"notifyOnAddCorrespondenceToCase"`
	// Whether to notify when a case is created or reopened.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/supportapp_slack_channel_configuration#notify_on_create_or_reopen_case SupportappSlackChannelConfiguration#notify_on_create_or_reopen_case}
	NotifyOnCreateOrReopenCase interface{} `field:"optional" json:"notifyOnCreateOrReopenCase" yaml:"notifyOnCreateOrReopenCase"`
	// Whether to notify when a case is resolved.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/supportapp_slack_channel_configuration#notify_on_resolve_case SupportappSlackChannelConfiguration#notify_on_resolve_case}
	NotifyOnResolveCase interface{} `field:"optional" json:"notifyOnResolveCase" yaml:"notifyOnResolveCase"`
}

