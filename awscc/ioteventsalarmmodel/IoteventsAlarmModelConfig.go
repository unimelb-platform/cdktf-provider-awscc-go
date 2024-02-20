package ioteventsalarmmodel

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IoteventsAlarmModelConfig struct {
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
	// Defines when your alarm is invoked.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#alarm_rule IoteventsAlarmModel#alarm_rule}
	AlarmRule *IoteventsAlarmModelAlarmRule `field:"required" json:"alarmRule" yaml:"alarmRule"`
	// The ARN of the role that grants permission to AWS IoT Events to perform its operations.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#role_arn IoteventsAlarmModel#role_arn}
	RoleArn *string `field:"required" json:"roleArn" yaml:"roleArn"`
	// Contains the configuration information of alarm state changes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#alarm_capabilities IoteventsAlarmModel#alarm_capabilities}
	AlarmCapabilities *IoteventsAlarmModelAlarmCapabilities `field:"optional" json:"alarmCapabilities" yaml:"alarmCapabilities"`
	// Contains information about one or more alarm actions.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#alarm_event_actions IoteventsAlarmModel#alarm_event_actions}
	AlarmEventActions *IoteventsAlarmModelAlarmEventActions `field:"optional" json:"alarmEventActions" yaml:"alarmEventActions"`
	// A brief description of the alarm model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#alarm_model_description IoteventsAlarmModel#alarm_model_description}
	AlarmModelDescription *string `field:"optional" json:"alarmModelDescription" yaml:"alarmModelDescription"`
	// The name of the alarm model.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#alarm_model_name IoteventsAlarmModel#alarm_model_name}
	AlarmModelName *string `field:"optional" json:"alarmModelName" yaml:"alarmModelName"`
	// The value used to identify a alarm instance.
	//
	// When a device or system sends input, a new alarm instance with a unique key value is created. AWS IoT Events can continue to route input to its corresponding alarm instance based on this identifying information.
	//
	// This parameter uses a JSON-path expression to select the attribute-value pair in the message payload that is used for identification. To route the message to the correct alarm instance, the device must send a message payload that contains the same attribute-value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#key IoteventsAlarmModel#key}
	Key *string `field:"optional" json:"key" yaml:"key"`
	// A non-negative integer that reflects the severity level of the alarm.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#severity IoteventsAlarmModel#severity}
	Severity *float64 `field:"optional" json:"severity" yaml:"severity"`
	// An array of key-value pairs to apply to this resource.
	//
	// For more information, see [Tag](https://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-resource-tags.html).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#tags IoteventsAlarmModel#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}

