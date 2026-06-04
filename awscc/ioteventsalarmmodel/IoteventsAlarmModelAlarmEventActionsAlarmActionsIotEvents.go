package ioteventsalarmmodel


type IoteventsAlarmModelAlarmEventActionsAlarmActionsIotEvents struct {
	// The name of the ITE input where the data is sent.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotevents_alarm_model#input_name IoteventsAlarmModel#input_name}
	InputName *string `field:"optional" json:"inputName" yaml:"inputName"`
	// You can configure the action payload when you send a message to an ITE input.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/awscc/1.49.0/docs/resources/iotevents_alarm_model#payload IoteventsAlarmModel#payload}
	Payload *IoteventsAlarmModelAlarmEventActionsAlarmActionsIotEventsPayload `field:"optional" json:"payload" yaml:"payload"`
}

