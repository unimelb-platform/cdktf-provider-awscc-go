package fmsnotificationchannel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FmsNotificationChannelConfig struct {
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
	// A resource ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/fms_notification_channel#sns_role_name FmsNotificationChannel#sns_role_name}
	SnsRoleName *string `field:"required" json:"snsRoleName" yaml:"snsRoleName"`
	// A resource ARN.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/fms_notification_channel#sns_topic_arn FmsNotificationChannel#sns_topic_arn}
	SnsTopicArn *string `field:"required" json:"snsTopicArn" yaml:"snsTopicArn"`
}

