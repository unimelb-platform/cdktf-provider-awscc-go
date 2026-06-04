package securitylakesubscribernotification

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecuritylakeSubscriberNotificationConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/securitylake_subscriber_notification#notification_configuration SecuritylakeSubscriberNotification#notification_configuration}.
	NotificationConfiguration *SecuritylakeSubscriberNotificationNotificationConfiguration `field:"required" json:"notificationConfiguration" yaml:"notificationConfiguration"`
	// The ARN for the subscriber.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/securitylake_subscriber_notification#subscriber_arn SecuritylakeSubscriberNotification#subscriber_arn}
	SubscriberArn *string `field:"required" json:"subscriberArn" yaml:"subscriberArn"`
}

