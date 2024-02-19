package ioteventsalarmmodel


type IoteventsAlarmModelAlarmEventActionsAlarmActionsSns struct {
	// The ARN of the Amazon SNS target where the message is sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#target_arn IoteventsAlarmModel#target_arn}
	TargetArn *string `field:"required" json:"targetArn" yaml:"targetArn"`
	// Information needed to configure the payload.
	//
	// By default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the alarm model instance and the event triggered the action. To configure the action payload, you can use `contentExpression`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#payload IoteventsAlarmModel#payload}
	Payload *IoteventsAlarmModelAlarmEventActionsAlarmActionsSnsPayload `field:"optional" json:"payload" yaml:"payload"`
}

