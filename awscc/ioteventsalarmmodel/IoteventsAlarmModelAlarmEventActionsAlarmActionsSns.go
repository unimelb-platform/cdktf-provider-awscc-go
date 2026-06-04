package ioteventsalarmmodel


type IoteventsAlarmModelAlarmEventActionsAlarmActionsSns struct {
	// You can configure the action payload when you send a message as an Amazon SNS push notification.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotevents_alarm_model#payload IoteventsAlarmModel#payload}
	Payload *IoteventsAlarmModelAlarmEventActionsAlarmActionsSnsPayload `field:"optional" json:"payload" yaml:"payload"`
	// The ARN of the Amazon SNS target where the message is sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotevents_alarm_model#target_arn IoteventsAlarmModel#target_arn}
	TargetArn *string `field:"optional" json:"targetArn" yaml:"targetArn"`
}

