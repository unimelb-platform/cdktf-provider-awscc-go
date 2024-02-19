package chatbotslackchannelconfiguration

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ChatbotSlackChannelConfigurationConfig struct {
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
	// The name of the configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/chatbot_slack_channel_configuration#configuration_name ChatbotSlackChannelConfiguration#configuration_name}
	ConfigurationName *string `field:"required" json:"configurationName" yaml:"configurationName"`
	// The ARN of the IAM role that defines the permissions for AWS Chatbot.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/chatbot_slack_channel_configuration#iam_role_arn ChatbotSlackChannelConfiguration#iam_role_arn}
	IamRoleArn *string `field:"required" json:"iamRoleArn" yaml:"iamRoleArn"`
	// The id of the Slack channel.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/chatbot_slack_channel_configuration#slack_channel_id ChatbotSlackChannelConfiguration#slack_channel_id}
	SlackChannelId *string `field:"required" json:"slackChannelId" yaml:"slackChannelId"`
	// The id of the Slack workspace.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/chatbot_slack_channel_configuration#slack_workspace_id ChatbotSlackChannelConfiguration#slack_workspace_id}
	SlackWorkspaceId *string `field:"required" json:"slackWorkspaceId" yaml:"slackWorkspaceId"`
	// The list of IAM policy ARNs that are applied as channel guardrails.
	//
	// The AWS managed 'AdministratorAccess' policy is applied as a default if this is not set.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/chatbot_slack_channel_configuration#guardrail_policies ChatbotSlackChannelConfiguration#guardrail_policies}
	GuardrailPolicies *[]*string `field:"optional" json:"guardrailPolicies" yaml:"guardrailPolicies"`
	// Specifies the logging level for this configuration:ERROR,INFO or NONE.
	//
	// This property affects the log entries pushed to Amazon CloudWatch logs
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/chatbot_slack_channel_configuration#logging_level ChatbotSlackChannelConfiguration#logging_level}
	LoggingLevel *string `field:"optional" json:"loggingLevel" yaml:"loggingLevel"`
	// ARNs of SNS topics which delivers notifications to AWS Chatbot, for example CloudWatch alarm notifications.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/chatbot_slack_channel_configuration#sns_topic_arns ChatbotSlackChannelConfiguration#sns_topic_arns}
	SnsTopicArns *[]*string `field:"optional" json:"snsTopicArns" yaml:"snsTopicArns"`
	// Enables use of a user role requirement in your chat configuration.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/chatbot_slack_channel_configuration#user_role_required ChatbotSlackChannelConfiguration#user_role_required}
	UserRoleRequired interface{} `field:"optional" json:"userRoleRequired" yaml:"userRoleRequired"`
}

