package ioteventsalarmmodel


type IoteventsAlarmModelAlarmEventActionsAlarmActionsIotEvents struct {
	// The name of the AWS IoT Events input where the data is sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#input_name IoteventsAlarmModel#input_name}
	InputName *string `field:"required" json:"inputName" yaml:"inputName"`
	// Information needed to configure the payload.
	//
	// By default, AWS IoT Events generates a standard payload in JSON for any action. This action payload contains all attribute-value pairs that have the information about the alarm model instance and the event triggered the action. To configure the action payload, you can use `contentExpression`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/0.70.0/docs/resources/iotevents_alarm_model#payload IoteventsAlarmModel#payload}
	Payload *IoteventsAlarmModelAlarmEventActionsAlarmActionsIotEventsPayload `field:"optional" json:"payload" yaml:"payload"`
}

