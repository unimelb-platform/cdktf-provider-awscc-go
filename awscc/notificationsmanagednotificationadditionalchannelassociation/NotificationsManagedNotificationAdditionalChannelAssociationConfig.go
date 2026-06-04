package notificationsmanagednotificationadditionalchannelassociation

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NotificationsManagedNotificationAdditionalChannelAssociationConfig struct {
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
	// ARN identifier of the channel. Example: arn:aws:chatbot::123456789012:chat-configuration/slack-channel/security-ops.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/notifications_managed_notification_additional_channel_association#channel_arn NotificationsManagedNotificationAdditionalChannelAssociation#channel_arn}
	ChannelArn *string `field:"required" json:"channelArn" yaml:"channelArn"`
	// ARN identifier of the Managed Notification. Example: arn:aws:notifications::381491923782:managed-notification-configuration/category/AWS-Health/sub-category/Billing.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/notifications_managed_notification_additional_channel_association#managed_notification_configuration_arn NotificationsManagedNotificationAdditionalChannelAssociation#managed_notification_configuration_arn}
	ManagedNotificationConfigurationArn *string `field:"required" json:"managedNotificationConfigurationArn" yaml:"managedNotificationConfigurationArn"`
}

